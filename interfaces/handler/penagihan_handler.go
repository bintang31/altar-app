package interfaces

import (
	"altar-app/application"
	"altar-app/domain/entity"
	"altar-app/infrastructure/auth"
	"altar-app/infrastructure/queue/producer"
	"altar-app/infrastructure/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Penagihans struct defines the dependencies that will be used
type Penagihans struct {
	pn application.PenagihanAppInterface
	us application.UserAppInterface
	pt application.PetugasAppInterface
	tr application.TransactionsAppInterface
	pl application.PelangganAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewPenagihans constructor
func NewPenagihans(pn application.PenagihanAppInterface, us application.UserAppInterface, pt application.PetugasAppInterface, tr application.TransactionsAppInterface, pl application.PelangganAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Penagihans {
	return &Penagihans{
		pn: pn,
		us: us,
		pt: pt,
		tr: tr,
		pl: pl,
		rd: rd,
		tk: tk,
	}
}

//GetPenagihans : Get All Data Penagihan
func (p *Penagihans) GetPenagihans(c *gin.Context) {
	var err error
	//Check if the user is authenticated first
	metadata, err := p.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := p.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	//Get user profile by user Login
	user, err := p.us.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("userID :%+v\n", user.ID)
	penagihans := entity.PenagihansSrKolektifs{} //customize user
	var paramFilter *entity.PenagihansParams
	if err := c.ShouldBindQuery(&paramFilter); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_filter": "invalid param filter",
		})
		return
	}
	paramFilter.UserID = userID
	//us, err = application.UserApp.GetUsers()
	penagihans, err = p.pn.GetPenagihanByParam(paramFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var penagihanPelanggan = make(map[string]interface{})
	penagihanPelanggan["penagihan_pelanggan"] = penagihans
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("030102").SetData(penagihanPelanggan).Build(c))
}

//BayarTagihanPelanggan : Handler untuk Bayar Tagihan Pelanggan to loket
func (p *Penagihans) BayarTagihanPelanggan(c *gin.Context) {
	var postDataTerima *entity.Bayar
	if err := c.ShouldBindJSON(&postDataTerima); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	var penagihan *entity.Penagihan
	var responseloket *entity.ResponseLoket
	var tokenErr = map[string]string{}
	rb := &response.ResponseBuilder{}
	var err error
	//Check if the user is authenticated first
	metadata, err := p.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := p.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	//Get user profile by user Login
	user, err := p.us.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//Get Petugas profile by user Login
	petugas, err := p.pt.GetProfilePetugas(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	penagihan, err = p.pn.GetPenagihanByNosamb(postDataTerima.Nosamb)
	fmt.Printf("Penagihan :%+v\n", penagihan)
	//validate the request:
	if user.Pin != postDataTerima.Pin {
		c.JSON(http.StatusUnprocessableEntity, rb.SetResponse("020103").SetData("Invalid PIN").Build(c))
		return
	}
	if penagihan.TotalTagihan > petugas.SisaLimit {
		c.JSON(http.StatusUnprocessableEntity, rb.SetResponse("020104").SetData("Total Tagihan Melebihi Limit Petugas").Build(c))
		return
	}

	postDataTerima.Pdam = user.Pdam
	postDataTerima.UserLoket = user.Name
	responseloket, tokenErr = p.pn.BayarTagihanByNosamb(postDataTerima)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, tokenErr)
		return
	}

	if responseloket.Message != "Payment success" {
		c.JSON(http.StatusUnprocessableEntity, rb.SetResponse("020102").SetData(responseloket.Message).Build(c))
		return
	}

	var trxInsert = entity.Transaction{}
	trxInsert.TotalDrd = penagihan.TagihanAir
	trxInsert.Total = penagihan.TotalTagihan
	trxInsert.Status = 2
	trxInsert.Notes = postDataTerima.Notes
	trxInsert.Pdam = penagihan.KodePdam
	trxInsert.CreatedBy = user.ID
	trxInsert.Jenis = penagihan.StatusKolektif
	trxInsert.Pelanggan = penagihan.Nosamb
	trxInsert.Denda = penagihan.TotalDenda
	trxInsert.PeriodeBilling = penagihan.PeriodeTagihan
	trxInsert.TotalAir = penagihan.TotalTagihanAir
	trxInsert.TotalNonair = penagihan.TotalTagihanNonair
	trxInsert.LoketMessage = responseloket.Message
	trxInsert.LoketMessageCode = responseloket.Code
	trxInsert.LoketMessageStatus = responseloket.Message
	_, tokenErr = p.tr.SaveTransactions(&trxInsert)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, tokenErr)
		return
	}

	var drdUpdate = entity.Drd{}
	drdUpdate.Nosamb = penagihan.Nosamb
	drdUpdate.Lunas = "1"
	drdUpdate.TransactionsID = trxInsert.ID
	_, tokenErr = p.pl.UpdateDrdByNosamb(&drdUpdate)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, tokenErr)
		return
	}

	var nonairUpdate = entity.Nonair{}
	nonairUpdate.Nomor = penagihan.Nosamb
	nonairUpdate.Lunas = "1"
	nonairUpdate.TransactionsID = trxInsert.ID
	_, tokenErr = p.pl.UpdateNonAirByNosamb(&nonairUpdate)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, tokenErr)
		return
	}

	//fmt.Printf("Response From Loket :%+v\n", responseloket)
	var penagihanPelanggan = make(map[string]interface{})
	penagihanPelanggan["transaction_id"] = trxInsert.ID
	penagihanPelanggan["message"] = responseloket.Message

	c.JSON(http.StatusOK, rb.SetResponse("030102").SetData(penagihanPelanggan).Build(c))
}

//BayarTagihanPelangganAsync : Bayar Tagihan Bulk Nosamb
func (p *Penagihans) BayarTagihanPelangganAsync(c *gin.Context) {
	var postDataTerima entity.TransactionsBulkData
	if err := c.ShouldBindJSON(&postDataTerima); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	var tokenErr = map[string]string{}
	//Check if the user is authenticated first
	metadata, err := p.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := p.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	//Get user profile by user Login
	user, err := p.us.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//validate the request:
	if user.Pin != postDataTerima.Pin {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"message": "Invalid PIN"})
		return
	}
	for _, sr := range postDataTerima.SambunganRumah {
		var trxInsert = entity.Transaction{}
		trxInsert.Pelanggan = sr.Nosamb
		trxInsert.TotalDrd = 100
		trxInsert.Denda = 100
		trxInsert.Total = 100
		trxInsert.Status = 5
		trxInsert.Pdam = "MJI"
		trxInsert.CreatedBy = 2
		trxInsert.Jenis = "sambungan_rumah"
		_, tokenErr = p.tr.SaveTransactions(&trxInsert)
		if tokenErr != nil {
			c.JSON(http.StatusInternalServerError, tokenErr)
			return
		}
	}

	for _, kl := range postDataTerima.Kolektif {
		var trxKolektifInsert = entity.TransactionsKolektif{}
		trxKolektifInsert.KodeKolektif = kl.Kodekolektif
		trxKolektifInsert.Pelanggan = kl.TotalPelanggan
		trxKolektifInsert.TotalTagihan = 200000
		trxKolektifInsert.Pdam = "MJI"
		trxKolektifInsert.CreatedBy = 3
		_, tokenErr = p.tr.SaveTransactionsKolektif(&trxKolektifInsert)
		if tokenErr != nil {
			c.JSON(http.StatusInternalServerError, tokenErr)
			return
		}
		for _, klp := range kl.PelangganKolektif {
			var trxInsert = entity.Transaction{}
			trxInsert.Pelanggan = klp.Nosamb
			trxInsert.TotalDrd = 100
			trxInsert.Denda = 100
			trxInsert.Total = 100
			trxInsert.Status = 5
			trxInsert.Pdam = "MJI"
			trxInsert.CreatedBy = 3
			trxInsert.Jenis = "kolektif"
			_, tokenErr = p.tr.SaveTransactions(&trxInsert)
			if tokenErr != nil {
				c.JSON(http.StatusInternalServerError, tokenErr)
				return
			}
		}
	}

	payload := make(map[string]interface{})
	payload["action"] = "payment_nosamb"
	payload["nosamb"] = "nomor"
	payload["pdam"] = user.Pdam
	payload["userid"] = user.ID
	//queue.SendWithParam(param)
	data, _ := json.Marshal(payload)
	producer.Producer.CreateItem("mobileloket-process", string(data))

	var penagihanPelanggan = make(map[string]interface{})
	penagihanPelanggan["penagihan_pelanggan"] = postDataTerima
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("030102").SetData(penagihanPelanggan).Build(c))
}

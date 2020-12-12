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
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewPenagihans constructor
func NewPenagihans(pn application.PenagihanAppInterface, us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Penagihans {
	return &Penagihans{
		pn: pn,
		us: us,
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
	//us, err = application.UserApp.GetUsers()
	penagihans, err = p.pn.GetPenagihansByUserPDAM(user.ID)
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
	//validate the request:
	if user.Pin != postDataTerima.Pin {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"message": "Invalid PIN"})
		return
	}
	penagihan, err = p.pn.GetPenagihanByNosamb(postDataTerima.Nosamb)
	fmt.Printf("userID :%+v\n", penagihan.Nama)
	postDataTerima.Pdam = user.Pdam
	responseloket, tokenErr = p.pn.BayarTagihanByNosamb(postDataTerima)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, tokenErr)
		return
	}
	fmt.Printf("Response From Loket :%+v\n", responseloket)
	var penagihanPelanggan = make(map[string]interface{})
	penagihanPelanggan["penagihan_pelanggan"] = postDataTerima
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("030102").SetData(penagihanPelanggan).Build(c))
}

//BayarTagihanPelangganBulk : Bayar Tagihan Bulk Nosamb
func (p *Penagihans) BayarTagihanPelangganBulk(c *gin.Context) {
	var postDataTerima entity.Bayar
	if err := c.ShouldBindJSON(&postDataTerima); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
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
	payload := make(map[string]interface{})
	payload["action"] = "payment_nosamb"
	payload["nosamb"] = postDataTerima.Nosamb
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

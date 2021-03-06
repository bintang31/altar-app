package interfaces

import (
	"altar-app/application"
	"altar-app/domain/entity"
	"altar-app/infrastructure/auth"
	"altar-app/infrastructure/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Pelanggans struct defines the dependencies that will be used
type Pelanggans struct {
	pl application.PelangganAppInterface
	pn application.PenagihanAppInterface
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewPelanggans constructor
func NewPelanggans(pl application.PelangganAppInterface, pn application.PenagihanAppInterface, us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Pelanggans {
	return &Pelanggans{
		pl: pl,
		pn: pn,
		us: us,
		rd: rd,
		tk: tk,
	}
}

//GetPelanggans : Get All Pelanggan
func (p *Pelanggans) GetPelanggans(c *gin.Context) {
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
	fmt.Printf("userID :%+v\n", user.Pdam)
	pelanggans := entity.Pelanggans{} //customize user
	//us, err = application.UserApp.GetUsers()
	pelanggans, err = p.pl.GetPelanggans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pelanggans)
}

//GetTagihanPelanggan : Get Tagihan Pelanggan by Nosamb
func (p *Pelanggans) GetTagihanPelanggan(c *gin.Context) {
	var penagihan *entity.Penagihan
	nosamb := c.Param("nosamb")
	var err error
	tagihanair := entity.Drds{}                 //customize tagihanair
	var tagihannonair []map[string]interface{}  //customize tagihannonair
	var riwayattagihan []map[string]interface{} //customize tagihannonair
	//us, err = application.UserApp.GetUsers()
	tagihanair, err = p.pl.GetTagihanAirPelanggansByNosamb(nosamb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	tagihannonair, err = p.pl.GetTagihanNonAirPelanggansByPeriode(nosamb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	riwayattagihan, err = p.pl.GetRiwayatTagihanByNosamb(nosamb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	penagihan, err = p.pn.GetPenagihanByNosamb(nosamb)
	fmt.Printf("userID :%+v\n", penagihan.Nama)
	var tagihanpelanggan = make(map[string]interface{})
	tagihanpelanggan["penagihan_billing_nonair"] = tagihannonair
	tagihanpelanggan["penagihan_billing"] = tagihanair
	tagihanpelanggan["penagihan_pelanggan"] = penagihan
	tagihanpelanggan["riwayat_billing"] = riwayattagihan
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("010102").SetData(tagihanpelanggan).Build(c))
}

//InquiryLoketTagihanPelanggan : Get Inquiry Loket Tagihan Pelanggan by Nosamb
func (p *Pelanggans) InquiryLoketTagihanPelanggan(c *gin.Context) {
	var postDataTerima *entity.InputInquiryPelanggan
	if err := c.ShouldBindJSON(&postDataTerima); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	var err error
	tagihanair := entity.RekairDetails{}

	tagihanair, err = p.pl.InquiryLoketTagihanAirByNosamb(postDataTerima)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var penagihanPelanggan = make(map[string]interface{})
	penagihanPelanggan["penagihan_pelanggan"] = tagihanair
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("030102").SetData(penagihanPelanggan).Build(c))
}

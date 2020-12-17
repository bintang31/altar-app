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

//Petugass struct defines the dependencies that will be used
type Petugass struct {
	pt application.PetugasAppInterface
	pn application.PenagihanAppInterface
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewPetugass constructor
func NewPetugass(pt application.PetugasAppInterface, pn application.PenagihanAppInterface, us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Petugass {
	return &Petugass{
		pt: pt,
		pn: pn,
		us: us,
		rd: rd,
		tk: tk,
	}
}

// GetDataPetugas godoc
// @Summary Retrieves petugas data
// @Produce json
// @success 200 @Success 200 {object} entity.PetugasData
// @Router /get_data [get]
// GetDataPetugas : Get All Data Petugas
func (pts *Petugass) GetDataPetugas(c *gin.Context) {
	var err error
	//Check if the user is authenticated first
	metadata, err := pts.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := pts.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	//Get user profile by user Login
	user, err := pts.us.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("userID :%+v\n", user)

	penagihans := entity.PenagihansSrKolektifs{} //customize user
	//us, err = application.UserApp.GetUsers()
	penagihans, err = pts.pn.GetPenagihansByUserPDAM(user.ID)

	tagihanair := entity.Drds{} //customize user
	//us, err = application.UserApp.GetUsers()
	tagihanair, err = pts.pt.GetTagihanAirByPetugas(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	petugasalldata := entity.PetugasData{}
	petugasalldata.PenagihanPelanggan = penagihans
	petugasalldata.PenagihanBilling = tagihanair
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("010101").SetData(petugasalldata).Build(c))
}

//GetProfilePetugas : Get Data Profile Petugas
func (pts *Petugass) GetProfilePetugas(c *gin.Context) {
	var err error
	//Check if the user is authenticated first
	metadata, err := pts.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := pts.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	//Get user profile by user Login
	petugas, err := pts.pt.GetProfilePetugas(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("userID :%+v\n", petugas)

	var petugasalldata = make(map[string]interface{})
	petugasalldata["profile_petugas"] = petugas
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("010101").SetData(petugasalldata).Build(c))
}

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

	petugas, err := pts.pt.GetProfilePetugas(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	penagihans := entity.PenagihansSrKolektifs{} //customize user
	//us, err = application.UserApp.GetUsers()
	penagihans, err = pts.pn.GetPenagihansByUserPDAM(user.ID)

	tagihanair := entity.DrdbyPetugass{} //customize user
	//us, err = application.UserApp.GetUsers()
	tagihanair, err = pts.pt.GetTagihanAirByPetugas(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	tagihanairkolektif := entity.DrdbyPetugass{} //customize user
	//us, err = application.UserApp.GetUsers()
	tagihanairkolektif, err = pts.pt.GetTagihanAirKolektifByPetugas(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	tagihannonair := entity.Nonairs{} //customize user
	//us, err = application.UserApp.GetUsers()
	tagihannonair, err = pts.pt.GetTagihanNonAirByPetugas(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	petugasalldata := entity.PetugasData{}
	petugasalldata.PenagihanPelanggan = penagihans
	petugasalldata.PenagihanBilling = tagihanair
	petugasalldata.PenagihanBillingKolektif = tagihanairkolektif
	petugasalldata.PenagihanBillingNonair = tagihannonair
	petugasalldata.Petugas = petugas
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

//GetAllPetugas : Get All  Petugas
func (pts *Petugass) GetAllPetugas(c *gin.Context) {
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
	petugas, err := pts.pt.GetPetugas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("userID :%+v\n", userID)
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("010101").SetData(petugas).Build(c))
}

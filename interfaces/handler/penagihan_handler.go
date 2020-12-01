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

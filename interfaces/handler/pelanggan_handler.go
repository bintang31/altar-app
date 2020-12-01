package interfaces

import (
	"altar-app/application"
	"altar-app/domain/entity"
	"altar-app/infrastructure/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Pelanggans struct defines the dependencies that will be used
type Pelanggans struct {
	pl application.PelangganAppInterface
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewPelanggans constructor
func NewPelanggans(pl application.PelangganAppInterface, us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Pelanggans {
	return &Pelanggans{
		pl: pl,
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

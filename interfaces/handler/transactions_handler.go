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

//Transactions struct defines the dependencies that will be used
type Transactions struct {
	tr application.TransactionsAppInterface
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewTransactions constructor
func NewTransactions(tr application.TransactionsAppInterface, us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Transactions {
	return &Transactions{
		tr: tr,
		us: us,
		rd: rd,
		tk: tk,
	}
}

//GetTransactions : Get All Data Transactions
func (t *Transactions) GetTransactions(c *gin.Context) {
	var postDataTerima entity.Bayar
	if err := c.ShouldBindJSON(&postDataTerima); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	var err error
	//Check if the user is authenticated first
	metadata, err := t.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := t.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	//Get user profile by user Login
	user, err := t.us.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//validate the request:

	if user.Pin != postDataTerima.Pin {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"message": "Invalid PIN"})
		return
	}
	fmt.Printf("userID :%+v\n", user)
	var penagihanPelanggan = make(map[string]interface{})
	penagihanPelanggan["penagihan_pelanggan"] = postDataTerima
	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("030102").SetData(penagihanPelanggan).Build(c))
}

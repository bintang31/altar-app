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

//Roles struct defines the dependencies that will be used
type Roles struct {
	rl application.RoleAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewRoles constructor
func NewRoles(rl application.RoleAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Roles {
	return &Roles{
		rl: rl,
		rd: rd,
		tk: tk,
	}
}

//GetRoles : Get All Roles
func (r *Roles) GetRoles(c *gin.Context) {

	//check is the user is authenticated first
	metadata, err := r.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := r.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	fmt.Printf("nosamb :%+v\n", userID)

	roles := entity.Roles{} //customize user
	//us, err = application.UserApp.GetUsers()
	roles, err = r.rl.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rb := &response.ResponseBuilder{}

	c.JSON(http.StatusOK, rb.SetResponse("030102").SetData(roles).Build(c))
}

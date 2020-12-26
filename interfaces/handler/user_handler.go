package interfaces

import (
	"altar-app/application"
	"altar-app/domain/entity"
	"altar-app/infrastructure/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//Users struct defines the dependencies that will be used
type Users struct {
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

//NewUsers constructor
func NewUsers(us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Users {
	return &Users{
		us: us,
		rd: rd,
		tk: tk,
	}
}

//SaveUser : Save new user
func (s *Users) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}
	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser.PublicUser())
}

//GetUsers : Get Data ALl User
func (s *Users) GetUsers(c *gin.Context) {
	users := entity.Users{} //customize user
	var err error
	//us, err = application.UserApp.GetUsers()
	users, err = s.us.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users.PublicUsers())
}

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} entity.User
// @Router /users/{id} [get]
// GetUser : get user By ID
func (s *Users) GetUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := s.us.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user.PublicUser())
}

//UpdateUser : Update user
func (s *Users) UpdateUser(c *gin.Context) {
	var userdata entity.User
	if err := c.ShouldBindJSON(&userdata); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	var err error

	//Check if the user is authenticated first
	metadata, err := s.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := s.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	fmt.Printf("userID :%+v\n", userID)
	newUser, dbErr := s.us.UpdateUser(&userdata)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, dbErr)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

//GetProfileUser : Profile User
func (s *Users) GetProfileUser(c *gin.Context) {
	//Check if the user is authenticated first
	metadata, err := s.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	//lookup the metadata in redis:
	userID, err := s.rd.FetchAuth(metadata.TokenUuid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	fmt.Printf("userID :%+v\n", userID)
	user, err := s.us.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

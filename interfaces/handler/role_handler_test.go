package interfaces

import (
	"altar-app/domain/entity"
	"altar-app/infrastructure/auth"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRoles_Success(t *testing.T) {
	//Mock extracting metadata
	fakeToken.ExtractTokenMetadataFn = func(r *http.Request) (*auth.AccessDetails, error) {
		return &auth.AccessDetails{
			TokenUuid: "0237817a-1546-4ca3-96a4-17621c237f6b",
			UserId:    1,
		}, nil
	}
	//Mocking the fetching of token metadata from redis
	fakeAuth.FetchAuthFn = func(uuid string) (uint64, error) {
		return 1, nil
	}
	roleApp.GetRolesFn = func() ([]entity.Role, error) {
		//remember we are running sensitive info such as email and password
		return []entity.Role{
			{
				ID:   1,
				Name: "Admin",
			},
			{
				ID:   2,
				Name: "Petugas Online",
			},
		}, nil
	}
	r := gin.Default()
	r.GET("/roles", rs.GetRoles)

	req, err := http.NewRequest(http.MethodGet, "/roles", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var roles []entity.Role

	err = json.Unmarshal(rr.Body.Bytes(), &roles)

	assert.Equal(t, rr.Code, 200)
	assert.EqualValues(t, len(roles), 2)
}

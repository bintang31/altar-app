package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var httpCode = map[string]int{
	"200": http.StatusOK,
	"400": http.StatusBadRequest,
	"401": http.StatusUnauthorized,
	"403": http.StatusForbidden,
	"404": http.StatusNotFound,
	"408": http.StatusRequestTimeout,
	"429": http.StatusTooManyRequests,
	"500": http.StatusInternalServerError,
}

type ResponseMessage struct {
	Response   interface{}            `json:"response"`
	Data       interface{}            `json:"data"`
	Version    interface{}            `json:"version"`
	Properties interface{}            `json:"properties"`
	Params     map[string]interface{} `json:"params"`
}

type ResponseSet struct {
	Response   bool `json:"response"`
	Data       bool `json:"data"`
	DataShow   bool `json:"datashow"`
	Version    bool `json:"version"`
	Properties bool `json:"properties"`
	Params     bool `json:"params"`
}

type ResponseBuilder struct {
	Set     ResponseSet
	Message ResponseMessage
}

type VersionApp struct {
	App   string `json:"app"`
	Build string `json:"build"`
	Date  string `json:"date"`
}

type ResponseApp struct {
	Result interface{} `json:"result"`
	Alias  interface{} `json:"alias"`
}

type MessageApp struct {
	Response   interface{} `json:"response"`
	Properties interface{} `json:"properties"`
}

func (rb *ResponseBuilder) Build(c *gin.Context) map[string]interface{} {
	m := make(map[string]interface{})
	// var httpStatus int

	version_app := GetVersion()

	if rb.Set.Response {
		m["response"] = rb.Message.Response
		m["properties"] = rb.Message.Properties
	}

	if rb.Set.Data {
		if rb.Set.DataShow {
			m["data"] = rb.Message.Data
		}
	}

	if rb.Set.Params {
		for k, v := range rb.Message.Params {
			m[k] = v
		}
	}

	m["version"] = version_app

	return m

}

func (rb *ResponseBuilder) SetResponse(str string) *ResponseBuilder {
	mssg, datashow := GetMessage(str)
	rb.Message.Response = mssg["response"]
	rb.Message.Properties = mssg["properties"]
	rb.Set.Response = true
	rb.Set.DataShow = datashow
	return rb
}

func (rb *ResponseBuilder) SetData(data interface{}) *ResponseBuilder {
	rb.Set.Data = true
	rb.Message.Data = data
	return rb
}

func (rb *ResponseBuilder) SetParams(key string, data interface{}) *ResponseBuilder {
	p := make(map[string]interface{})
	rb.Set.Params = true
	p[key] = data
	rb.Message.Params = p
	return rb
}

func GetVersion() VersionApp {
	var directory = "/Users/andhikarestama/Documents/go/src/mobileloket-app/infrastructure/response"
	var e error

	var version_obj VersionApp
	var version_json []byte

	if version_json, e = ioutil.ReadFile(directory + "/version.json"); e != nil {
		version_json = []byte(`{ "app" : "-", "build" : "-", "date" : "-" }`)
	}

	version_txt := string(version_json)
	json.Unmarshal([]byte(version_txt), &version_obj)

	return version_obj
}

func GetMessage(str string) (map[string]interface{}, bool) {
	var directory = "/Users/andhikarestama/Documents/go/src/mobileloket-app/infrastructure/response"
	var e error

	var mssg_obj ResponseApp
	var mssg_json []byte
	var std_mssg map[string]interface{}
	var err_mssg map[string]interface{}

	dataShow := true

	err_mssg = GetErrorMessage()

	if mssg_json, e = ioutil.ReadFile(directory + "/response.json"); e != nil {
		std_mssg = err_mssg
		dataShow = false
	} else {
		response_txt := string(mssg_json)
		json.Unmarshal([]byte(response_txt), &mssg_obj)

		kode_response, _ := mssg_obj.Result.(map[string]interface{})
		alias_reponse, _ := mssg_obj.Alias.(map[string]interface{})

		if _, ok := kode_response[str]; ok {
			std_mssg, _ = kode_response[str].(map[string]interface{})
		} else {
			if _, ok := alias_reponse[str]; ok {
				str_alias := fmt.Sprintf("%v", alias_reponse[str])
				if _, ok := kode_response[str_alias]; ok {
					std_mssg, _ = kode_response[str_alias].(map[string]interface{})
				} else {
					std_mssg = err_mssg
					dataShow = false
				}
			} else {
				std_mssg = err_mssg
				dataShow = false
			}
		}
	}

	return std_mssg, dataShow
}

func GetErrorMessage() map[string]interface{} {
	var directory = "/Users/andhikarestama/Documents/go/src/mobileloket-app/infrastructure/response"
	var e error

	var err_obj MessageApp
	var err_json []byte
	var err_mssg map[string]interface{}

	if err_json, e = ioutil.ReadFile(directory + "/response-error.json"); e != nil {
		err_app := []byte(`{"response":{"code" : "50000000", "status" : "ERROR", "message" : "Application error"}, "properties":{}}`)
		mssg_txt := string(err_app)

		json.Unmarshal([]byte(mssg_txt), &err_obj)
		err_mssg = map[string]interface{}{
			"response":   err_obj.Response,
			"properties": err_obj.Properties,
		}
	} else {

		mssg_txt := string(err_json)

		json.Unmarshal([]byte(mssg_txt), &err_obj)
		err_mssg = map[string]interface{}{
			"response":   err_obj.Response,
			"properties": err_obj.Properties,
		}
	}

	return err_mssg
}

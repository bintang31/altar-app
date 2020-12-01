package task

import (
	logs "altar-app/infrastructure/logger"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
)

type TerimaTagihan struct {
	Action string `json:"action"`
	Nosamb string `json:"nosamb"`
	Pdam   string `json:"pdam"`
	Userid int    `json:"userid"`
}

//RunTask : Task Assigment
var RunTask = func(payload string) {
	var tagihan TerimaTagihan
	err := json.Unmarshal([]byte(payload), &tagihan)
	if err != nil {
		logs.Error().Msgf("RunTask-error: %s", err)
	}
	if tagihan.Action == "terima_kolektif" {
		spew.Dump("===> Running Task Payment Kolektif")

	}

}

package task

import (
	"altar-app/application/config"
	logs "altar-app/infrastructure/logger"
	"altar-app/infrastructure/persistence"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
)

//TerimaTagihan : Payload Terima Tagihan
type TerimaTagihan struct {
	Action string `json:"action"`
	Nosamb string `json:"nosamb"`
	Pdam   string `json:"pdam"`
	Userid int    `json:"userid"`
}

//RunTask : Task Assigment
var RunTask = func(payload string) {
	var tagihan TerimaTagihan
	conf := config.LoadAppConfig("postgres")
	dbdriver := conf.Driver
	host := conf.Host
	password := conf.Password
	user := conf.User
	dbname := conf.DBName
	port := conf.Port
	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	pelanggans := NewPelanggansTask(services.Pelanggan, services.Penagihan, services.User)
	err = json.Unmarshal([]byte(payload), &tagihan)
	if err != nil {
		logs.Error().Msgf("RunTask-error: %s", err)
	}
	if tagihan.Action == "terima_kolektif" {
		spew.Dump("===> Running Task Payment Kolektif")

	}
	if tagihan.Action == "payment_nosamb" {
		spew.Dump("===> Running Task Payment Nosamb")
		pelanggans.PelanggansTask(tagihan.Nosamb)

	}

}

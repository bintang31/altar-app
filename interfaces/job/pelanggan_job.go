package interfaces

import (
	"altar-app/application"
	"altar-app/application/config"
	"altar-app/domain/entity"
	"altar-app/infrastructure/queue/producer"
	"encoding/json"
	"fmt"
)

//PelangganJobs struct defines the dependencies that will be used
type PelangganJobs struct {
	pl application.PelangganAppInterface
	pn application.PenagihanAppInterface
	us application.UserAppInterface
}

//NewPelangganJobs constructor
func NewPelangganJobs(pl application.PelangganAppInterface, pn application.PenagihanAppInterface, us application.UserAppInterface) *PelangganJobs {
	return &PelangganJobs{
		pl: pl,
		pn: pn,
		us: us,
	}
}

//GetPelangganAlls : Get All Pelanggan
func (p *PelangganJobs) GetPelangganAlls() {
	conf := config.LoadAppConfig("amqp")
	payload := make(map[string]interface{})
	payload["action"] = "terima_kolektif"
	payload["nosamb"] = "08111009"
	payload["pdam"] = "MJI"
	//queue.SendWithParam(param)
	data, _ := json.Marshal(payload)
	producer.Producer.CreateItem(conf.QueueName, string(data))
}

//InquiryLoket : InquiryLoket Pelanggan
func (p *PelangganJobs) InquiryLoket() {
	var input = entity.InputInquiryPelanggan{}
	input.Nosamb = "091310001"
	input.Pdam = "MJI"
	var err error
	tagihanair := entity.RekairDetails{}

	tagihanair, err = p.pl.InquiryLoketTagihanAirByNosamb(&input)

	if err != nil {
		fmt.Printf("userID :%+v\n", err)
		return
	}
	fmt.Printf("userID :%+v\n", tagihanair)
}

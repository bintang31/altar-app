package interfaces

import (
	"altar-app/application"
	"altar-app/application/config"
	"altar-app/domain/entity"
	logger "altar-app/infrastructure/logger"
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

//InquiryLoket : InquiryLoket Pelanggan
func (p *PelangganJobs) InquiryLoket() {
	var err error
	penagihan := entity.PenagihansSrKolektifs{}

	penagihan, err = p.pn.GetPenagihansInquiryProcess()

	conf := config.LoadAppConfig("amqp")
	for _, p := range penagihan {
		payload := make(map[string]interface{})
		payload["action"] = "inquiry_nosamb"
		payload["nosamb"] = p.Nosamb
		payload["pdam"] = p.KodePdam
		//queue.SendWithParam(param)
		data, _ := json.Marshal(payload)
		producer.Producer.CreateItem(conf.QueueName, string(data))
	}

	if err != nil {
		logger.InfoLogHandler("Worker Stopped")
		return
	}
	fmt.Printf("userID :%+v\n", penagihan)
}

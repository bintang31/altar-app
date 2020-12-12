package interfaces

import (
	"altar-app/application"
	"altar-app/application/config"
	"altar-app/infrastructure/queue/producer"
	"encoding/json"
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

package task

import (
	"altar-app/application"
	"altar-app/domain/entity"
	"fmt"
)

//PelanggansTask struct defines the dependencies that will be used
type PelanggansTask struct {
	pl application.PelangganAppInterface
	pn application.PenagihanAppInterface
	us application.UserAppInterface
}

//NewPelanggansTask constructor
func NewPelanggansTask(pl application.PelangganAppInterface, pn application.PenagihanAppInterface, us application.UserAppInterface) *PelanggansTask {
	return &PelanggansTask{
		pl: pl,
		pn: pn,
		us: us,
	}
}

//PelanggansTask : All Pelanggan Task
func (p *PelanggansTask) PelanggansTask(nosamb string) (a map[string]interface{}) {
	data := make(map[string]interface{})
	var err error
	var penagihan *entity.Penagihan
	penagihan, err = p.pn.GetPenagihanByNosamb(nosamb)
	if err != nil {
		fmt.Printf("userID :%+v\n", "Error Log")
		return
	}
	fmt.Printf("userID :%+v\n", penagihan.Nama)
	return data
}

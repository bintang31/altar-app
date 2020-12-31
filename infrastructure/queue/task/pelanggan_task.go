package task

import (
	"altar-app/application"
	"altar-app/domain/entity"
	"fmt"
	"strconv"
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

//InquiryPelanggansTask : Inquiry Pelanggan Task
func (p *PelanggansTask) InquiryPelanggansTask(nosamb string, pdam string) (a map[string]interface{}) {
	data := make(map[string]interface{})
	var err error
	var tokenErr = map[string]string{}
	var postDataTerima = entity.InputInquiryPelanggan{}
	postDataTerima.Nosamb = nosamb
	postDataTerima.Pdam = pdam
	tagihanair := entity.RekairDetails{}
	tagihanair, err = p.pl.InquiryLoketTagihanAirByNosamb(&postDataTerima)
	if err != nil {
		fmt.Printf("tagihanair :%+v\n", "Error Log")
		return
	}
	tagihannonair := entity.NonAirDetails{}
	tagihannonair, err = p.pl.InquiryLoketTagihanNonAirByNosamb(&postDataTerima)
	if err != nil {
		fmt.Printf("tagihannonair :%+v\n", "Error Log")
		return
	}
	angsuran := entity.AngsuranDetails{}
	angsuran, err = p.pl.InquiryLoketAngsuranByNosamb(&postDataTerima)
	if err != nil {
		fmt.Printf("angsuran :%+v\n", "Error Log")
		return
	}
	if len(tagihanair) > 0 {
		for _, sr := range tagihanair {
			var drdUpdate = entity.Drd{}
			drdUpdate.Nosamb = sr.Nosamb
			drdUpdate.Periode = sr.Periode
			drdUpdate.Total = sr.Tagihan
			drdUpdate.TransactionsID = 0
			_, tokenErr = p.pl.UpdateDrdByNosamb(&drdUpdate)
			if tokenErr != nil {
				fmt.Printf("userID :%+v\n", "Error Log")
				return
			}
		}
	}

	if len(tagihannonair) > 0 {
		fmt.Printf("tagihannonair :%+v\n", tagihannonair)
	}

	if len(angsuran) > 0 {
		for _, ar := range angsuran {

			var nonairUpdate = entity.Nonair{}
			year := ar.Periode[0:4]

			month := ar.Periode[4:6]

			bulan := bulanConvert(month) + " " + year

			nonairUpdate.Urutan = ar.Urutan
			nonairUpdate.Nomor = ar.Nomor
			nonairUpdate.Pdam = pdam
			nonairUpdate.Periode = ar.Periode
			nonairUpdate.Lunas = "0"
			nonairUpdate.Jenis = ar.Jenis
			nonairUpdate.DendaTunggakan = ar.DendaTunggakan
			nonairUpdate.Angsur = strconv.Itoa(ar.FlagAngsur)
			nonairUpdate.Flag = ar.Flag
			nonairUpdate.Bulan = bulan
			nonairUpdate.NoAngsuran = ar.Noangsuran
			nonairUpdate.Termin = ar.Termin
			nonairUpdate.KetJenis = ar.KetJenis
			nonairUpdate.DibebankanKepada = ar.DibebankanKepada
			nonairUpdate.DendaTunggakan = ar.DendaTunggakan
			nonairUpdate.Administrasi = ar.Administrasi
			nonairUpdate.Total = ar.Total
			nonairUpdate.JenisTagihan = ar.JenisKeterangan
			_, tokenErr = p.pl.InsertAngsuranByNosamb(&nonairUpdate)
			if tokenErr != nil {
				fmt.Printf("userID :%+v\n", tokenErr)
				return
			}
		}
	}

	return data
}

func bulanConvert(month string) (a string) {
	bulan := ""
	switch month {
	case "01":
		bulan = "Januari"
	case "02":
		bulan = "Februari"
	case "03":
		bulan = "Maret"
	case "04":
		bulan = "April"
	case "05":
		bulan = "Mei"
	case "06":
		bulan = "Juni"
	case "07":
		bulan = "Juli"
	case "08":
		bulan = "Agustus"
	case "09":
		bulan = "September"
	case "10":
		bulan = "Oktober"
	case "11":
		bulan = "November"
	case "12":
		bulan = "Desember"
	default:
		bulan = "Januari"
	}
	return bulan
}

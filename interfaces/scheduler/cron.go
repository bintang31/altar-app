package interfaces

import (
	"altar-app/application/config"
	"altar-app/infrastructure/persistence"
	"altar-app/interfaces/job"
	"gopkg.in/robfig/cron.v2"
)

// InitCronInfo : CronJOB Routes to function backend
func InitCronInfo() {
	c := cron.New()
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
	pelanggans := interfaces.NewPelangganJobs(services.Pelanggan, services.Penagihan, services.User)
	_, _ = c.AddFunc("2 * * * * *", pelanggans.InquiryLoket)
	//_, _ = c.AddFunc("15 * * * * *", job.CronExportData)
	//_, _ = c.AddFunc("15 * * * * *", job.CronUploadData)
	//_, _ = c.AddFunc("5 * * * *", func() { fmt.Println("Every 5 minutes" + time.Now().Format("2006-01-02T00:00:00+07:00")) })
	//_, _ = c.AddFunc("* * * * *", func() { fmt.Println("Every 1 minutes" + time.Now().Format("2006-01-02T00:00:00+07:00")) })
	c.Start()

}

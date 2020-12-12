package routes

import (
	"altar-app/application/config"
	"altar-app/infrastructure/auth"
	"altar-app/infrastructure/persistence"
	"altar-app/interfaces/handler"
	"altar-app/interfaces/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//API : Handler API Interfacing to FrontEnd
func API() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	conf := config.LoadAppConfig("postgres")
	dbdriver := conf.Driver
	host := conf.Host
	password := conf.Password
	user := conf.User
	dbname := conf.DBName
	port := conf.Port

	//redis details
	confRedis := config.LoadAppConfig("redis")
	redisHost := confRedis.Host
	redisPort := confRedis.Port
	redisPassword := confRedis.Password

	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.Automigrate()

	redisService, err := auth.NewRedisDB(redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()

	users := interfaces.NewUsers(services.User, redisService.Auth, tk)
	roles := interfaces.NewRoles(services.Role, redisService.Auth, tk)
	pelanggans := interfaces.NewPelanggans(services.Pelanggan, services.Penagihan, services.User, redisService.Auth, tk)
	penagihans := interfaces.NewPenagihans(services.Penagihan, services.User, redisService.Auth, tk)
	petugas := interfaces.NewPetugass(services.Petugas, services.Penagihan, services.User, redisService.Auth, tk)
	transactions := interfaces.NewTransactions(services.Transaksi, services.User, redisService.Auth, tk)
	authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS
	v1 := r.Group("/v1/api")
	//user routes
	v1.POST("/users", middleware.AuthMiddleware(), users.SaveUser)
	v1.GET("/users", users.GetUsers)
	v1.GET("/users/:user_id", users.GetUser)

	//role routes
	v1.GET("/roles", middleware.AuthMiddleware(), roles.GetRoles)

	//pelanggans routes
	v1.GET("/pelanggans", middleware.AuthMiddleware(), pelanggans.GetPelanggans)
	v1.GET("/pelanggans/:nosamb/tagihan", middleware.AuthMiddleware(), pelanggans.GetTagihanPelanggan)

	//penagihans routes
	v1.GET("/penagihans", middleware.AuthMiddleware(), penagihans.GetPenagihans)
	v1.POST("/bayar_sync", middleware.AuthMiddleware(), penagihans.BayarTagihanPelangganBulk)
	v1.POST("/bayar", middleware.AuthMiddleware(), penagihans.BayarTagihanPelanggan)

	//transactions routes
	v1.POST("/transactions", middleware.AuthMiddleware(), transactions.GetTransactions)

	//petugas routes
	v1.GET("/petugas_data", middleware.AuthMiddleware(), petugas.GetDataPetugas)
	v1.GET("/petugas/profile", middleware.AuthMiddleware(), petugas.GetProfilePetugas)

	//authentication routes
	v1.POST("/login", authenticate.Login)
	v1.POST("/logout", authenticate.Logout)
	v1.POST("/refresh", authenticate.Refresh)

	//Starting the application
	appPort := os.Getenv("PORT") //using heroku host
	if appPort == "" {
		appPort = "1123" //localhost
	}
	log.Fatal(r.Run(":" + appPort))
}

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
	pelanggans := interfaces.NewPelanggans(services.Pelanggan, services.User, redisService.Auth, tk)
	penagihans := interfaces.NewPenagihans(services.Penagihan, services.User, redisService.Auth, tk)
	authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	//user routes
	r.POST("/users", middleware.AuthMiddleware(), users.SaveUser)
	r.GET("/users", users.GetUsers)
	r.GET("/users/:user_id", users.GetUser)

	//role routes
	r.GET("/roles", middleware.AuthMiddleware(), roles.GetRoles)

	//pelanggans routes
	r.GET("/pelanggans", middleware.AuthMiddleware(), pelanggans.GetPelanggans)

	//penagihans routes
	r.GET("/penagihans", middleware.AuthMiddleware(), penagihans.GetPenagihans)

	//authentication routes
	r.POST("/login", authenticate.Login)
	r.POST("/logout", authenticate.Logout)
	r.POST("/refresh", authenticate.Refresh)

	//Starting the application
	appPort := os.Getenv("PORT") //using heroku host
	if appPort == "" {
		appPort = "8888" //localhost
	}
	log.Fatal(r.Run(":" + appPort))
}

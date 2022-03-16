package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deepsourcelabs/apollo/cmd/apollo/routes"
	dc "github.com/deepsourcelabs/apollo/controllers/dependency"
	hc "github.com/deepsourcelabs/apollo/controllers/health"
	healthDb "github.com/deepsourcelabs/apollo/drivers/db/health/redis"
	"github.com/deepsourcelabs/apollo/drivers/redis"
	"github.com/deepsourcelabs/apollo/entity/health"
	"github.com/labstack/echo/v4"
	flag "github.com/spf13/pflag"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/deepsourcelabs/apollo/docs"
	"github.com/deepsourcelabs/apollo/logging"
	"github.com/deepsourcelabs/apollo/utils"
)

func main() {
	// set up echo
	e := echo.New()

	// expose Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// set up flags
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.String("path", "", "path to config file")
	err := f.Parse(os.Args[1:])
	if err != nil {
		log.Println("failed to parse args")
		f.Usage()
	}

	filePath, err := f.GetString("path")
	if err != nil {
		log.Println("failed to get config path")
		f.Usage()
	}

	if filePath == "" {
		log.Println("no value for config path provided")
		f.Usage()
	}

	// get configs from config file
	conf, err := utils.GetConfig(filePath)
	if err != nil {
		log.Fatalln("failed to load config:", err)
	}

	// set up logger
	logger := logging.NewLogger(conf.Logging.File)

	// get DB connection
	Conn := redis.GetConn(conf)

	// get timeout duration from config
	timeout := utils.GetTimeout(conf)

	// set up a new health use case
	healthUsecase := health.NewUseCase(healthDb.NewHealthRepo(Conn, logger), timeout)

	// set up controllers
	healthController := hc.NewHealthController(*healthUsecase)
	dependencyController := dc.NewDependencyController(*healthUsecase)
	rc := routes.Controllers{
		DependencyController: dependencyController,
		HealthController:     healthController,
	}
	rc.InitRoutes(e)

	// warn user if server address is empty
	if conf.Server.Addr == "" {
		log.Println("[WARNING] failed to get server address, using random address.")
	}

	// start server
	log.Println(e.Start(conf.Server.Addr))
}

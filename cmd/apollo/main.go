package main

import (
	"fmt"
	"log"
	"os"

	"github.com/burntcarrot/apollo/cmd/apollo/routes"
	dc "github.com/burntcarrot/apollo/controllers/dependency"
	hc "github.com/burntcarrot/apollo/controllers/health"
	healthDb "github.com/burntcarrot/apollo/drivers/db/health/redis"
	"github.com/burntcarrot/apollo/drivers/redis"
	"github.com/burntcarrot/apollo/entity/health"
	"github.com/labstack/echo/v4"
	flag "github.com/spf13/pflag"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/burntcarrot/apollo/docs"
	"github.com/burntcarrot/apollo/logging"
	"github.com/burntcarrot/apollo/utils"
)

func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.String("path", "", "path to config file")
	f.Parse(os.Args[1:])

	filePath, err := f.GetString("path")
	if err != nil {
		log.Println("failed to get config path")
		f.Usage()
	}

	if filePath == "" {
		log.Println("no value for config path provided")
		f.Usage()
	}

	conf := utils.GetConfig(filePath)

	logger := logging.NewLogger(conf.Logging.File)

	Conn := redis.GetConn(conf)
	timeout := utils.GetTimeout(conf)

	healthUsecase := health.NewUseCase(healthDb.NewHealthRepo(Conn, logger), timeout)

	healthController := hc.NewHealthController(*healthUsecase)
	dependencyController := dc.NewDependencyController(*healthUsecase)
	rc := routes.Controllers{
		DependencyController: dependencyController,
		HealthController:     healthController,
	}
	rc.InitRoutes(e)

	if conf.Server.Addr == "" {
		log.Println("failed to get server address, using random address.")
	}

	log.Println(e.Start(conf.Server.Addr))
}

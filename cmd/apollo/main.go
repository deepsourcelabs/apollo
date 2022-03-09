package main

import (
	"log"
	"time"

	"github.com/burntcarrot/apollo/cmd/apollo/routes"
	dc "github.com/burntcarrot/apollo/controllers/dependency"
	hc "github.com/burntcarrot/apollo/controllers/health"
	healthDbRedis "github.com/burntcarrot/apollo/drivers/db/health/redis"
	"github.com/burntcarrot/apollo/drivers/redis"
	"github.com/burntcarrot/apollo/entity/health"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/burntcarrot/apollo/docs"
)

func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	dbConfig := redis.DBConfig{
		Addr: "localhost:6379",
	}
	Conn := dbConfig.InitDB()
	timeout := time.Duration(time.Minute * 5)

	healthUsecase := health.NewUseCase(healthDbRedis.NewHealthRepo(Conn), timeout)

	healthController := hc.NewHealthController(*healthUsecase)
	dependencyController := dc.NewDependencyController(*healthUsecase)
	rc := routes.Controllers{
		DependencyController: dependencyController,
		HealthController:     healthController,
	}
	rc.InitRoutes(e)
	log.Println(e.Start(":8080"))
}

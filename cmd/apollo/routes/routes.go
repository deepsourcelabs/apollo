package routes

import (
	"github.com/burntcarrot/apollo/controllers/dependency"
	"github.com/burntcarrot/apollo/controllers/health"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controllers struct {
	HealthController     *health.HealthController
	DependencyController *dependency.DependencyController
}

func (c *Controllers) InitRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.Use(middleware.Recover())

	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// routes
	api.POST("/register", c.DependencyController.Register)
	api.GET("/health", c.HealthController.GetPrimaryHealthCheck)
	api.GET("/health/:id", c.HealthController.GetHealthCheck)
}

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

// @title Apollo
// @version 1.0
// @description A distributed deep health check system
// @contact.name DeepSource
// @contact.url https://deepsource.io/
// @BasePath /api/v1
func (c *Controllers) InitRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
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
	api.GET("/health", c.HealthController.GetHealthCheck)
	api.GET("/health/:id", c.HealthController.GetHealthCheckByID)
}

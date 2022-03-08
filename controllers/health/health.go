package health

import (
	"net/http"
	"strconv"

	"github.com/burntcarrot/apollo/controllers"
	"github.com/burntcarrot/apollo/entity/health"
	"github.com/labstack/echo/v4"
)

type HealthController struct {
	Usecase health.Usecase
}

func NewHealthController(u health.Usecase) *HealthController {
	return &HealthController{
		Usecase: u,
	}
}

func (h *HealthController) GetPrimaryHealthCheck(c echo.Context) error {
	ctx := c.Request().Context()

	services, err := h.Usecase.GetAllServices(ctx)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError)
	}

	results, err := getResults(services)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError)
	}

	resp := health.Response{
		Results: results,
	}

	return controllers.Success(c, resp)
}

func (h *HealthController) GetHealthCheck(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError)
	}

	services, err := h.Usecase.GetServices(ctx, uint(id))
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError)
	}

	results, err := getResults(services)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError)
	}

	resp := health.Response{
		Results: results,
	}

	return controllers.Success(c, resp)
}

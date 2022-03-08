package dependency

import (
	"net/http"

	"github.com/burntcarrot/apollo/controllers"
	"github.com/burntcarrot/apollo/entity/health"
	"github.com/labstack/echo/v4"
)

type DependencyController struct {
	Usecase health.Usecase
}

func NewDependencyController(u health.Usecase) *DependencyController {
	return &DependencyController{
		Usecase: u,
	}
}

func (d *DependencyController) Register(c echo.Context) error {
	req := RegisterRequest{}
	c.Bind(&req)

	ctx := c.Request().Context()

	srv, err := d.Usecase.GetServices(ctx, req.ID)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError)
	}

	if len(srv) != 0 {
		return controllers.Error(c, http.StatusBadRequest)
	}

	healthDomain := health.Domain{
		Name: req.Name,
		URI:  req.URI,
	}

	_, err = d.Usecase.CreateService(ctx, healthDomain, req.ID)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError)
	}

	return controllers.Success(c, "woohoo")
}

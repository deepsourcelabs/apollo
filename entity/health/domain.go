package health

import "context"

type Domain struct {
	Name string
	URI  string
}

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Err   string `json:"err"`
}

type DomainRepo interface {
	CreateService(ctx context.Context, hd Domain, id uint) (Domain, error)
	GetServicesByID(ctx context.Context, id uint) ([]Domain, error)
	GetAllServices(ctx context.Context) ([]Domain, error)
}

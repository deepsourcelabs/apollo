package health

import "context"

type Domain struct {
	Name string
	URI  string
}

type DomainRepo interface {
	CreateService(ctx context.Context, hd Domain, id uint) (Domain, error)
	GetServicesByID(ctx context.Context, id uint) ([]Domain, error)
	GetAllServices(ctx context.Context) ([]Domain, error)
}

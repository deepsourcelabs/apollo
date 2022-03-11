package health

import "context"

type Domain struct {
	Name string
	URI  string
}

type DomainRepo interface {
	CreateService(ctx context.Context, domain Domain, id string) (Domain, error)
	GetServicesByID(ctx context.Context, id string) ([]Domain, error)
	GetAllServices(ctx context.Context) ([]Domain, error)
}

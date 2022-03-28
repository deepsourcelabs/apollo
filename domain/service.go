package domain

import "context"

// Service represents the details of the service registered to the Apollo server.
type Service struct {
	Name string
	URI  string
}

type ServiceRepo interface {
	CreateService(ctx context.Context, srv Service, id string) (Service, error)
	GetServicesByID(ctx context.Context, id string) ([]Service, error)
	GetAllServices(ctx context.Context) ([]Service, error)
}

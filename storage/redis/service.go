package storage

import (
	"context"

	"github.com/deepsourcelabs/apollo/domain"
	"github.com/go-redis/redis/v8"
)

// Service represents the details of the service registered to the Apollo server.
type Service struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

type serviceStore struct {
	db *redis.Client
}

func NewServiceStore(conn *redis.Client) domain.ServiceRepo {
	return &serviceStore{
		db: conn,
	}
}

func (store *serviceStore) CreateService(ctx context.Context, srv domain.Service, id string) (domain.Service, error) {
	return domain.Service{}, nil
}

func (store *serviceStore) GetAllServices(ctx context.Context) ([]domain.Service, error) {
	return []domain.Service{}, nil
}

func (store *serviceStore) GetServicesByID(ctx context.Context, id string) ([]domain.Service, error) {
	return []domain.Service{}, nil
}

package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	dbHealth "github.com/burntcarrot/apollo/drivers/db/health"
	"github.com/burntcarrot/apollo/entity/health"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// MAX_FETCH_ROWS is a constant denoting the maximum amount of records to be fetched.
const MAX_FETCH_ROWS = 9 * 100000

type HealthRepo struct {
	Conn   *redis.Client
	logger *zap.SugaredLogger
}

func NewHealthRepo(conn *redis.Client, logger *zap.SugaredLogger) *HealthRepo {
	return &HealthRepo{
		Conn:   conn,
		logger: logger,
	}
}

func (h *HealthRepo) CreateService(ctx context.Context, srv health.Domain, id string) (health.Domain, error) {
	createdService := dbHealth.Service{
		Name: srv.Name,
		URI:  srv.URI,
	}

	serviceData, err := json.Marshal(createdService)
	if err != nil {
		h.logger.Errorf("[health/redis/CreateService] failed to marshal: %v\n", err)
		return health.Domain{}, errors.New("failed to marshal")
	}

	key := fmt.Sprintf("apollo:%s:services", id)
	err = h.Conn.RPush(ctx, key, serviceData).Err()
	if err != nil {
		h.logger.Errorf("[health/redis/CreateService] failed to push to %s: %v\n", key, err)
		return health.Domain{}, errors.New("failed to create service")
	}

	// push service data
	srvKey := "apollo:services"
	err = h.Conn.RPush(ctx, srvKey, serviceData).Err()
	if err != nil {
		h.logger.Errorf("[health/redis/CreateService] failed to push to %s: %v\n", srvKey, err)
		return health.Domain{}, errors.New("failed to create service")
	}

	return createdService.ToDomain(), nil
}

func (h *HealthRepo) GetServicesByID(ctx context.Context, id string) ([]health.Domain, error) {
	key := fmt.Sprintf("apollo:%s:services", id)
	serviceData, err := h.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		h.logger.Errorf("[health/redis/GetServicesByID] failed to fetch %s: %v\n", key, err)
		return []health.Domain{}, errors.New("failed to fetch services")
	}

	srv := new(dbHealth.Service)
	var services []health.Domain

	for _, service := range serviceData {
		if err := json.Unmarshal([]byte(service), srv); err != nil {
			h.logger.Errorf("[health/redis/GetServicesByID] failed to unmarshal: %v\n", err)
			return []health.Domain{}, errors.New("failed to unmarshal")
		}

		services = append(services, srv.ToDomain())
	}

	return services, nil
}

func (h *HealthRepo) GetAllServices(ctx context.Context) ([]health.Domain, error) {
	key := "apollo:services"
	serviceData, err := h.Conn.LRange(ctx, key, 0, MAX_FETCH_ROWS).Result()
	if err != nil {
		h.logger.Errorf("[health/redis/GetAllServices] failed to fetch %s: %v\n", key, err)
		return []health.Domain{}, errors.New("failed to fetch services")
	}

	srv := new(dbHealth.Service)
	var services []health.Domain

	for _, service := range serviceData {
		if err := json.Unmarshal([]byte(service), srv); err != nil {
			h.logger.Errorf("[health/redis/GetAllServices] failed to unmarshal: %v\n", err)
			return []health.Domain{}, errors.New("failed to unmarshal")
		}

		services = append(services, srv.ToDomain())
	}

	return services, nil
}

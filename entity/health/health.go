package health

import (
	"context"
	"time"
)

type Usecase struct {
	Repo       DomainRepo
	ctxTimeout time.Duration
}

func NewUseCase(repo DomainRepo, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:       repo,
		ctxTimeout: timeout,
	}
}

func (u *Usecase) CreateService(ctx context.Context, hd Domain, id uint) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.Repo.CreateService(ctx, hd, id)
}

func (u *Usecase) GetServices(ctx context.Context, id uint) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.Repo.GetServicesByID(ctx, id)
}

func (u *Usecase) GetAllServices(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.Repo.GetAllServices(ctx)
}

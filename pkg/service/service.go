package service

import (
	"context"

	"github.com/popeskul/audit-logger/pkg/domain"
	"github.com/popeskul/audit-logger/pkg/repository"
)

type Logger interface {
	Insert(ctx context.Context, item domain.LogItem) error
}

type LogService struct {
	repo *repository.Repository
}

func NewLogService(repo *repository.Repository) *LogService {
	return &LogService{
		repo: repo,
	}
}

func (s *LogService) Log(ctx context.Context, req *domain.LogRequest) error {
	item := domain.LogItem{
		Action:    req.Action.String(),
		Entity:    req.Entity.String(),
		EntityID:  req.EntityId,
		Timestamp: req.Timestamp.AsTime(),
	}

	return s.repo.Insert(ctx, item)
}

package handler

import (
	"context"
	"github.com/popeskul/audit-logger/pkg/domain"
	"github.com/popeskul/audit-logger/pkg/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.LogService
	domain.UnimplementedAuditServiceServer
}

func NewHandler(services *service.LogService) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Log(ctx context.Context, req *domain.LogRequest) (*domain.Empty, error) {
	err := h.services.Log(ctx, req)
	if err != nil {
		logrus.Errorf("error while logging: %v", err)
		return &domain.Empty{}, err
	}

	logrus.Infof("log request: %v", req)

	return &domain.Empty{}, nil
}

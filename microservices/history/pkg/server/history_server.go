package server

import (
	"context"
	"github.com/Askalag/aska/microservices/history/pkg/service"
	status_v1 "github.com/Askalag/aska/microservices/history/proto/status/v1"
)

type HistoryServer struct {
	hs service.History
}

func NewHistoryServer(s service.History) *HistoryServer {
	return &HistoryServer{hs: s}
}

func (s *HistoryServer) Status(_ context.Context, _ *status_v1.StatusRequest) (*status_v1.StatusResponse, error) {
	st, err := s.hs.Status()
	if err != nil {
		return nil, err
	}
	status := &status_v1.StatusResponse{Status: st}
	return status, nil
}

package service

import (
	"context"
	status_v1 "github.com/Askalag/protolib/gen/proto/go/status/v1"
	"github.com/askalag/aska/microservices/webapp/pkg"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HistoryService struct {
	tcp      pkg.ServicesTCP
	conn     *grpc.ClientConn
	clientHS status_v1.StatusServiceClient
}

func NewHistoryService(tcp pkg.ServicesTCP) *HistoryService {
	conn, err := grpc.Dial(tcp.HistoryAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(err)
	}

	clientHS := status_v1.NewStatusServiceClient(conn)

	return &HistoryService{
		tcp:      tcp,
		conn:     conn,
		clientHS: clientHS,
	}
}

func (s *HistoryService) GrpcStatus(ctx context.Context, req *status_v1.StatusRequest) (*status_v1.StatusResponse, error) {
	status, err := s.clientHS.Status(ctx, req)
	if err != nil {
		return nil, err
	}
	return status, nil
}

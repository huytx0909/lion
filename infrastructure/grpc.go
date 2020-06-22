package infrastructure

import (
	"context"
	"errors"
	lion_service "lion/proto/proto"
	"lion/usecase"
)

type LionServiceServer struct {
	useCase usecase.LionUseCaseInterface
}

func (l LionServiceServer) FindPrey(ctx context.Context, request *lion_service.FindPreyRequest) (*lion_service.FindPreyResponse, error) {
	var preyName string
	if request.GetPreyName() == "" {
		return nil, errors.New("empty prey name")
	}
	preyName = request.GetPreyName()
	return &lion_service.FindPreyResponse{PreyName: preyName}, nil
}

func NewLionServiceServer(lionUC usecase.LionUseCaseInterface) *LionServiceServer {
	return &LionServiceServer{useCase: lionUC}
}

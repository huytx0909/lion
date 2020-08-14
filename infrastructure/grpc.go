package infrastructure

import (
	"context"
	"errors"
	"fmt"
	lion_service "lion/proto/proto"
	"lion/usecase"
)

type LionServiceServer struct {
	useCase usecase.LionUseCaseInterface
}

func (l LionServiceServer) AddPrey(ctx context.Context, request *lion_service.AddPreyRequest) (*lion_service.AddPreyResponse, error) {
	if request.PreyName != "" {
		return &lion_service.AddPreyResponse{
			PreyName: request.PreyName+" added",
		}, nil
	}
	return nil, errors.New("empty prey name")
}

func (l LionServiceServer) GetPrey(ctx context.Context, request *lion_service.GetPreyRequest) (*lion_service.GetPreyResponse, error) {
	fmt.Print("func GetPrey")
	if request.GetId() == "" {
		return nil, errors.New("empty id")
	}
	preyName := "prey number" + request.Id
	return &lion_service.GetPreyResponse{PreyName: preyName}, nil
}

func NewLionServiceServer(lionUC usecase.LionUseCaseInterface) *LionServiceServer {
	return &LionServiceServer{useCase: lionUC}
}

package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"lion/infrastructure"
	"lion/usecase"
	"net"

	lion_service "lion/proto/proto"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "run http server",
	Run:   runGrpc,
}

func runGrpc(command *cobra.Command, args []string) {
	grpcServer := grpc.NewServer()

	// UseCase
	uc := usecase.NewLionUseCase()

	lis, err := net.Listen("tcp", "localhost:5678")
	if err != nil {
		panic(err)
	}

	lionServer := infrastructure.NewLionServiceServer(uc)

	lion_service.RegisterLionServiceServer(grpcServer, lionServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

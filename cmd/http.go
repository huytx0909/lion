package cmd

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	lion_service "lion/proto/proto"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "run http server",
	Run:   runHttp,
}

func runHttp(command *cobra.Command, args []string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := lion_service.RegisterLionServiceHandlerFromEndpoint(ctx, mux, "localhost:5678", opts)
	if err != nil {
		panic(err)
	}
	handler := cors.AllowAll().Handler(mux)
	fmt.Println("listening...")
	err = http.ListenAndServe(":1234", handler)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/EFinish/leibniz/be/services/argument-manager/config"
	logger "github.com/EFinish/leibniz/be/utilities/logger"
	protoOut "github.com/EFinish/leibniz/proto/gen/argumentmanager/v1"
)

type (
	argumentManager struct {
		conf   *config.Config
		logger logger.LeibnizLogger
	}
	ArgumentManagerServiceServer struct {
		protoOut.UnimplementedArgumentManagerServer
	}
)

func NewServer() *ArgumentManagerServiceServer {
	return &ArgumentManagerServiceServer{}
}

var aa *argumentManager

func main() {
	conf := config.GetConfig()
	logger := logger.InitLogger("argument-manager")
	aa = &argumentManager{
		logger: logger,
		conf:   conf,
	}

	fmt.Printf("Listening on port %v\n", conf.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.GrpcPort))

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	protoOut.RegisterArgumentManagerServer(grpcServer, &ArgumentManagerServiceServer{})
	aa.logger.Infof("Serving gRPC on port %v\n", conf.GrpcPort)

	go func() {
		aa.logger.Fatalf("failed to serve grpc server: %w", grpcServer.Serve(lis))
	}()

	// create apiGW server
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%v", conf.GrpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		aa.logger.Fatalf("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	err = protoOut.RegisterArgumentManagerHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", "9000"),
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:9000")

	aa.logger.Fatalf("failed to listen and serve API GW server: %w", gwServer.ListenAndServe())
}

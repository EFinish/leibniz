package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"

	"github.com/EFinish/leibniz/be/services/argument-access/config"
	logger "github.com/EFinish/leibniz/be/utilities/logger"
	protoOut "github.com/EFinish/leibniz/proto/gen/argumentaccess/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type (
	argumentAccess struct {
		conf                            *config.Config
		logger                          logger.LeibnizLogger
		premisesCollection              *mongo.Collection
		subjectsCollection              *mongo.Collection
		predicatesCollection            *mongo.Collection
		propositionsCollection          *mongo.Collection
		conditionalStatementsCollection *mongo.Collection
		argumentsCollection             *mongo.Collection
	}
	ArgumentAccessServiceServer struct {
		protoOut.UnimplementedArgumentAccessServer
	}
)

var aa *argumentAccess

func main() {
	conf := config.GetConfig()
	logger := logger.InitLogger("argument-access")
	aa = &argumentAccess{
		logger: logger,
		conf:   conf,
	}

	bgContext := context.Background()

	ctx, cancel := context.WithTimeout(bgContext, 10*time.Second)
	defer cancel()

	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(aa.conf.ArgumentDbURL))
	if err != nil {
		fmt.Printf("unable to connect to argument database: %v", err)

		panic(err)
	}

	ctx, cancel = context.WithTimeout(bgContext, 2*time.Second)
	defer cancel()

	aa.logger.Infof("pinging argument database")

	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("unable to ping to argument database: %v", err)

		panic(err)
	}

	aa.logger.Infof("connected to mongo DB for argument database")

	aa.premisesCollection = dbClient.Database("argument").Collection("premises")
	aa.subjectsCollection = dbClient.Database("argument").Collection("subjects")
	aa.predicatesCollection = dbClient.Database("argument").Collection("predicates")
	aa.propositionsCollection = dbClient.Database("argument").Collection("propositions")
	aa.conditionalStatementsCollection = dbClient.Database("argument").Collection("conditional_statements")
	aa.argumentsCollection = dbClient.Database("argument").Collection("arguments")

	fmt.Printf("Listening on port %v\n", conf.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.GrpcPort))

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	protoOut.RegisterArgumentAccessServer(grpcServer, &ArgumentAccessServiceServer{})
	aa.logger.Infof("Serving gRPC on port %v\n", conf.GrpcPort)

	go func() {
		aa.logger.Fatalf("failed to serve grpc server: %w", grpcServer.Serve(lis))
	}()

	// create apiGW server
	// conn, err := grpc.DialContext(
	// 	context.Background(),
	// 	fmt.Sprintf("0.0.0.0:%v", conf.GrpcPort),
	// 	grpc.WithBlock(),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )

	gwmux := runtime.NewServeMux()
	// err = protoOut.RegisterBuildingdataAccessHandler(context.Background(), gwmux, conn)
	// if err != nil {
	// 	aa.logger.Fatalf("Failed to register gateway: %w", err)
	// }

	aa.logger.Infof("API GW server serving on port %s", aa.conf.GrpcPort)

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", aa.conf.GrpcPort),
		Handler: gwmux,
	}

	aa.logger.Fatalf("failed to listen and serve API GW server: %w", gwServer.ListenAndServe())
}

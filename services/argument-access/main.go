package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"
	logger "github.com/EFinish/leibniz/utilities/logger"
)

type (
	argumentAccess struct {
		logger               logger.LeibnizLogger
		argumentsCollection  *mongo.Collection
		premisesCollection   *mongo.Collection
		subjectsCollection   *mongo.Collection
		predicatesCollection *mongo.Collection
	}
	ArgumentAccessServiceServer struct {
		protoOut.UnimplementedArgumentAccessServer
	}
)

var aa *argumentAccess

func main() {
	logger := logger.InitLogger("argument-access")
	aa = &argumentAccess{
		logger: logger,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	argumentDBClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo-db:27017/argument"))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = argumentDBClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = argumentDBClient.Ping(ctx, readpref.Primary())

	if err != nil {
		panic(err)
	}

	fmt.Println("connected to mongo DB for argument database")

	aa.argumentsCollection = argumentDBClient.Database("argument").Collection("arguments")
	aa.premisesCollection = argumentDBClient.Database("argument").Collection("premises")
	aa.subjectsCollection = argumentDBClient.Database("argument").Collection("subjects")
	aa.predicatesCollection = argumentDBClient.Database("argument").Collection("predicates")

	fmt.Println("Listening on port 9002")

	lis, err := net.Listen("tcp", ":9002")

	if err != nil {
		panic(err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	argumentAccessServer := ArgumentAccessServiceServer{}
	protoOut.RegisterArgumentAccessServer(grpcServer, &argumentAccessServer)

	aa.logger.Infof("gRPC server starting")
	if err := grpcServer.Serve(lis); err != nil {
		aa.logger.Fatalf("failed to start gRPC server: %v", err)
	}
}

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
	"github.com/EFinish/leibniz/services/argument-access/config"
	logger "github.com/EFinish/leibniz/utilities/logger"
)

type (
	argumentAccess struct {
		conf                 *config.Config
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
	// dbDisconnect := func() {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// 	defer cancel()
	// 	if err := dbClient.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }

	ctx, cancel = context.WithTimeout(bgContext, 2*time.Second)
	defer cancel()
	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("unable to ping to argument database: %v", err)

		panic(err)
	}

	fmt.Println("connected to mongo DB for argument database")

	aa.argumentsCollection = dbClient.Database("argument").Collection("arguments")
	aa.premisesCollection = dbClient.Database("argument").Collection("premises")
	aa.subjectsCollection = dbClient.Database("argument").Collection("subjects")
	aa.predicatesCollection = dbClient.Database("argument").Collection("predicates")

	fmt.Printf("Listening on port %v\n", conf.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.GrpcPort))

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

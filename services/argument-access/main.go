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
)

type ArgumentAccess struct {
	argumentsCollection  *mongo.Collection
	premisesCollection   *mongo.Collection
	subjectsCollection   *mongo.Collection
	predicatesCollection *mongo.Collection
}

var aa *ArgumentAccess

func main() {
	aa = &ArgumentAccess{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	argumentDBClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo-db:27017/workout"))

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

	// workoutAccessServer := WorkoutAccessServiceServer{}
	// protoOut.RegisterWorkoutAccessServer(grpcServer, &workoutAccessServer)

	// wa.isReadyForRequests = true

	// wa.logger.Infof("gRPC server starting")
	// if err := grpcServer.Serve(lis); err != nil {
	// 	wa.logger.Fatalf("failed to start gRPC server: %v", err)
	// }
}

package integration_tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/mgo.v2/bson"

	protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"
)

var (
	argumentAccessClient protoOut.ArgumentAccessClient
	argumentsCollection  *mongo.Collection
	premisesCollection   *mongo.Collection
	subjectsCollection   *mongo.Collection
	predicatesCollection *mongo.Collection
)

func getEnv(env string, defaultVal string) string {
	value, present := os.LookupEnv(env)
	if !present {
		return defaultVal
	}
	return value
}

func TestMain(m *testing.M) {
	initializeArgumentAccessClient()
	initializeMongoDBConnections()

	dbCleanUp(context.Background())

	// Run tests
	exitCode := m.Run()
	os.Exit(exitCode)
}

func initializeArgumentAccessClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())
	grpcHost := getEnv("ARGUMENT_ACCESS_TEST_GRPC_HOST", "localhost")
	grpcPort := getEnv("ARGUMENT_ACCESS_TEST_GRPC_PORT", "9002")
	serverAddr := fmt.Sprintf("%s:%s", grpcHost, grpcPort)
	conn, err := grpc.DialContext(ctx, serverAddr, opts...)

	if err != nil {
		panic(fmt.Errorf("creating argument access test client: dialing grpc: %w", err))
	}

	argumentAccessClient = protoOut.NewArgumentAccessClient(conn)

	defer conn.Close()
}

func initializeMongoDBConnections() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	argumentDBUrl := getEnv("ARGUMENT_ACCESS_TEST_ARGUMENT_DB_URL", "mongodb://localhost:27017/argument")
	argumentDbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(argumentDBUrl))

	if err != nil {
		panic(fmt.Errorf("creating argument DB client: %w", err))
	}
	defer func() {
		if err := argumentDbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = pingMongoDBClient(argumentDbClient)

	if err != nil {
		panic(fmt.Errorf("testing argument DB client: %w", err))
	}

	subjectsCollection = argumentDbClient.Database("argument").Collection("subjects")
	predicatesCollection = argumentDbClient.Database("argument").Collection("predicates")
	premisesCollection = argumentDbClient.Database("argument").Collection("premises")
	argumentsCollection = argumentDbClient.Database("argument").Collection("arguments")
}

func pingMongoDBClient(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := client.Ping(ctx, readpref.Primary())

	if err != nil {
		return fmt.Errorf("while pinging via mongoDB client: %w", err)
	}

	fmt.Println("Connected to user tokens MongoDB")

	return nil
}

func dbCleanUp(ctx context.Context) {
	_, err := subjectsCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		panic(fmt.Errorf("db clean up for subjects collection: %w", err))
	}

	_, err = predicatesCollection.DeleteMany(ctx, bson.M{})

	if err != nil {
		panic(fmt.Errorf("db clean up for predicates collection: %w", err))
	}

	_, err = premisesCollection.DeleteMany(ctx, bson.M{})

	if err != nil {
		panic(fmt.Errorf("db clean up for premises collection: %w", err))
	}

	_, err = argumentsCollection.DeleteMany(ctx, bson.M{})

	if err != nil {
		panic(fmt.Errorf("db clean up for arguments collection: %w", err))
	}
}

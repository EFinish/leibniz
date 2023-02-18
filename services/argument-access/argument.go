package main

import (
	"context"
	"fmt"
	"time"

	protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func insertArgument(ctx context.Context, argument *protoOut.Argument) (insertedArgument *protoOut.Argument, err error) {
	insertArgument := bson.M{
		"created_at":             time.Now(),
		"updated_at":             time.Now(),
		"deleted_at":             nil,
		"title":                  argument.Title,
		"premises":               argument.Premises,
		"conditional_statements": argument.ConditionalStatements,
		"conclusion_premise":     argument.ConclusionPremise,
	}

	aa.logger.Infof("Inserting new arguments record")

	vehicle, err := aa.argumentsCollection.InsertOne(ctx, insertArgument)

	if err != nil {
		return nil, fmt.Errorf("in arguments InsertOne(): %w", err)
	}

	insertedID := vehicle.InsertedID.(primitive.ObjectID).Hex()
	insertedArgument = &protoOut.Argument{
		Id:                    insertedID,
		CreatedAt:             argument.CreatedAt,
		UpdatedAt:             argument.UpdatedAt,
		DeletedAt:             argument.DeletedAt,
		Title:                 argument.Title,
		Premises:              argument.Premises,
		ConditionalStatements: argument.ConditionalStatements,
		ConclusionPremise:     argument.ConclusionPremise,
	}

	aa.logger.Infof("new arguments record %v", insertedID)

	return insertedArgument, nil
}

func getArguments(ctx context.Context, ArgumentID string) ([]*protoOut.Argument, error) {
	filter := bson.M{}

	if len(ArgumentID) > 0 {
		filter["id"] = ArgumentID
	}

	cursor, err := aa.argumentsCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []protoOut.Argument{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var preds []*protoOut.Argument
	for _, pred := range found {
		preds = append(preds, &pred)
	}

	return preds, nil
}

func updateArgument(ctx context.Context, argument *protoOut.Argument) (*protoOut.Argument, error) {
	aa.logger.Infof("Updating Argument record ID %v", argument.Id)

	updateArgument := bson.M{
		"created_at":             time.Now(),
		"updated_at":             time.Now(),
		"deleted_at":             nil,
		"title":                  argument.Title,
		"premises":               argument.Premises,
		"conditional_statements": argument.ConditionalStatements,
		"conclusion_premise":     argument.ConclusionPremise,
	}

	objID, err := primitive.ObjectIDFromHex(argument.Id)

	if err != nil {
		return nil, fmt.Errorf("in primitive.ObjectIDFromHex(%v): %w", argument.Id, err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateArgument}
	upsert := false
	returnDoc := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	}
	vehicle := protoOut.Argument{}
	err = aa.argumentsCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	return &vehicle, nil
}

func deleteArgument(ctx context.Context, ArgumentId string) (*int64, error) {
	filter := bson.M{}
	filter["_id"] = ArgumentId

	result, err := aa.argumentsCollection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("occurs in ArgumentsCollection DeleteMany: %w", err)
	}

	aa.logger.Debugf("Deleted %v records from Arguments database", result.DeletedCount)

	return &result.DeletedCount, nil
}

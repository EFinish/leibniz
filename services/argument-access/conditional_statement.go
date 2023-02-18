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

func insertConditionalStatement(ctx context.Context, conditionalStatement *protoOut.ConditionalStatement) (insertedConditionalStatement *protoOut.ConditionalStatement, err error) {
	insertConditionalStatement := bson.M{
		"created_at":       time.Now(),
		"updated_at":       time.Now(),
		"deleted_at":       nil,
		"if_premise":       conditionalStatement.IfPremise,
		"then_premise":     conditionalStatement.ThenPremise,
		"if_proposition":   conditionalStatement.IfProposition,
		"then_proposition": conditionalStatement.ThenProposition,
	}

	aa.logger.Infof("Inserting new conditionalStatements record")

	vehicle, err := aa.conditionalStatementsCollection.InsertOne(ctx, insertConditionalStatement)

	if err != nil {
		return nil, fmt.Errorf("in conditionalStatements InsertOne(): %w", err)
	}

	insertedID := vehicle.InsertedID.(primitive.ObjectID).Hex()
	insertedConditionalStatement = &protoOut.ConditionalStatement{
		Id:              insertedID,
		CreatedAt:       conditionalStatement.CreatedAt,
		UpdatedAt:       conditionalStatement.UpdatedAt,
		DeletedAt:       conditionalStatement.DeletedAt,
		IfPremise:       conditionalStatement.IfPremise,
		ThenPremise:     conditionalStatement.ThenPremise,
		IfProposition:   conditionalStatement.IfProposition,
		ThenProposition: conditionalStatement.ThenProposition,
	}

	aa.logger.Infof("new conditionalStatements record %v", insertedID)

	return insertedConditionalStatement, nil
}

func getConditionalStatements(ctx context.Context, ConditionalStatementID string) ([]*protoOut.ConditionalStatement, error) {
	filter := bson.M{}

	if len(ConditionalStatementID) > 0 {
		filter["id"] = ConditionalStatementID
	}

	cursor, err := aa.conditionalStatementsCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []protoOut.ConditionalStatement{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var preds []*protoOut.ConditionalStatement
	for _, pred := range found {
		preds = append(preds, &pred)
	}

	return preds, nil
}

func updateConditionalStatement(ctx context.Context, conditionalStatement *protoOut.ConditionalStatement) (*protoOut.ConditionalStatement, error) {
	aa.logger.Infof("Updating ConditionalStatement record ID %v", conditionalStatement.Id)

	updateConditionalStatement := bson.M{
		"created_at":       time.Now(),
		"updated_at":       time.Now(),
		"deleted_at":       nil,
		"if_premise":       conditionalStatement.IfPremise,
		"then_premise":     conditionalStatement.ThenPremise,
		"if_proposition":   conditionalStatement.IfProposition,
		"then_proposition": conditionalStatement.ThenProposition,
	}

	objID, err := primitive.ObjectIDFromHex(conditionalStatement.Id)

	if err != nil {
		return nil, fmt.Errorf("in primitive.ObjectIDFromHex(%v): %w", conditionalStatement.Id, err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateConditionalStatement}
	upsert := false
	returnDoc := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	}
	vehicle := protoOut.ConditionalStatement{}
	err = aa.conditionalStatementsCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	return &vehicle, nil
}

func deleteConditionalStatement(ctx context.Context, ConditionalStatementId string) (*int64, error) {
	filter := bson.M{}
	filter["_id"] = ConditionalStatementId

	result, err := aa.conditionalStatementsCollection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("occurs in ConditionalStatementsCollection DeleteMany: %w", err)
	}

	aa.logger.Debugf("Deleted %v records from ConditionalStatements database", result.DeletedCount)

	return &result.DeletedCount, nil
}

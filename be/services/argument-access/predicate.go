package main

import (
	"context"
	"fmt"
	"time"

	protoOut "github.com/EFinish/leibniz/proto/gen/argumentaccess/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func insertPredicate(ctx context.Context, predicate *protoOut.Predicate) (insertedPredicate *protoOut.Predicate, err error) {
	insertPredicate := bson.M{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"body":       predicate.Body,
	}

	aa.logger.Infof("Inserting new predicates record")

	vehicle, err := aa.predicatesCollection.InsertOne(ctx, insertPredicate)

	if err != nil {
		return nil, fmt.Errorf("in predicates InsertOne(): %w", err)
	}

	insertedID := vehicle.InsertedID.(primitive.ObjectID).Hex()
	insertedPredicate = &protoOut.Predicate{
		Id:        insertedID,
		CreatedAt: predicate.CreatedAt,
		UpdatedAt: predicate.UpdatedAt,
		DeletedAt: predicate.DeletedAt,
		Body:      predicate.Body,
	}

	aa.logger.Infof("new predicates record %v", insertedID)

	return insertedPredicate, nil
}

func getPredicates(ctx context.Context, predicateID *string) ([]*protoOut.Predicate, error) {
	filter := bson.M{}

	if predicateID != nil && len(*predicateID) > 0 {
		filter["id"] = predicateID
	}

	cursor, err := aa.predicatesCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []protoOut.Predicate{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var preds []*protoOut.Predicate
	for _, pred := range found {
		preds = append(preds, &pred)
	}

	return preds, nil
}

func updatePredicate(ctx context.Context, Predicate *protoOut.Predicate) (*protoOut.Predicate, error) {
	aa.logger.Infof("Updating Predicate record ID %v", Predicate.Id)

	updatePredicate := bson.M{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"body":       Predicate.Body,
	}

	objID, err := primitive.ObjectIDFromHex(Predicate.Id)

	if err != nil {
		return nil, fmt.Errorf("in primitive.ObjectIDFromHex(%v): %w", Predicate.Id, err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatePredicate}
	upsert := false
	returnDoc := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	}
	vehicle := protoOut.Predicate{}
	err = aa.predicatesCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	return &vehicle, nil
}

func deletePredicate(ctx context.Context, PredicateId string) (*int64, error) {
	filter := bson.M{}
	filter["_id"] = PredicateId

	result, err := aa.predicatesCollection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("occurs in PredicatesCollection DeleteMany: %w", err)
	}

	aa.logger.Debugf("Deleted %v records from Predicates database", result.DeletedCount)

	return &result.DeletedCount, nil
}

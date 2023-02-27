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

func insertPremise(ctx context.Context, premise *protoOut.Premise) (insertedPremise *protoOut.Premise, err error) {
	insertPremise := bson.M{
		"created_at":          time.Now(),
		"updated_at":          time.Now(),
		"deleted_at":          nil,
		"subject":             premise.Subject,
		"subject_quantifier":  premise.SubjectQuantifier,
		"predicate":           premise.Predicate,
		"predicate_qualifier": premise.PredicateQualifier,
	}

	aa.logger.Infof("Inserting new premises record")

	vehicle, err := aa.premisesCollection.InsertOne(ctx, insertPremise)

	if err != nil {
		return nil, fmt.Errorf("in premises InsertOne(): %w", err)
	}

	insertedID := vehicle.InsertedID.(primitive.ObjectID).Hex()
	insertedPremise = &protoOut.Premise{
		Id:                 insertedID,
		CreatedAt:          premise.CreatedAt,
		UpdatedAt:          premise.UpdatedAt,
		DeletedAt:          premise.DeletedAt,
		Subject:            premise.Subject,
		SubjectQuantifier:  premise.SubjectQuantifier,
		Predicate:          premise.Predicate,
		PredicateQualifier: premise.PredicateQualifier,
	}

	aa.logger.Infof("new premises record %v", insertedID)

	return insertedPremise, nil
}

func getPremises(ctx context.Context, PremiseID string) ([]*protoOut.Premise, error) {
	filter := bson.M{}

	if len(PremiseID) > 0 {
		filter["id"] = PremiseID
	}

	cursor, err := aa.premisesCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []protoOut.Premise{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var preds []*protoOut.Premise
	for _, pred := range found {
		preds = append(preds, &pred)
	}

	return preds, nil
}

func updatePremise(ctx context.Context, premise *protoOut.Premise) (*protoOut.Premise, error) {
	aa.logger.Infof("Updating Premise record ID %v", premise.Id)

	updatePremise := bson.M{
		"created_at":          time.Now(),
		"updated_at":          time.Now(),
		"deleted_at":          nil,
		"subject":             premise.Subject,
		"subject_quantifier":  premise.SubjectQuantifier,
		"predicate":           premise.Predicate,
		"predicate_qualifier": premise.PredicateQualifier,
	}

	objID, err := primitive.ObjectIDFromHex(premise.Id)

	if err != nil {
		return nil, fmt.Errorf("in primitive.ObjectIDFromHex(%v): %w", premise.Id, err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatePremise}
	upsert := false
	returnDoc := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	}
	vehicle := protoOut.Premise{}
	err = aa.premisesCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	return &vehicle, nil
}

func deletePremise(ctx context.Context, PremiseId string) (*int64, error) {
	filter := bson.M{}
	filter["_id"] = PremiseId

	result, err := aa.premisesCollection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("occurs in PremisesCollection DeleteMany: %w", err)
	}

	aa.logger.Debugf("Deleted %v records from Premises database", result.DeletedCount)

	return &result.DeletedCount, nil
}

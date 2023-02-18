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

func insertProposition(ctx context.Context, proposition *protoOut.Proposition) (insertedProposition *protoOut.Proposition, err error) {
	insertProposition := bson.M{
		"created_at":                   time.Now(),
		"updated_at":                   time.Now(),
		"deleted_at":                   nil,
		"proposition_type":             proposition.PropositionType,
		"proposition_sub_premises":     proposition.PropositionSubPremises,
		"proposition_sub_propositions": proposition.PropositionSubPropositions,
	}

	aa.logger.Infof("Inserting new propositions record")

	vehicle, err := aa.propositionsCollection.InsertOne(ctx, insertProposition)

	if err != nil {
		return nil, fmt.Errorf("in propositions InsertOne(): %w", err)
	}

	insertedID := vehicle.InsertedID.(primitive.ObjectID).Hex()
	insertedProposition = &protoOut.Proposition{
		Id:                         insertedID,
		CreatedAt:                  proposition.CreatedAt,
		UpdatedAt:                  proposition.UpdatedAt,
		DeletedAt:                  proposition.DeletedAt,
		PropositionType:            proposition.PropositionType,
		PropositionSubPremises:     proposition.PropositionSubPremises,
		PropositionSubPropositions: proposition.PropositionSubPropositions,
	}

	aa.logger.Infof("new propositions record %v", insertedID)

	return insertedProposition, nil
}

func getPropositions(ctx context.Context, PropositionID string) ([]*protoOut.Proposition, error) {
	filter := bson.M{}

	if len(PropositionID) > 0 {
		filter["id"] = PropositionID
	}

	cursor, err := aa.propositionsCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []protoOut.Proposition{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var preds []*protoOut.Proposition
	for _, pred := range found {
		preds = append(preds, &pred)
	}

	return preds, nil
}

func updateProposition(ctx context.Context, proposition *protoOut.Proposition) (*protoOut.Proposition, error) {
	aa.logger.Infof("Updating Proposition record ID %v", proposition.Id)

	updateProposition := bson.M{
		"created_at":                   time.Now(),
		"updated_at":                   time.Now(),
		"deleted_at":                   nil,
		"proposition_type":             proposition.PropositionType,
		"proposition_sub_premises":     proposition.PropositionSubPremises,
		"proposition_sub_propositions": proposition.PropositionSubPropositions,
	}

	objID, err := primitive.ObjectIDFromHex(proposition.Id)

	if err != nil {
		return nil, fmt.Errorf("in primitive.ObjectIDFromHex(%v): %w", proposition.Id, err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateProposition}
	upsert := false
	returnDoc := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	}
	vehicle := protoOut.Proposition{}
	err = aa.propositionsCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	return &vehicle, nil
}

func deleteProposition(ctx context.Context, PropositionId string) (*int64, error) {
	filter := bson.M{}
	filter["_id"] = PropositionId

	result, err := aa.propositionsCollection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("occurs in PropositionsCollection DeleteMany: %w", err)
	}

	aa.logger.Debugf("Deleted %v records from Propositions database", result.DeletedCount)

	return &result.DeletedCount, nil
}

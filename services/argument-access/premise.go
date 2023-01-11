package main

import (
	"context"
	"fmt"
	"time"

	protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"
	conversion "github.com/EFinish/leibniz/utilities/conversion"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type PremiseJsonConvertable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty"`
	DeletedAt string             `bson:"deleted_at,omitempty"`
	Body      string             `bson:"body,omitempty"`
}

func (s *PremiseJsonConvertable) toProto() (*protoOut.Premise, error) {
	// TODO figure out format of datetimes from mongodb
	createdAt, err := conversion.StringToTimestamp(s.CreatedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting created at timestamp for Premise: %w", err)
	}

	updatedAt, err := conversion.StringToTimestamp(s.UpdatedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting updated at timestamp for Premise: %w", err)
	}

	deletedAt, err := conversion.StringToTimestamp(s.DeletedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting deleted at timestamp for Premise: %w", err)
	}

	return &protoOut.Premise{
		Id:        s.ID.Hex(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
		Body:      s.Body,
	}, nil
}

func insertPremise(ctx context.Context, Premise *protoOut.Premise) (insertedPremise *protoOut.Premise, err error) {
	insertPremise := bson.M{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"body":       Premise.Body,
	}

	aa.logger.Infof("Inserting new Premises record")

	vehicle, err := aa.PremisesCollection.InsertOne(ctx, insertPremise)

	if err != nil {
		return nil, fmt.Errorf("in Premises InsertOne(): %w", err)
	}

	insertedID := vehicle.InsertedID.(primitive.ObjectID).Hex()
	insertedPremise = &protoOut.Premise{
		Id:        insertedID,
		CreatedAt: Premise.CreatedAt,
		UpdatedAt: Premise.UpdatedAt,
		DeletedAt: Premise.DeletedAt,
		Body:      Premise.Body,
	}

	aa.logger.Infof("new Premises record %v", insertedID)

	return insertedPremise, nil
}

func getPremises(ctx context.Context, PremiseID string) ([]*protoOut.Premise, error) {
	filter := bson.M{}

	if len(PremiseID) > 0 {
		filter["id"] = PremiseID
	}

	cursor, err := aa.PremisesCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []PremiseJsonConvertable{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var subs []*protoOut.Premise
	for _, sub := range found {
		created, err := time.Parse(time.RFC3339, sub.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("while parsing createdat: %w", err)
		}

		updated, err := time.Parse(time.RFC3339, sub.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("while parsing updatedat: %w", err)
		}

		deleted, err := time.Parse(time.RFC3339, sub.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("while parsing deletedat: %w", err)
		}

		subs = append(subs, &protoOut.Premise{
			Id:        sub.ID.Hex(),
			CreatedAt: timestamppb.New(created),
			UpdatedAt: timestamppb.New(updated),
			DeletedAt: timestamppb.New(deleted),
			Body:      sub.Body,
		})
	}

	return subs, nil
}

func updatePremise(ctx context.Context, Premise *protoOut.Premise) (*protoOut.Premise, error) {
	aa.logger.Infof("Updating Premise record ID %v", Premise.Id)

	updatePremise := bson.M{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"body":       Premise.Body,
	}

	objID, err := primitive.ObjectIDFromHex(Premise.Id)

	if err != nil {
		return nil, fmt.Errorf("in primitive.ObjectIDFromHex(%v): %w", Premise.Id, err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatePremise}
	upsert := false
	returnDoc := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	}
	vehicle := PremiseJsonConvertable{}
	err = aa.PremisesCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	protoPremise, err := vehicle.toProto()

	if err != nil {
		return nil, fmt.Errorf("during protoification: %w", err)
	}

	return protoPremise, nil
}

func deletePremise(ctx context.Context, PremiseId string) (*int64, error) {
	filter := bson.M{}
	filter["_id"] = PremiseId

	result, err := aa.PremisesCollection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("occurs in PremisesCollection DeleteMany: %w", err)
	}

	aa.logger.Debugf("Deleted %v records from Premises database", result.DeletedCount)

	return &result.DeletedCount, nil
}

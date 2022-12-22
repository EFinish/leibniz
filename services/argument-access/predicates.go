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

type PredicateJsonConvertable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty"`
	DeletedAt string             `bson:"deleted_at,omitempty"`
	Body      string             `bson:"body,omitempty"`
}

func (s *PredicateJsonConvertable) toProto() (*protoOut.Predicate, error) {
	// TODO figure out format of datetimes from mongodb
	createdAt, err := conversion.StringToTimestamp(s.CreatedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting created at timestamp for predicate: %w", err)
	}

	updatedAt, err := conversion.StringToTimestamp(s.UpdatedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting updated at timestamp for predicate: %w", err)
	}

	deletedAt, err := conversion.StringToTimestamp(s.DeletedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting deleted at timestamp for predicate: %w", err)
	}

	return &protoOut.Predicate{
		Id:        s.ID.Hex(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
		Body:      s.Body,
	}, nil
}

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

func getPredicates(ctx context.Context, PredicateID string) ([]*protoOut.Predicate, error) {
	filter := bson.M{}

	if len(PredicateID) > 0 {
		filter["id"] = PredicateID
	}

	cursor, err := aa.predicatesCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []PredicateJsonConvertable{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var preds []*protoOut.Predicate
	for _, pred := range found {
		created, err := time.Parse(time.RFC3339, pred.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("while parsing createdat: %w", err)
		}

		updated, err := time.Parse(time.RFC3339, pred.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("while parsing updatedat: %w", err)
		}

		deleted, err := time.Parse(time.RFC3339, pred.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("while parsing deletedat: %w", err)
		}

		preds = append(preds, &protoOut.Predicate{
			Id:        pred.ID.Hex(),
			CreatedAt: timestamppb.New(created),
			UpdatedAt: timestamppb.New(updated),
			DeletedAt: timestamppb.New(deleted),
			Body:      pred.Body,
		})
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
	vehicle := PredicateJsonConvertable{}
	err = aa.predicatesCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	protoPredicate, err := vehicle.toProto()

	if err != nil {
		return nil, fmt.Errorf("during protoification: %w", err)
	}

	return protoPredicate, nil
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

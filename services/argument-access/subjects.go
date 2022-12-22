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

type SubjectJsonConvertable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty"`
	DeletedAt string             `bson:"deleted_at,omitempty"`
	Body      string             `bson:"body,omitempty"`
}

func (s *SubjectJsonConvertable) toProto() (*protoOut.Subject, error) {
	// TODO figure out format of datetimes from mongodb
	createdAt, err := conversion.StringToTimestamp(s.CreatedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting created at timestamp for subject: %w", err)
	}

	updatedAt, err := conversion.StringToTimestamp(s.UpdatedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting updated at timestamp for subject: %w", err)
	}

	deletedAt, err := conversion.StringToTimestamp(s.DeletedAt, "TODO")

	if err != nil {
		return nil, fmt.Errorf("getting deleted at timestamp for subject: %w", err)
	}

	return &protoOut.Subject{
		Id:        s.ID.Hex(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
		Body:      s.Body,
	}, nil
}

func insertSubject(ctx context.Context, subject *protoOut.Subject) (insertedSubject *protoOut.Subject, err error) {
	insertSubject := bson.M{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"body":       subject.Body,
	}

	aa.logger.Infof("Inserting new subjects record")

	vehicle, err := aa.subjectsCollection.InsertOne(ctx, insertSubject)

	if err != nil {
		return nil, fmt.Errorf("in subjects InsertOne(): %w", err)
	}

	insertedID := vehicle.InsertedID.(primitive.ObjectID).Hex()
	insertedSubject = &protoOut.Subject{
		Id:        insertedID,
		CreatedAt: subject.CreatedAt,
		UpdatedAt: subject.UpdatedAt,
		DeletedAt: subject.DeletedAt,
		Body:      subject.Body,
	}

	aa.logger.Infof("new subjects record %v", insertedID)

	return insertedSubject, nil
}

func getSubjects(ctx context.Context, subjectID string) ([]*protoOut.Subject, error) {
	filter := bson.M{}

	if len(subjectID) > 0 {
		filter["id"] = subjectID
	}

	cursor, err := aa.subjectsCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDb error: %w", err)
	}

	found := []SubjectJsonConvertable{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var subs []*protoOut.Subject
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

		subs = append(subs, &protoOut.Subject{
			Id:        sub.ID.Hex(),
			CreatedAt: timestamppb.New(created),
			UpdatedAt: timestamppb.New(updated),
			DeletedAt: timestamppb.New(deleted),
			Body:      sub.Body,
		})
	}

	return subs, nil
}

func updateSubject(ctx context.Context, subject *protoOut.Subject) (*protoOut.Subject, error) {
	aa.logger.Infof("Updating subject record ID %v", subject.Id)

	updateSubject := bson.M{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"body":       subject.Body,
	}

	objID, err := primitive.ObjectIDFromHex(subject.Id)

	if err != nil {
		return nil, fmt.Errorf("in primitive.ObjectIDFromHex(%v): %w", subject.Id, err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateSubject}
	upsert := false
	returnDoc := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	}
	vehicle := SubjectJsonConvertable{}
	err = aa.subjectsCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	protoSubject, err := vehicle.toProto()

	if err != nil {
		return nil, fmt.Errorf("during protoification: %w", err)
	}

	return protoSubject, nil
}

func deleteSubject(ctx context.Context, subjectId string) (*int64, error) {
	filter := bson.M{}
	filter["_id"] = subjectId

	result, err := aa.subjectsCollection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("occurs in subjectsCollection DeleteMany: %w", err)
	}

	aa.logger.Debugf("Deleted %v records from subjects database", result.DeletedCount)

	return &result.DeletedCount, nil
}

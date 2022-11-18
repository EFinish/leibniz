package main

import (
	"context"
	"fmt"
	"time"

	protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"
	conversion "github.com/EFinish/leibniz/utilities/conversion"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// TODO test this function
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

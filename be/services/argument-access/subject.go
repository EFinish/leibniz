package main

import (
	"context"
	"fmt"
	"time"

	protoOut "github.com/EFinish/leibniz/proto/gen/argumentaccess/v1"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/mgo.v2/bson"
)

func insertSubject(ctx context.Context, subject *protoOut.Subject) (insertedSubject *protoOut.Subject, err error) {
	subject.CreatedAt = timestamppb.Now()
	subject.UpdatedAt = timestamppb.Now()
	subject.DeletedAt = nil

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

	found := []protoOut.Subject{}
	err = cursor.All(ctx, &found)

	if err != nil {
		return nil, fmt.Errorf("while decoding: %w", err)
	}

	var subs []*protoOut.Subject
	for _, sub := range found {
		subs = append(subs, &sub)
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
	vehicle := protoOut.Subject{}
	err = aa.subjectsCollection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&vehicle)

	if err != nil {
		return nil, fmt.Errorf("in FindOneAndUpdate: %w", err)
	}

	return &vehicle, nil
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

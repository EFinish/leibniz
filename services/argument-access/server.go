package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"
)

type PredicateJsonConvertable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty"`
	DeletedAt string             `bson:"deleted_at,omitempty"`
	Body      string             `bson:"body,omitempty"`
}

type PremiseJsonConvertable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty"`
	DeletedAt string             `bson:"deleted_at,omitempty"`
	// Subject                   string             `bson:"subject,omitempty"`
	SubjectInclusivenessLevel string `bson:"subject_inclusiveness_level,omitempty"`
	// Predicate                 string             `bson:"predicate,omitempty"`
	// SubPremises               string             `bson:"sub_premises,omitempty"`
	ParallelCondition string `bson:"parallel_condition,omitempty"`
	// ParallelConditionPremise  string             `bson:"parallel_condition_premise,omitempty"`
	// ImplicationPremise        string             `bson:"implication_premise,omitempty"`
}

type ArgumentJSONConvertable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	UpdatedAt string             `bson:"updated_at,omitempty"`
	DeletedAt string             `bson:"deleted_at,omitempty"`
	Title     string             `bson:"title,omitempty"`
	// Premises          string             `bson:"premises,omitempty"`
	// ConclusionPremise string             `bson:"conclusion_premise,omitempty"`
}

func (s *ArgumentAccessServiceServer) CreateSubject(ctx context.Context, req *protoOut.CreateSubjectRequest) (*protoOut.CreateSubjectResponse, error) {
	if req.Subject == nil {
		return nil, errors.New("empty subject")
	}
	if len(req.Subject.Id) > 0 {
		return nil, errors.New("no ids allowed during subject creation")
	}

	subject, err := insertSubject(ctx, req.Subject)

	if err != nil {
		return nil, fmt.Errorf("during subject insertion: %w", err)
	}

	return &protoOut.CreateSubjectResponse{
		Subject: subject,
	}, nil
}

// func (s *ArgumentAccessServiceServer) ReadSubject(ctx context.Context, req *protoOut.ReadSubjectRequest) (*protoOut, protoOut.ReadSubjectResponse, error) {

// }

func (s *ArgumentAccessServiceServer) UpdateSubject(ctx context.Context, req *protoOut.UpdateSubjectRequest) (*protoOut.UpdateSubjectResponse, error) {
	if req.Subject == nil {
		return nil, errors.New("empty subject")
	}
	if len(req.Subject.Id) == 0 {
		return nil, errors.New("id required for updating subject")
	}

	subject, err := updateSubject(ctx, req.Subject)

	if err != nil {
		return nil, fmt.Errorf("during subject updating: %w", err)
	}

	return &protoOut.UpdateSubjectResponse{
		Subject: subject,
	}, nil
}

func (s *ArgumentAccessServiceServer) DeleteSubject(ctx context.Context, req *protoOut.DeleteSubjectRequest) (*protoOut.DeleteSubjectResponse, error) {
	if len(req.SubjectId) == 0 {
		return nil, errors.New("id required for deleting subject")
	}

	numDeleted, err := deleteSubject(ctx, req.SubjectId)

	if err != nil {
		return nil, fmt.Errorf("during subject deletion: %w", err)
	}

	return &protoOut.DeleteSubjectResponse{
		DeletedCount: *numDeleted,
	}, nil
}

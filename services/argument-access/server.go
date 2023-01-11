package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	protoOut "github.com/EFinish/leibniz/proto/gen/go/argumentaccess/v1"
)

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

func (s *ArgumentAccessServiceServer) ReadSubject(ctx context.Context, req *protoOut.ReadSubjectRequest) (*protoOut.ReadSubjectResponse, error) {
	subjects, err := getSubjects(ctx, req.SubjectId)

	if err != nil {
		return nil, fmt.Errorf("during subject reading: %w", err)
	}

	return &protoOut.ReadSubjectResponse{
		Subjects: subjects,
	}, nil
}

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

func (s *ArgumentAccessServiceServer) CreatePredicate(ctx context.Context, req *protoOut.CreatePredicateRequest) (*protoOut.CreatePredicateResponse, error) {
	if req.Predicate == nil {
		return nil, errors.New("empty Predicate")
	}
	if len(req.Predicate.Id) > 0 {
		return nil, errors.New("no ids allowed during Predicate creation")
	}

	Predicate, err := insertPredicate(ctx, req.Predicate)

	if err != nil {
		return nil, fmt.Errorf("during Predicate insertion: %w", err)
	}

	return &protoOut.CreatePredicateResponse{
		Predicate: Predicate,
	}, nil
}

func (s *ArgumentAccessServiceServer) ReadPredicate(ctx context.Context, req *protoOut.ReadPredicateRequest) (*protoOut.ReadPredicateResponse, error) {
	Predicates, err := getPredicates(ctx, req.PredicateId)

	if err != nil {
		return nil, fmt.Errorf("during Predicate reading: %w", err)
	}

	return &protoOut.ReadPredicateResponse{
		Predicates: Predicates,
	}, nil
}

func (s *ArgumentAccessServiceServer) UpdatePredicate(ctx context.Context, req *protoOut.UpdatePredicateRequest) (*protoOut.UpdatePredicateResponse, error) {
	if req.Predicate == nil {
		return nil, errors.New("empty Predicate")
	}
	if len(req.Predicate.Id) == 0 {
		return nil, errors.New("id required for updating Predicate")
	}

	Predicate, err := updatePredicate(ctx, req.Predicate)

	if err != nil {
		return nil, fmt.Errorf("during Predicate updating: %w", err)
	}

	return &protoOut.UpdatePredicateResponse{
		Predicate: Predicate,
	}, nil
}

func (s *ArgumentAccessServiceServer) DeletePredicate(ctx context.Context, req *protoOut.DeletePredicateRequest) (*protoOut.DeletePredicateResponse, error) {
	if len(req.PredicateId) == 0 {
		return nil, errors.New("id required for deleting Predicate")
	}

	numDeleted, err := deletePredicate(ctx, req.PredicateId)

	if err != nil {
		return nil, fmt.Errorf("during Predicate deletion: %w", err)
	}

	return &protoOut.DeletePredicateResponse{
		DeletedCount: *numDeleted,
	}, nil
}

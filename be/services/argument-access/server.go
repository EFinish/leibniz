package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	protoOut "github.com/EFinish/leibniz/proto/gen/argumentaccess/v1"
)

type ArgumentJSONConvertable struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt primitive.ObjectID `bson:"created_at,omitempty"`
	UpdatedAt primitive.ObjectID `bson:"updated_at,omitempty"`
	DeletedAt primitive.ObjectID `bson:"deleted_at,omitempty"`
	Title     string             `bson:"title,omitempty"`
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
	subjects, err := getSubjects(ctx, *req.SubjectId)

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

func (s *ArgumentAccessServiceServer) CreatePremise(ctx context.Context, req *protoOut.CreatePremiseRequest) (*protoOut.CreatePremiseResponse, error) {
	if req.Premise == nil {
		return nil, errors.New("empty Premise")
	}
	if len(req.Premise.Id) > 0 {
		return nil, errors.New("no ids allowed during Premise creation")
	}

	Premise, err := insertPremise(ctx, req.Premise)

	if err != nil {
		return nil, fmt.Errorf("during Premise insertion: %w", err)
	}

	return &protoOut.CreatePremiseResponse{
		Premise: Premise,
	}, nil
}

func (s *ArgumentAccessServiceServer) ReadPremise(ctx context.Context, req *protoOut.ReadPremiseRequest) (*protoOut.ReadPremiseResponse, error) {
	Premises, err := getPremises(ctx, req.PremiseId)

	if err != nil {
		return nil, fmt.Errorf("during Premise reading: %w", err)
	}

	return &protoOut.ReadPremiseResponse{
		Premises: Premises,
	}, nil
}

func (s *ArgumentAccessServiceServer) UpdatePremise(ctx context.Context, req *protoOut.UpdatePremiseRequest) (*protoOut.UpdatePremiseResponse, error) {
	if req.Premise == nil {
		return nil, errors.New("empty Premise")
	}
	if len(req.Premise.Id) == 0 {
		return nil, errors.New("id required for updating Premise")
	}

	Premise, err := updatePremise(ctx, req.Premise)

	if err != nil {
		return nil, fmt.Errorf("during Premise updating: %w", err)
	}

	return &protoOut.UpdatePremiseResponse{
		Premise: Premise,
	}, nil
}

func (s *ArgumentAccessServiceServer) DeletePremise(ctx context.Context, req *protoOut.DeletePremiseRequest) (*protoOut.DeletePremiseResponse, error) {
	if len(req.PremiseId) == 0 {
		return nil, errors.New("id required for deleting Premise")
	}

	numDeleted, err := deletePremise(ctx, req.PremiseId)

	if err != nil {
		return nil, fmt.Errorf("during Premise deletion: %w", err)
	}

	return &protoOut.DeletePremiseResponse{
		DeletedCount: *numDeleted,
	}, nil
}

func (s *ArgumentAccessServiceServer) CreateProposition(ctx context.Context, req *protoOut.CreatePropositionRequest) (*protoOut.CreatePropositionResponse, error) {
	if req.Proposition == nil {
		return nil, errors.New("empty Proposition")
	}
	if len(req.Proposition.Id) > 0 {
		return nil, errors.New("no ids allowed during Proposition creation")
	}

	Proposition, err := insertProposition(ctx, req.Proposition)

	if err != nil {
		return nil, fmt.Errorf("during Proposition insertion: %w", err)
	}

	return &protoOut.CreatePropositionResponse{
		Proposition: Proposition,
	}, nil
}

func (s *ArgumentAccessServiceServer) ReadProposition(ctx context.Context, req *protoOut.ReadPropositionRequest) (*protoOut.ReadPropositionResponse, error) {
	Propositions, err := getPropositions(ctx, req.PropositionId)

	if err != nil {
		return nil, fmt.Errorf("during Proposition reading: %w", err)
	}

	return &protoOut.ReadPropositionResponse{
		Propositions: Propositions,
	}, nil
}

func (s *ArgumentAccessServiceServer) UpdateProposition(ctx context.Context, req *protoOut.UpdatePropositionRequest) (*protoOut.UpdatePropositionResponse, error) {
	if req.Proposition == nil {
		return nil, errors.New("empty Proposition")
	}
	if len(req.Proposition.Id) == 0 {
		return nil, errors.New("id required for updating Proposition")
	}

	Proposition, err := updateProposition(ctx, req.Proposition)

	if err != nil {
		return nil, fmt.Errorf("during Proposition updating: %w", err)
	}

	return &protoOut.UpdatePropositionResponse{
		Proposition: Proposition,
	}, nil
}

func (s *ArgumentAccessServiceServer) DeleteProposition(ctx context.Context, req *protoOut.DeletePropositionRequest) (*protoOut.DeletePropositionResponse, error) {
	if len(req.PropositionId) == 0 {
		return nil, errors.New("id required for deleting Proposition")
	}

	numDeleted, err := deleteProposition(ctx, req.PropositionId)

	if err != nil {
		return nil, fmt.Errorf("during Proposition deletion: %w", err)
	}

	return &protoOut.DeletePropositionResponse{
		DeletedCount: *numDeleted,
	}, nil
}

func (s *ArgumentAccessServiceServer) CreateConditionalStatement(ctx context.Context, req *protoOut.CreateConditionalStatementRequest) (*protoOut.CreateConditionalStatementResponse, error) {
	if req.ConditionalStatement == nil {
		return nil, errors.New("empty ConditionalStatement")
	}
	if len(req.ConditionalStatement.Id) > 0 {
		return nil, errors.New("no ids allowed during ConditionalStatement creation")
	}

	ConditionalStatement, err := insertConditionalStatement(ctx, req.ConditionalStatement)

	if err != nil {
		return nil, fmt.Errorf("during ConditionalStatement insertion: %w", err)
	}

	return &protoOut.CreateConditionalStatementResponse{
		ConditionalStatement: ConditionalStatement,
	}, nil
}

func (s *ArgumentAccessServiceServer) ReadConditionalStatement(ctx context.Context, req *protoOut.ReadConditionalStatementRequest) (*protoOut.ReadConditionalStatementResponse, error) {
	ConditionalStatements, err := getConditionalStatements(ctx, req.ConditionalStatementId)

	if err != nil {
		return nil, fmt.Errorf("during ConditionalStatement reading: %w", err)
	}

	return &protoOut.ReadConditionalStatementResponse{
		ConditionalStatements: ConditionalStatements,
	}, nil
}

func (s *ArgumentAccessServiceServer) UpdateConditionalStatement(ctx context.Context, req *protoOut.UpdateConditionalStatementRequest) (*protoOut.UpdateConditionalStatementResponse, error) {
	if req.ConditionalStatement == nil {
		return nil, errors.New("empty ConditionalStatement")
	}
	if len(req.ConditionalStatement.Id) == 0 {
		return nil, errors.New("id required for updating ConditionalStatement")
	}

	ConditionalStatement, err := updateConditionalStatement(ctx, req.ConditionalStatement)

	if err != nil {
		return nil, fmt.Errorf("during ConditionalStatement updating: %w", err)
	}

	return &protoOut.UpdateConditionalStatementResponse{
		ConditionalStatement: ConditionalStatement,
	}, nil
}

func (s *ArgumentAccessServiceServer) DeleteConditionalStatement(ctx context.Context, req *protoOut.DeleteConditionalStatementRequest) (*protoOut.DeleteConditionalStatementResponse, error) {
	if len(req.ConditionalStatementId) == 0 {
		return nil, errors.New("id required for deleting ConditionalStatement")
	}

	numDeleted, err := deleteConditionalStatement(ctx, req.ConditionalStatementId)

	if err != nil {
		return nil, fmt.Errorf("during ConditionalStatement deletion: %w", err)
	}

	return &protoOut.DeleteConditionalStatementResponse{
		DeletedCount: *numDeleted,
	}, nil
}

func (s *ArgumentAccessServiceServer) CreateCArgument(ctx context.Context, req *protoOut.CreateArgumentRequest) (*protoOut.CreateArgumentResponse, error) {
	if req.Argument == nil {
		return nil, errors.New("empty Argument")
	}
	if len(req.Argument.Id) > 0 {
		return nil, errors.New("no ids allowed during Argument creation")
	}

	Argument, err := insertArgument(ctx, req.Argument)

	if err != nil {
		return nil, fmt.Errorf("during Argument insertion: %w", err)
	}

	return &protoOut.CreateArgumentResponse{
		Argument: Argument,
	}, nil
}

func (s *ArgumentAccessServiceServer) ReadArgument(ctx context.Context, req *protoOut.ReadArgumentRequest) (*protoOut.ReadArgumentResponse, error) {
	Arguments, err := getArguments(ctx, req.ArgumentId)

	if err != nil {
		return nil, fmt.Errorf("during Argument reading: %w", err)
	}

	return &protoOut.ReadArgumentResponse{
		Arguments: Arguments,
	}, nil
}

func (s *ArgumentAccessServiceServer) UpdateArgument(ctx context.Context, req *protoOut.UpdateArgumentRequest) (*protoOut.UpdateArgumentResponse, error) {
	if req.Argument == nil {
		return nil, errors.New("empty Argument")
	}
	if len(req.Argument.Id) == 0 {
		return nil, errors.New("id required for updating Argument")
	}

	Argument, err := updateArgument(ctx, req.Argument)

	if err != nil {
		return nil, fmt.Errorf("during Argument updating: %w", err)
	}

	return &protoOut.UpdateArgumentResponse{
		Argument: Argument,
	}, nil
}

func (s *ArgumentAccessServiceServer) DeleteArgument(ctx context.Context, req *protoOut.DeleteArgumentRequest) (*protoOut.DeleteArgumentResponse, error) {
	if len(req.ArgumentId) == 0 {
		return nil, errors.New("id required for deleting Argument")
	}

	numDeleted, err := deleteArgument(ctx, req.ArgumentId)

	if err != nil {
		return nil, fmt.Errorf("during Argument deletion: %w", err)
	}

	return &protoOut.DeleteArgumentResponse{
		DeletedCount: *numDeleted,
	}, nil
}

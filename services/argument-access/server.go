package main

import (
	"context"
	"errors"

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
		aa.logger.Errorf("empty subject")
		return nil, errors.New("empty subject")
	}
	if len(req.Subject.Id) == 0 {
		aa.logger.Errorf("no ids allowed during subject creation")
		return nil, errors.New("no ids allowed during subject creation")
	}

	return &protoOut.CreateSubjectResponse{
		Subject: nil,
	}, nil
}

// func (s *BrandAccessServiceServer) FilterBrandConfig(ctx context.Context, req *protoOut.FilterBrandConfigRequest) (*protoOut.FilterBrandConfigResponse, error) {

// 	brandId := req.BrandId
// 	if len(brandId) == 0 {
// 		err := serr.InvalidArgumentError("brand_id is required").SetExternalMsgFromError()
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}
// 	filter := bson.M{"brand_id": brandId}

// 	cursor, err := ba.configurationCollection.Find(ctx, filter)
// 	if err != nil {
// 		err = fmt.Errorf("MongoDb error: %w", err)
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}

// 	foundbcs := []BrandConfigJSONConvertable{}
// 	err = cursor.All(ctx, &foundbcs)
// 	if err != nil {
// 		err = fmt.Errorf("while decoding: %w", err)
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}
// 	if len(foundbcs) == 0 {
// 		err = serr.NotFoundError("No brand config found for brand id: %s", brandId).SetExternalMsgFromError()
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}

// 	var bcs []*protoOut.BrandConfig
// 	for _, bc := range foundbcs {
// 		bcs = append(bcs, &protoOut.BrandConfig{
// 			Id:                                  bc.ID.Hex(),
// 			BrandId:                             bc.BrandID,
// 			BrandName:                           bc.BrandName,
// 			GarminConsumerKey:                   bc.GarminConsumerKey,
// 			GarminConsumerSecret:                bc.GarminConsumerSecret,
// 			GarminFailureUrl:                    bc.GarminFailureUrl,
// 			GarminFailureAlreadyConnectedUrl:    bc.GarminFailureAlreadyConnectedUrl,
// 			GarminSuccessUrl:                    bc.GarminSuccessUrl,
// 			PolarConsumerKey:                    bc.PolarConsumerKey,
// 			PolarConsumerSecret:                 bc.PolarConsumerSecret,
// 			PolarFailureUrl:                     bc.PolarFailureUrl,
// 			PolarFailureAlreadyConnectedUrl:     bc.PolarFailureAlreadyConnectedUrl,
// 			PolarSuccessUrl:                     bc.PolarSuccessUrl,
// 			MilestoneWebhookUrl:                 bc.MilestoneWebhookUrl,
// 			CommunityMilestoneWebhookUrl:        bc.CommunityMilestoneWebhookUrl,
// 			UserConnectedWebhookUrl:             bc.UserConnectedWebhookUrl,
// 			StravaConsumerKey:                   bc.StravaConsumerKey,
// 			StravaConsumerSecret:                bc.StravaConsumerSecret,
// 			StravaFailureUrl:                    bc.StravaFailureUrl,
// 			StravaFailureAlreadyConnectedUrl:    bc.StravaFailureAlreadyConnectedUrl,
// 			StravaSuccessUrl:                    bc.StravaSuccessUrl,
// 			AuthType:                            bc.AuthType,
// 			AuthJwtRsaPublicKey:                 bc.AuthJWTRSAPublicKey,
// 			FitbitClientId:                      bc.FitbitClientId,
// 			FitbitClientSecret:                  bc.FitbitClientSecret,
// 			FitbitFailureUrl:                    bc.FitbitFailureUrl,
// 			FitbitFailureAlreadyConnectedUrl:    bc.FitbitFailureAlreadyConnectedUrl,
// 			FitbitSuccessUrl:                    bc.FitbitSuccessUrl,
// 			FitbitSubscriptionVerificationCode:  bc.FitbitSubscriptionVerificationCode,
// 			GoogleFitClientId:                   bc.GoogleFitClientId,
// 			GoogleFitClientSecret:               bc.GoogleFitClientSecret,
// 			GoogleFitFailureUrl:                 bc.GoogleFitFailureUrl,
// 			GoogleFitSuccessUrl:                 bc.GoogleFitSuccessUrl,
// 			GoogleFitFailureAlreadyConnectedUrl: bc.GoogleFitFailureAlreadyConnectedUrl,
// 		})
// 	}

// 	return &protoOut.FilterBrandConfigResponse{Result: bcs}, nil
// }

// func (s *BrandAccessServiceServer) DeleteBrandConfig(ctx context.Context, req *protoOut.DeleteBrandConfigRequest) (*protoOut.DeleteBrandConfigResponse, error) {
// 	brandId := req.BrandId
// 	if len(brandId) == 0 {
// 		err := serr.InvalidArgumentError("Brand_id is required").SetExternalMsgFromError()
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}

// 	result, err := ba.configurationCollection.DeleteOne(ctx, bson.M{"brand_id": brandId})
// 	if err != nil {
// 		err = fmt.Errorf("MongoDb error: %w", err)
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}
// 	if result.DeletedCount == 0 {
// 		err = serr.NotFoundError(fmt.Sprintf("No deleteable Brand-config for brand_id: %s", brandId))
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}

// 	return &protoOut.DeleteBrandConfigResponse{DeletedCount: result.DeletedCount}, nil
// }

// func (s *BrandAccessServiceServer) UpsertBrandConfig(ctx context.Context, req *protoOut.UpsertBrandConfigRequest) (*protoOut.UpsertBrandConfigResponse, error) {
// 	if len(req.BrandId) == 0 {
// 		err := serr.InvalidArgumentError("Brand_id is required").SetExternalMsgFromError()
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}

// 	set := bson.M{}
// 	req.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
// 		set[string(fd.Name())] = v.String()
// 		return true
// 	})
// 	update := bson.M{"$set": set}

// 	trueValue := true
// 	opts := &options.UpdateOptions{Upsert: &trueValue}

// 	result, err := ba.configurationCollection.UpdateOne(ctx, bson.M{"brand_id": req.BrandId}, update, opts)
// 	if err != nil {
// 		err = fmt.Errorf("mongoDB error: %w", err)
// 		return nil, serr.GrpcErr(ba.logger, err)
// 	}

// 	var r int64
// 	if result.UpsertedCount > 0 {
// 		r = result.UpsertedCount
// 	} else {
// 		r = result.ModifiedCount
// 	}

// 	return &protoOut.UpsertBrandConfigResponse{
// 		UpsertedCount: r,
// 	}, nil
// }

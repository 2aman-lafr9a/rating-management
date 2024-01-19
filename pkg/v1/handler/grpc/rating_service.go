package grpc

import (
	"Rating-management/internal/models"
	interfaces "Rating-management/pkg/v1"
	pb "Rating-management/proto"
	"context"
	"google.golang.org/grpc"
)

type RatingService struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedRatingServer
}

func NewServer3(grpcServer *grpc.Server, useCase interfaces.UseCaseInterface) {
	useGrpc := &RatingService{useCase: useCase}
	pb.RegisterRatingServer(grpcServer, useGrpc)
}

func (srv *RatingService) FindAll(context.Context, *pb.GetRatingsRequest) (*pb.GetRatingsResponse, error) {
	ratings, err := srv.useCase.FindAll()
	if ratings == nil {
		return &pb.GetRatingsResponse{}, nil
	}
	return srv.transformRatingsModel(ratings), err
}

func (srv *RatingService) FindById(ctx context.Context, req *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {
	rating, err := srv.useCase.FindById(req.OfferId, req.PlayerId)
	if rating == nil {
		return &pb.GetRatingResponse{}, nil
	}
	return srv.transformRatingModel(*rating), err
}

func (srv *RatingService) Update(ctx context.Context, req *pb.UpdateRatingRequest) (*pb.UpdateRatingResponse, error) {
	rating, _ := srv.useCase.FindById(req.Rating.OfferId, req.Rating.PlayerId)
	if rating == nil {
		_, err := srv.useCase.Create(srv.transformRatingRPC(req))
		if err != nil {
			return nil, err
		}
		return &pb.UpdateRatingResponse{
			Rating: &pb.RatingItem{
				Rating:   req.Rating.Rating,
				PlayerId: req.Rating.PlayerId,
				OfferId:  req.Rating.OfferId,
			},
		}, nil
	} else {
		err := srv.useCase.Update(srv.transformRatingRPC(req))
		if err != nil {
			return nil, err
		}
		return &pb.UpdateRatingResponse{}, err
	}
}

func (srv *RatingService) transformRatingsModel(ratings []*models.Rating) *pb.GetRatingsResponse {
	var ratingsRPC []*pb.RatingItem
	for _, rating := range ratings {
		ratingsRPC = append(ratingsRPC, &pb.RatingItem{
			Rating:   rating.Rating,
			PlayerId: rating.PlayerID,
			OfferId:  rating.OfferID,
		})
	}
	return &pb.GetRatingsResponse{Ratings: ratingsRPC}
}

func (srv *RatingService) transformRatingModel(rating models.Rating) *pb.GetRatingResponse {
	return &pb.GetRatingResponse{
		Rating: &pb.RatingItem{
			Rating:   rating.Rating,
			PlayerId: rating.PlayerID,
			OfferId:  rating.OfferID,
		},
	}
}

func (srv *RatingService) transformRatingRPC(req *pb.UpdateRatingRequest) *models.Rating {
	return &models.Rating{
		Rating:   req.Rating.Rating,
		PlayerID: req.Rating.PlayerId,
		OfferID:  req.Rating.OfferId,
	}
}

package tests

import (
	pb "Rating-management/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestCreateRating(t *testing.T) {
	conn, err := grpc.Dial("localhost:50006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewRatingClient(conn)

	request := &pb.UpdateRatingRequest{
		Rating: &pb.RatingItem{
			Rating:   5,
			PlayerId: "1",
			OfferId:  "1",
		},
	}

	_, err = client.Update(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling CreateRating: %s", err)
	}
	res, err := client.FindById(context.Background(), &pb.GetRatingRequest{
		PlayerId: "1",
		OfferId:  "1",
	})

	if res.Rating.Rating != 5 {
		t.Fatalf("Expected rating to be %d, got %d", request.Rating.Rating, res.Rating.Rating)
	}
}

func TestGetRating(t *testing.T) {
	conn, err := grpc.Dial("localhost:50006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewRatingClient(conn)

	request := &pb.GetRatingRequest{
		PlayerId: "1",
		OfferId:  "1",
	}

	res, err := client.FindById(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling GetRating: %s", err)
	}

	if res.Rating.Rating != 5 {
		t.Fatalf("Expected offerId to be %s, got %s", request.OfferId, res.Rating.OfferId)
	}
}

func TestGetRatings(t *testing.T) {
	conn, err := grpc.Dial("localhost:50006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewRatingClient(conn)

	request := &pb.GetRatingsRequest{}

	res, err := client.FindAll(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling GetRatings: %s", err)
	}

	if res.Ratings == nil {
		t.Fatalf("Expected ratings to be %v, got %v", request, res.Ratings)
	}
}

func TestUpdateRating(t *testing.T) {
	conn, err := grpc.Dial("localhost:50006", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewRatingClient(conn)

	request := &pb.UpdateRatingRequest{
		Rating: &pb.RatingItem{
			Rating:   5,
			PlayerId: "1",
			OfferId:  "1",
		},
	}

	_, err = client.Update(context.Background(), request)

	if err != nil {
		t.Fatalf("Error when calling UpdateRating: %s", err)
	}

	res, err := client.FindById(context.Background(), &pb.GetRatingRequest{
		PlayerId: "1",
		OfferId:  "1",
	})

	if res.Rating.Rating != 5 {
		t.Fatalf("Expected rating to be %d, got %d", request.Rating.Rating, res.Rating.Rating)
	}
}

package main

import (
	"Rating-management/internal/db"
	"Rating-management/internal/models"
	v1 "Rating-management/pkg/v1"
	handler "Rating-management/pkg/v1/handler/grpc"
	repo "Rating-management/pkg/v1/repository"
	usecase "Rating-management/pkg/v1/usecase"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	postgresDB := db.PostgresConn()
	migrations(postgresDB)

	lis, err := net.Listen("tcp", ":50006")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	ratingUseCase := initUserServer(postgresDB)
	handler.NewServer3(grpcServer, ratingUseCase)

	log.Fatalf("Failed to serve: %v", grpcServer.Serve(lis))
}

func initUserServer(postgresDB *gorm.DB) v1.UseCaseInterface {
	ratingRepo := repo.New(postgresDB)
	ratingUseCase := usecase.New(ratingRepo)
	return ratingUseCase
}

func migrations(sqlDB *gorm.DB) {
	err := sqlDB.Migrator().DropTable(&models.Rating{})
	err1 := sqlDB.Migrator().AutoMigrate(&models.Rating{})
	if err != nil || err1 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migration did run successfully")
	}
}

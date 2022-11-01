package facilities

import (
	"context"
	"log"
	"net/http"

	"github.com/andikem/kraken/pkg/db"
	gorm_generics "github.com/andikem/kraken/pkg/gorm"
	pb "github.com/andikem/kraken/pkg/grpc"
	"github.com/andikem/kraken/pkg/models"
)


type Server struct {
	H   db.Handler
}

func (s *Server) CreateFacility(ctx context.Context, req *pb.CreateFacilityRequest) (*pb.CreateFacilityResponse, error) {
	repository := gorm_generics.NewRepository[models.Facilities, models.GoFacilities](s.H.DB)

	facility := models.GoFacilities{
		
	}

	err := repository.Insert(ctx, &facility)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateFacilityResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateFacilityResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) UpdateFacility(ctx context.Context, req *pb.UpdateFacilityRequest) (*pb.CreateFacilityResponse, error) {
	repository := gorm_generics.NewRepository[models.Facilities, models.GoFacilities](s.H.DB)

	facility := models.GoFacilities{

	}

	err := repository.Update(ctx, &facility)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateFacilityResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateFacilityResponse{
		Status: http.StatusOK,
	}, nil
}

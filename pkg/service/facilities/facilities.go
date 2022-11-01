package facilities

import (
	"context"
	"log"
	"net/http"

	"github.com/andikem/kraken/pkg/db"
	gorm_generics "github.com/andikem/kraken/pkg/gorm"
	pb "github.com/andikem/kraken/pkg/grpc"
	"github.com/andikem/kraken/pkg/models"
	"github.com/google/uuid"
)


type Server struct {
	H   db.Handler
}

func (s *Server) CreateFacility(ctx context.Context, req *pb.CreateFacilityRequest) (*pb.CreateFacilityResponse, error) {
	repository := gorm_generics.NewRepository[models.Facilities, models.GoFacilities](s.H.DB)

	facility := models.GoFacilities{
		Status: req.Status,
		DisableNote: req.DisableNote,
		FacilityName: req.FacilityName,
		IsPort: req.IsPort,
		PortId: uuid.MustParse(req.PortId),
		TypeOfTerminal: req.TypeOfTerminal,
		ThirdPartyServices: req.ThirdPartyServices,
		// FacilityAddress: req.FacilityAddress,
		// FacilityCoordinate: req.FacilityCoordinate,
		EntityId: uuid.MustParse(req.EntityId),
		// FacilityManager: req.FacilityManager,
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
		FacilityId: uuid.MustParse(req.FacilityId),
		Status: req.Status,
		DisableNote: req.DisableNote,
		FacilityName: req.FacilityName,
		IsPort: req.IsPort,
		PortId: uuid.MustParse(req.PortId),
		TypeOfTerminal: req.TypeOfTerminal,
		ThirdPartyServices: req.ThirdPartyServices,
		// FacilityAddress: req.FacilityAddress,
		// FacilityCoordinate: req.FacilityCoordinate,
		EntityId: uuid.MustParse(req.EntityId),
		// FacilityManager: req.FacilityManager,
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

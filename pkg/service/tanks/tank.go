package tanks

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

func (s *Server) CreateTank(ctx context.Context, req *pb.CreateTankRequest) (*pb.CreateTankResponse, error) {
	repository := gorm_generics.NewRepository[models.Tank, models.GoTank](s.H.DB)

	tank := models.GoTank{
		Status: req.Status,
		DisableNote: req.DisableNote,
		TankName: req.TankName,
		// TankDescription: req.TankDescription,
		// TankCoordinate: req.TankCoordinate,
		FacilityId: uuid.MustParse(req.FacilityId),
	}

	err := repository.Insert(ctx, &tank)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateTankResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateTankResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) UpdateTank(ctx context.Context, req *pb.UpdateTankRequest) (*pb.CreateTankResponse, error) {
	repository := gorm_generics.NewRepository[models.Tank, models.GoTank](s.H.DB)

	tank := models.GoTank{
		TankId: uuid.MustParse(req.TankId),
		Status: req.Status,
		DisableNote: req.DisableNote,
		TankName: req.TankName,
		// TankDescription: req.TankDescription,
		// TankCoordinate: req.TankCoordinate,
		FacilityId: uuid.MustParse(req.FacilityId),
	}

	err := repository.Update(ctx, &tank)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateTankResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateTankResponse{
		Status: http.StatusOK,
	}, nil
}

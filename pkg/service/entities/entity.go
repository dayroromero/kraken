package entities

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

func (s *Server) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	repository := gorm_generics.NewRepository[models.Entity, models.GoEntity](s.H.DB)

	entity := models.GoEntity{
		Status: req.Status,
		DisableNote: req.DisableNote,
		EntityName: req.EntityName,
		FriendlyName: req.FriendlyName,
		CountryId: uuid.MustParse(req.CountryId),
		SegmentId: uuid.MustParse(req.SegmentId),
	}

	err := repository.Insert(ctx, &entity)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateEntityResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateEntityResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) UpdateEntity(ctx context.Context, req *pb.UpdateEntityRequest) (*pb.CreateEntityResponse, error) {
	repository := gorm_generics.NewRepository[models.Entity, models.GoEntity](s.H.DB)

	entity := models.GoEntity{
		EntityId: uuid.MustParse(req.EntityId),
		Status: req.Status,
		DisableNote: req.DisableNote,
		EntityName: req.EntityName,
		FriendlyName: req.FriendlyName,
		CountryId: uuid.MustParse(req.CountryId),
		SegmentId: uuid.MustParse(req.SegmentId),
	}

	err := repository.Update(ctx, &entity)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateEntityResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateEntityResponse{
		Status: http.StatusOK,
	}, nil
}

package bol

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

func (s *Server) CreateBol(ctx context.Context, req *pb.CreateBolRequest) (*pb.CreateBolResponse, error) {
	repository := gorm_generics.NewRepository[models.Bol, models.GoBol](s.H.DB)

	bol := models.GoBol{
		Status: req.Status,
		DisableNote: req.DisableNote,
		ShippingETA: req.ShippingETA.AsTime(),
		ShippingETD: req.ShippingETD.AsTime(),
		ArrivalDate: req.ArrivalDate.AsTime(),
		CloseDate: req.CloseDate.AsTime(),
		IsNullandVoid: req.IsNullandVoid,
		OriginalBOL: uuid.MustParse(req.OriginalBOL),
		IsDelivered: req.IsDelivered,
		ShippingLineId: uuid.MustParse(req.ShippingLineId),
		VesselId: uuid.MustParse(req.VesselId),
		Voyage: req.Voyage,
		ProductId: uuid.MustParse(req.ProductId),
		BPReceiverId: uuid.MustParse(req.BPReceiverId),
		Quantity: req.Quantity,
		DestinationPortId: uuid.MustParse(req.DestinationPortId),
	}

	err := repository.Insert(ctx, &bol)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateBolResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateBolResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) UpdateBol(ctx context.Context, req *pb.UpdateBolRequest) (*pb.CreateBolResponse, error) {
	repository := gorm_generics.NewRepository[models.Bol, models.GoBol](s.H.DB)

	bol := models.GoBol{
		ShippingBOLId: uuid.MustParse(req.ShippingBOLId),
		Status: req.Status,
		DisableNote: req.DisableNote,
		ShippingETA: req.ShippingETA.AsTime(),
		ShippingETD: req.ShippingETD.AsTime(),
		ArrivalDate: req.ArrivalDate.AsTime(),
		CloseDate: req.CloseDate.AsTime(),
		IsNullandVoid: req.IsNullandVoid,
		OriginalBOL: uuid.MustParse(req.OriginalBOL),
		IsDelivered: req.IsDelivered,
		ShippingLineId: uuid.MustParse(req.ShippingLineId),
		VesselId: uuid.MustParse(req.VesselId),
		Voyage: req.Voyage,
		ProductId: uuid.MustParse(req.ProductId),
		BPReceiverId: uuid.MustParse(req.BPReceiverId),
		Quantity: req.Quantity,
		DestinationPortId: uuid.MustParse(req.DestinationPortId),
	}

	err := repository.Update(ctx, &bol)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateBolResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateBolResponse{
		Status: http.StatusOK,
	}, nil
}

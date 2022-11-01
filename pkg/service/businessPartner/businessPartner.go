package businesspartner

import (
	"context"
	"encoding/json"
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

func (s *Server) CreateBusinessPartner(ctx context.Context, req *pb.CreateBusinessPartnerRequest) (*pb.CreateBusinessPartnerResponse, error) {
	repository := gorm_generics.NewRepository[models.BussinesPartner, models.GoBussinesPartner](s.H.DB)

	var NewContact models.Contact
	 json.Unmarshal(req.Contact, &NewContact)
	
	businesspartner := models.GoBussinesPartner{
		Status: req.Status,
		DisableNote: req.DisableNote,
		CustomerName: req.CustomerName,
		IsSupplier: req.IsSupplier,
		IsCustomer: req.IsCustomer,
		IsShippingAgent: req.IsShippingAgent,
		IsTrucker: req.IsTrucker,
		IsChild: req.IsChild,
		ParentCustomer: uuid.MustParse(req.ParentCustomer),
	}

	err := repository.Insert(ctx, &businesspartner)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateBusinessPartnerResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateBusinessPartnerResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) UpdateBusinessPartner(ctx context.Context, req *pb.UpdateBusinessPartnerRequest) (*pb.CreateBusinessPartnerResponse, error) {
	repository := gorm_generics.NewRepository[models.BussinesPartner, models.GoBussinesPartner](s.H.DB)

	bp := models.GoBussinesPartner{
		BusinessPartnerId: uuid.MustParse(req.BusinessPartnerId),
		Status: req.Status,
		DisableNote: req.DisableNote,
		CustomerName: req.CustomerName,
		IsSupplier: req.IsSupplier,
		IsCustomer: req.IsCustomer,
		IsShippingAgent: req.IsShippingAgent,
		IsTrucker: req.IsTrucker,
		IsChild: req.IsChild,
		ParentCustomer: uuid.MustParse(req.ParentCustomer),
	}

	err := repository.Update(ctx, &bp)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateBusinessPartnerResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateBusinessPartnerResponse{
		Status: http.StatusOK,
	}, nil
}
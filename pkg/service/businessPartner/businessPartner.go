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
	guuid "github.com/gofrs/uuid"
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

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.CreateProductResponse, error) {
	repository := gorm_generics.NewRepository[models.Product, models.GoProduct](s.H.DB)

	product := models.GoProduct{
		ProductId: 		   guuid.FromStringOrNil(req.ProductId),
		Status:            req.Status,
		DisableNote:       req.DisableNote,
		SKU:               req.SKU,
		ProductName:       req.ProductName,
		ProductDescriptor: req.ProductDescriptor,
		UnCode:            req.UnCode,
		ChrisCode:         req.ChrisCode,
		EinecNumber:       req.EinecNumber,
		HSCode:            req.HSCode,
		ChemUniqueId:      guuid.FromStringOrNil(req.ChemUniqueId),
	}

	err := repository.Update(ctx, &product)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateProductResponse{
			Status: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateProductResponse{
		Status: http.StatusOK,
	}, nil
}
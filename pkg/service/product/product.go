package product

import (
	"context"
	"log"
	"net/http"

	"github.com/andikem/kraken/pkg/db"
	gorm_generics "github.com/andikem/kraken/pkg/gorm"
	pb "github.com/andikem/kraken/pkg/grpc"
	"github.com/andikem/kraken/pkg/models"
	"github.com/gofrs/uuid"
)

type Server struct {
	H   db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	repository := gorm_generics.NewRepository[models.Product, models.GoProduct](s.H.DB)

	product := models.GoProduct{
		Status:            req.Status,
		DisableNote:       req.DisableNote,
		SKU:               req.SKU,
		ProductName:       req.ProductName,
		ProductDescriptor: req.ProductDescriptor,
		UnCode:            req.UnCode,
		ChrisCode:         req.ChrisCode,
		EinecNumber:       req.EinecNumber,
		HSCode:            req.HSCode,
		ChemUniqueId:      uuid.FromStringOrNil(req.ChemUniqueId),
	}

	err := repository.Insert(ctx, &product)
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

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.CreateProductResponse, error) {
	repository := gorm_generics.NewRepository[models.Product, models.GoProduct](s.H.DB)

	product := models.GoProduct{
		ProductId: 		   uuid.FromStringOrNil(req.ProductId),
		Status:            req.Status,
		DisableNote:       req.DisableNote,
		SKU:               req.SKU,
		ProductName:       req.ProductName,
		ProductDescriptor: req.ProductDescriptor,
		UnCode:            req.UnCode,
		ChrisCode:         req.ChrisCode,
		EinecNumber:       req.EinecNumber,
		HSCode:            req.HSCode,
		ChemUniqueId:      uuid.FromStringOrNil(req.ChemUniqueId),
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

func (s *Server) GetAllProducts(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	repository := gorm_generics.NewRepository[models.Product, models.GoProduct](s.H.DB)

	product := models.GoProduct{
		Status:            req.Status,
		DisableNote:       req.DisableNote,
		SKU:               req.SKU,
		ProductName:       req.ProductName,
		ProductDescriptor: req.ProductDescriptor,
		UnCode:            req.UnCode,
		ChrisCode:         req.ChrisCode,
		EinecNumber:       req.EinecNumber,
		HSCode:            req.HSCode,
		ChemUniqueId:      uuid.FromStringOrNil(req.ChemUniqueId),
	}

	err := repository.Insert(ctx, &product)
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

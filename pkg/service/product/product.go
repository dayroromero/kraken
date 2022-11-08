package product

import (
	"context"
	"encoding/json"
	"fmt"
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
		ChemUniqueId:      uuid.MustParse(req.ChemUniqueId),
	}

	err := repository.Insert(ctx, &product)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateProductResponse{
			ReqStatus: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateProductResponse{
		ReqStatus: http.StatusOK,
	}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.CreateProductResponse, error) {
	repository := gorm_generics.NewRepository[models.Product, models.GoProduct](s.H.DB)

	product := models.GoProduct{
		ProductId: 		  uuid.MustParse(req.ProductId),
		Status:            req.Status,
		DisableNote:       req.DisableNote,
		SKU:               req.SKU,
		ProductName:       req.ProductName,
		ProductDescriptor: req.ProductDescriptor,
		UnCode:            req.UnCode,
		ChrisCode:         req.ChrisCode,
		EinecNumber:       req.EinecNumber,
		HSCode:            req.HSCode,
		ChemUniqueId:      uuid.MustParse(req.ChemUniqueId),
	}

	err := repository.Update(ctx, &product)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateProductResponse{
			ReqStatus: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateProductResponse{
		ReqStatus: http.StatusOK,
	}, nil
}

func (s *Server) GetAllProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	//repository := gorm_generics.NewRepository[models.Product, models.GoProduct](s.H.DB)

	

	//products, err := repository.Find(ctx, &filters)

	filters, err := json.Marshal(req.Filters)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(filters)



	return &pb.GetProductsResponse{
		
		ReqStatus: http.StatusOK,
	}, nil
}

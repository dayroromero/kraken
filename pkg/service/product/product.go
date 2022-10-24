package product

import (
	"context"
	"log"
	"net/http"

	"github.com/andikem/kraken/pkg/config"
	"github.com/andikem/kraken/pkg/db"
	gorm_generics "github.com/andikem/kraken/pkg/gorm"
	pb "github.com/andikem/kraken/pkg/grpc"
	"github.com/andikem/kraken/pkg/models"
	"github.com/gofrs/uuid"
)

type Server struct {
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	env, err := config.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	dbUrl, err := config.GetAwsSecret(env)
	if err != nil {
		log.Fatal(err)
	}

	db := db.Init(dbUrl)

	repository := gorm_generics.NewRepository[models.Product, models.GoProduct](db.DB)

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

	err = repository.Insert(ctx, &product)
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

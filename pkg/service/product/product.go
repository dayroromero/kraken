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

	repository := gorm_generics.NewRepository[models.ProductGorm, models.Product](db.DB)

	product := models.Product{
		Status:            true,
		DisableNote:       "",
		SKU:               "",
		ProductName:       "Test",
		ProductDescriptor: "Test",
		UnCode:            "",
		ChrisCode:         "",
		EinecNumber:       "",
		HSCode:            "",
		ChemUniqueId:      uuid.UUID{},
	}

	err = repository.Insert(ctx, &product)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.CreateProductResponse{
		Status: http.StatusOK,
	}, nil
}


func (s *Server) UpsertChemical(ctx context.Context, req *pb.UpserChemicalRequest) (*pb.UpserChemicalResponse, error) {

	return &pb.UpserChemicalResponse{
		Status: http.StatusOK,
	}, nil
}
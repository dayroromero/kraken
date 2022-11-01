package routes

import (
	"log"

	"github.com/andikem/kraken/pkg/config"
	"github.com/andikem/kraken/pkg/db"
	pb "github.com/andikem/kraken/pkg/grpc"
	businesspartner "github.com/andikem/kraken/pkg/service/businessPartner"
	"github.com/andikem/kraken/pkg/service/facilities"
	"github.com/andikem/kraken/pkg/service/product"

	"google.golang.org/grpc"
)

func RegisterRoutes(grpcServer *grpc.Server) {
	env, err := config.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	dbUrl, err := config.GetAwsSecret(env)
	if err != nil {
		log.Fatal(err)
	}

	h := db.Init(dbUrl)

	s1 := product.Server{
		H: h,
	}

	pb.RegisterProductsServiceServer(grpcServer, &s1)

	s2 := businesspartner.Server{
		H: h,
	}

	pb.RegisterBusinessPartnerServiceServer(grpcServer, &s2)

	s3 := facilities.Server{
		H: h,
	}

	pb.RegisterFacilitiesServiceServer(grpcServer, &s3)
}

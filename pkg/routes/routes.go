package routes

import (
	"log"

	"github.com/andikem/kraken/pkg/config"
	"github.com/andikem/kraken/pkg/db"
	pb "github.com/andikem/kraken/pkg/grpc"
	"github.com/andikem/kraken/pkg/service/bol"
	businesspartner "github.com/andikem/kraken/pkg/service/businessPartner"
	"github.com/andikem/kraken/pkg/service/entities"
	"github.com/andikem/kraken/pkg/service/facilities"
	"github.com/andikem/kraken/pkg/service/product"
	"github.com/andikem/kraken/pkg/service/tanks"

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

	// Products Routes
	s1 := product.Server{
		H: h,
	}

	pb.RegisterProductsServiceServer(grpcServer, &s1)

	// Business Partner Routes
	s2 := businesspartner.Server{
		H: h,
	}

	pb.RegisterBusinessPartnerServiceServer(grpcServer, &s2)

	//Facilities Routes
	s3 := facilities.Server{
		H: h,
	}

	pb.RegisterFacilitiesServiceServer(grpcServer, &s3)

	//Entities Routes
	s4 := entities.Server{
		H: h,
	}

	pb.RegisterEntitiesServiceServer(grpcServer, &s4)

	//Tanks Routes
	s5 := tanks.Server{
		H: h,
	}

	pb.RegisterTankServiceServer(grpcServer, &s5)

	// Bol Routes
	s6 := bol.Server{
		H: h,
	}

	pb.RegisterShippingBOLServiceServer(grpcServer, &s6)

	// FacilityAssets Routes
	s7 := facilities.FServer{
		H: h,
	}

	pb.RegisterFacilityAssetServiceServer(grpcServer, &s7)
}

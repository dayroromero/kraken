package facilities

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


type FServer struct {
	H   db.Handler
}

func (s *FServer) CreateFacilityAsset(ctx context.Context, req *pb.CreateFacilityAssetRequest) (*pb.CreateFacilityAssetResponse, error) {
	repository := gorm_generics.NewRepository[models.FacilityAsset, models.GoFacilityAsset](s.H.DB)

	facility := models.GoFacilityAsset{
		Enabled : req.Enabled,
		Name : req.Name,
		Description : req.Description,
		InventoryId : req.InventoryId,
		AssetType : req.AssetType,
		FacilityId : uuid.MustParse(req.FacilityId),
	}

	err := repository.Insert(ctx, &facility)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateFacilityAssetResponse{
			ReqStatus: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateFacilityAssetResponse{
		ReqStatus: http.StatusOK,
	}, nil
}

func (s *FServer) UpdateFacilityAsset(ctx context.Context, req *pb.UpdateFacilityAssetRequest) (*pb.CreateFacilityAssetResponse, error) {
	repository := gorm_generics.NewRepository[models.FacilityAsset, models.GoFacilityAsset](s.H.DB)

	facility := models.GoFacilityAsset{
		FacilityAssetId: uuid.MustParse(req.FacilityAssetId),
		Enabled : req.Enabled,
		Name : req.Name,
		Description : req.Description,
		InventoryId : req.InventoryId,
		AssetType : req.AssetType,
		FacilityId : uuid.MustParse(req.FacilityId),
	}

	err := repository.Update(ctx, &facility)
	if err != nil {
		log.Fatal(err)

		return &pb.CreateFacilityAssetResponse{
			ReqStatus: http.StatusBadGateway,
		}, err
	}

	return &pb.CreateFacilityAssetResponse{
		ReqStatus: http.StatusOK,
	}, nil
}

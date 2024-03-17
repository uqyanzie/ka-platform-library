package gapi

import (
	"context"
	pb "github.com/uqyanzie/platform-library-grpc-api/pb"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlatformServer) GetPlatformList(ctx context.Context, req *pb.PlatformListRequest) (*pb.PlatformListResponse, error) {
	platforms, err := s.platformService.GetPlatformList()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get platform list: %v", err)
	}

	platformList := &pb.PlatformListResponse{}

	for _, platform := range platforms {
		platformList.Platforms = append(platformList.Platforms, &pb.Platform{
			Id:              platform.Id,
			UnitName:        platform.UnitName,
			UnitClass:       platform.UnitClass,
			OperationField:  platform.OperationField,
			GeneralCategory: platform.GeneralCategory,
			GeneralType:     platform.GeneralType,
			LethalityLevel:  platform.LethalityLevel,
			CruisingSpeed:   platform.CruisingSpeed,
			MaxSpeed:        platform.MaxSpeed,
			CombatRange:     platform.CombatRange,
			FuelLoad:        platform.FuelLoad,
			CountryOrigin:   platform.CountryOrigin,
		})
	}

	return platformList, nil
}


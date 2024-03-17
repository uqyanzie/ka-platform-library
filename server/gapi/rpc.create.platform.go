package gapi

import (
	"context"

	"github.com/uqyanzie/platform-library-grpc-api/models"
	pb "github.com/uqyanzie/platform-library-grpc-api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlatformServer) CreateNewPlatform(ctx context.Context, req *pb.NewPlatform) (*pb.Platform, error) {

	platform := &models.NewPlatform{
		UnitName:        req.GetUnitName(),
		UnitClass:       req.GetUnitClass(),
		OperationField:  req.GetOperationField(),
		GeneralCategory: req.GetGeneralCategory(),
		GeneralType:     req.GetGeneralType(),
		LethalityLevel:  req.GetLethalityLevel(),
		CruisingSpeed:   req.GetCruisingSpeed(),
		MaxSpeed:        req.GetMaxSpeed(),
		CombatRange:     req.GetCombatRange(),
		FuelLoad:        req.GetFuelLoad(),
		CountryOrigin:   req.GetCountryOrigin(),
	}

	res, err := s.platformService.CreateNewPlatform(platform)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create new platform: %v", err)
	}

	return &pb.Platform{
		Id:              res.Id,
		UnitName:        res.UnitName,
		UnitClass:       res.UnitClass,
		OperationField:  res.OperationField,
		GeneralCategory: res.GeneralCategory,
		GeneralType:     res.GeneralType,
		LethalityLevel:  res.LethalityLevel,
		CruisingSpeed:   res.CruisingSpeed,
		MaxSpeed:        res.MaxSpeed,
		CombatRange:     res.CombatRange,
		FuelLoad:        res.FuelLoad,
		CountryOrigin:   res.CountryOrigin,
	}, nil

}

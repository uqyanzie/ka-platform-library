package service

import (
	"context"
	"log"
	
	pb "grpc-platform/pb/platform"

	"github.com/google/uuid"
)

type PlatformServer struct {
	pb.UnimplementedPlatformServiceServer
}

func (s *PlatformServer) CreateNewPlatform(ctx context.Context, in *pb.NewPlatform) (*pb.Platform, error) {

	log.Printf("Received: %v", in.GetUnitName())
	platform := &pb.Platform{
		UnitName:        in.GetUnitName(),
		UnitClass:       in.GetUnitClass(),
		OperationField:  in.GetOperationField(),
		GeneralCategory: in.GetGeneralCategory(),
		GeneralType:     in.GetGeneralType(),
		LethalityLevel:  in.GetLethalityLevel(),
		CruisingSpeed:   in.GetCruisingSpeed(),
		MaxSpeed:        in.GetMaxSpeed(),
		CombatRange:     in.GetCombatRange(),
		FuelLoad:        in.GetFuelLoad(),
		CountryOrigin:   in.GetCountryOrigin(),
		Id:              uuid.New().String(),
	}

	return platform, nil

}
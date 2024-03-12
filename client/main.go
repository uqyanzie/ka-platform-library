package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-platform/pb/platform"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDRESS = "localhost:8080"
)

type PlatformUnit struct {
	Name            string
	Description     string
	Done            bool
	UnitName        string
	UnitClass       string
	OperationField  string
	GeneralCategory string
	GeneralType     string
	LethalityLevel  string
	CruisingSpeed   float32
	MaxSpeed        float32
	CombatRange     float32
	FuelLoad        float32
	CountryOrigin   string
}

func main() {
	
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	log.Printf("Connected to server %s", ADDRESS)

	defer conn.Close()

	c := pb.NewPlatformServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	platforms := []PlatformUnit{
		{
			UnitName:        "KRI Usman Harun",
			UnitClass:       "Bung Tomo Class",
			OperationField:  "Surface",
			GeneralCategory: "Surface Combatant",
			GeneralType:     "Corvettes",
			LethalityLevel:  "Medium",
			CruisingSpeed:   14.77,
			MaxSpeed:        24.27,
			CombatRange:     50.57,
			FuelLoad:        200,
			CountryOrigin:   "Indonesia",
		},
	}

	for _, platform := range platforms {
		res, err := c.CreateNewPlatform(ctx, &pb.NewPlatform{
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

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`
           ID : %s
           Unit Name : %s
           Unit Class : %s
           Operating Field : %v,
       `, res.GetId(), res.GetUnitName(), res.GetUnitClass(), res.GetOperationField())
	}

}

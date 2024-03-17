package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/uqyanzie/platform-library-grpc-api/models"
)

type PlatformRepository interface {
	CreateNewPlatform(platform *models.NewPlatform) (*models.Platform, error)
	GetPlatformList() ([]*models.Platform, error)
}
type PlatformNeo4jRepository struct {
	ctx    context.Context
	Driver neo4j.DriverWithContext
}

func NewPlatformRepository(ctx context.Context, driver neo4j.DriverWithContext) PlatformRepository {
	return &PlatformNeo4jRepository{ctx: ctx, Driver: driver}
}

func (pl *PlatformNeo4jRepository) CreateNewPlatform(platform *models.NewPlatform) (*models.Platform, error) {
	newPlatform := &models.Platform{
		Id:              uuid.New().String(),
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
	}

	session := pl.Driver.NewSession(pl.ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer session.Close(pl.ctx)

	_, err := session.ExecuteWrite(pl.ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := "CREATE (p:Platform {id: $id, unitName: $unitName, unitClass: $unitClass, operationField: $operationField, generalCategory: $generalCategory, generalType: $generalType, lethalityLevel: $lethalityLevel, cruisingSpeed: $cruisingSpeed, maxSpeed: $maxSpeed, combatRange: $combatRange, fuelLoad: $fuelLoad, countryOrigin: $countryOrigin})"

		params := map[string]interface{}{
			"id":              newPlatform.Id,
			"unitName":        newPlatform.UnitName,
			"unitClass":       newPlatform.UnitClass,
			"operationField":  newPlatform.OperationField,
			"generalCategory": newPlatform.GeneralCategory,
			"generalType":     newPlatform.GeneralType,
			"lethalityLevel":  newPlatform.LethalityLevel,
			"cruisingSpeed":   newPlatform.CruisingSpeed,
			"maxSpeed":        newPlatform.MaxSpeed,
			"combatRange":     newPlatform.CombatRange,
			"fuelLoad":        newPlatform.FuelLoad,
			"countryOrigin":   newPlatform.CountryOrigin,
		}

		result, err := tx.Run(pl.ctx, query, params)

		if err != nil {
			return nil, err
		}

		return result, nil
	})

	if err != nil {
		return nil, err
	}

	return newPlatform, nil

}

func (pl *PlatformNeo4jRepository) GetPlatformList() ([]*models.Platform, error) {
	session := pl.Driver.NewSession(pl.ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})

	defer session.Close(pl.ctx)

	query := "MATCH (p:Platform) RETURN p"
	res, err := session.Run(pl.ctx, query, map[string]interface{}{})

	if err != nil {
		return nil, err
	}

	var platforms []*models.Platform

	for res.Next(pl.ctx) {
		record := res.Record()
		if value, ok := record.Get("p"); ok {
			node := value.(neo4j.Node)
			props := node.Props

			platform := &models.Platform{}

			err := mapstructure.Decode(props, platform)
			if err != nil {
				return nil, err
			}

			platforms = append(platforms, platform)
		}
	}

	return platforms, nil
}

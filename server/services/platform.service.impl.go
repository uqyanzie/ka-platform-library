package services

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/uqyanzie/platform-library-grpc-api/models"
	"github.com/uqyanzie/platform-library-grpc-api/repository"
)

type PlatformServiceImpl struct {
	ctx                context.Context
	driver             neo4j.DriverWithContext
	platformRepository repository.PlatformRepository
}

func NewPlatformService(ctx context.Context, driver neo4j.DriverWithContext) PlatformService {
	return &PlatformServiceImpl{ctx: ctx, driver: driver, platformRepository: repository.NewPlatformRepository(ctx, driver)}
}

func (pl *PlatformServiceImpl) CreateNewPlatform(platform *models.NewPlatform) (*models.Platform, error) {

	res, err := pl.platformRepository.CreateNewPlatform(platform)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pl *PlatformServiceImpl) GetPlatformList() ([]*models.Platform, error) {
	return nil, nil
}

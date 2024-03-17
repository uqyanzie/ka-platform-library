package services

import "github.com/uqyanzie/platform-library-grpc-api/models"

type PlatformService interface {
	CreateNewPlatform(platform *models.NewPlatform) (*models.Platform, error)
	GetPlatformList() ([]*models.Platform, error)
}

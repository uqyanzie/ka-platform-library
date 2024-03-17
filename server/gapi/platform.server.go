package gapi

import (
	pb "github.com/uqyanzie/platform-library-grpc-api/pb"
	"github.com/uqyanzie/platform-library-grpc-api/repository"
	"github.com/uqyanzie/platform-library-grpc-api/services"
)

type PlatformServer struct {
	pb.UnimplementedPlatformServiceServer
	platformRepository repository.PlatformRepository
	platformService    services.PlatformService
}

func NewGrpcPlatformServer(platformRepository repository.PlatformRepository, platformService services.PlatformService) (*PlatformServer, error) {
	platformServer := &PlatformServer{
		platformRepository: platformRepository,
		platformService:    platformService,
	}

	return platformServer, nil
}

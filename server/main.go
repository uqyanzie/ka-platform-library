package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/uqyanzie/platform-library-grpc-api/config"
	pb "github.com/uqyanzie/platform-library-grpc-api/pb"
	"github.com/uqyanzie/platform-library-grpc-api/repository"
	"google.golang.org/grpc"

	"github.com/uqyanzie/platform-library-grpc-api/services"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/uqyanzie/platform-library-grpc-api/controllers"
	"github.com/uqyanzie/platform-library-grpc-api/gapi"
	"github.com/uqyanzie/platform-library-grpc-api/routes"
	"google.golang.org/grpc/reflection"
)

var (
	server *gin.Engine
	ctx    context.Context

	// userService         services.UserService
	// UserController      controllers.UserController
	// UserRouteController routes.UserRouteController

	// authCollection      *mongo.Collection
	// authService         services.AuthService
	// AuthController      controllers.AuthController
	// AuthRouteController routes.AuthRouteController

	// // 👇 Create the Post Variables
	platformService         services.PlatformService
	PlatformController      controllers.PlatformController
	platformRepository      repository.PlatformRepository
	PlatformRouteController routes.PlatformRouteController
)

func connectNeo4j() (neo4j.DriverWithContext, error) {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	driver, err := neo4j.NewDriverWithContext(config.NEO4J_URI, neo4j.BasicAuth(config.NEO4J_USER, config.NEO4J_PASS, ""))
	if err != nil {
		log.Fatal("cannot create neo4j driver: ", err)
		panic(err)
	}

	//defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatal("cannot connect to neo4j: ", err)
		panic(err)
	}

	return driver, err
}

func init() {

	ctx = context.Background()

	driver, err := connectNeo4j()

	if err != nil {
		log.Fatal("cannot connect to neo4j: ", err)
		panic(err)
	}

	fmt.Println("Neo4j successfully connected...")

	platformService = services.NewPlatformService(ctx, driver)
	platformRepository = repository.NewPlatformRepository(ctx, driver)
	PlatformController = controllers.NewPlatformController(platformService)
	PlatformRouteController = routes.NewPlatformRouteController(PlatformController)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	startGrpcServer(config)
}

func startGrpcServer(config config.Config) {

	platformServer, err := gapi.NewGrpcPlatformServer(platformService, platformRepository)
	
	if err != nil {
		log.Fatal("cannot create grpc platformServer: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPlatformServiceServer(grpcServer, platformServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}

func startGinServer(config config.Config) {
	// value, err := redisclient.Get(ctx, "test").Result()

	// if err == redis.Nil {
	// 	fmt.Println("key: test does not exist")
	// } else if err != nil {
	// 	panic(err)
	// }

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	// router.GET("/healthchecker", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	// })

	// AuthRouteController.AuthRoute(router, userService)
	// UserRouteController.UserRoute(router, userService)
	// // 👇 Post Route
	// PostRouteController.PostRoute(router)

	PlatformRouteController.PlatformRoutes(router)
	log.Fatal(server.Run(":" + config.Port))
}

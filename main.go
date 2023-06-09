package main

import (
	"usermanagement/infrastructure"
	"usermanagement/infrastructure/hasher"
	"usermanagement/infrastructure/usertokengenerator"
	"usermanagement/repository"
	"usermanagement/repository/jsondb"
	"usermanagement/service"
	"usermanagement/transport"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Setup infrastructure
	infra := infrastructure.NewInfrastructure()
	infra.SetHasher(hasher.NewDefaultHasher())
	infra.SetUserTokenGenerator(usertokengenerator.NewJWT())

	// Setup repository
	repo := repository.NewRepository()
	repo.SetUser(jsondb.NewUserRepository())

	// Setup service
	svc := service.NewUserImpl(infra, repo)

	// Setup transport
	tpr := transport.NewController(svc)
	fbr := fiber.New(
		fiber.Config{
			ErrorHandler: transport.ErrHandler,
		},
	)
	tpr.Start(fbr)
}

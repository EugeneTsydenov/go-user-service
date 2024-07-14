package app

import (
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/internal/adapter/grpc_adapter"
	"github.com/EugeneTsydenov/go-user-service/internal/domain/repository"
	"github.com/EugeneTsydenov/go-user-service/internal/service"
	"github.com/EugeneTsydenov/go-user-service/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type serviceProvider struct {
	userService        *service.Service
	userRepository     *repository.Repository
	grpcConfig         string
	userImplementation *grpc_adapter.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

//Service provider getters:

func (provider *serviceProvider) UserService() *service.Service {
	if provider.userService == nil {
		userService := service.New(provider.UserRepository())
		err := provider.SetUserService(userService)
		if err != nil {
			return nil
		}
	}

	return provider.userService
}

func (provider *serviceProvider) UserRepository() *repository.Repository {
	if provider.userRepository == nil {
		db, err := provider.DB()
		if err != nil {
			return nil
		}
		userRepository := repository.New(db)
		err = provider.SetUserRepository(userRepository)
		if err != nil {
			return nil
		}
	}

	return provider.userRepository
}

func (provider *serviceProvider) GrpcConfig() string {
	return provider.grpcConfig
}

func (provider *serviceProvider) UserImplementation() *grpc_adapter.Implementation {
	if provider.userImplementation == nil {
		userImplementation := grpc_adapter.NewImplementation(provider.UserService())
		err := provider.SetUserImplementation(userImplementation)
		if err != nil {
			return nil
		}
	}
	return provider.userImplementation
}

//Service provider setters:

func (provider *serviceProvider) SetUserService(userService *service.Service) error {
	if userService == nil {
		return fmt.Errorf("%v: userService can't be empty", userService)
	}
	provider.userService = userService
	return nil
}

func (provider *serviceProvider) SetUserRepository(userRepository *repository.Repository) error {
	if userRepository == nil {
		return fmt.Errorf("%v: userRepository can't be empty", userRepository)
	}
	provider.userRepository = userRepository
	return nil
}

func (provider *serviceProvider) SetUserImplementation(userImplementation *grpc_adapter.Implementation) error {
	if userImplementation == nil {
		return fmt.Errorf("%v: userImplementation can't be empty", userImplementation)
	}
	provider.userImplementation = userImplementation
	return nil
}

func (provider *serviceProvider) SetGrpcConfig(grpcConfig string) error {
	if grpcConfig == "" {
		return fmt.Errorf("%v: grpcConfig can't be empty", grpcConfig)
	}
	return nil
}

func (provider *serviceProvider) DB() (*gorm.DB, error) {
	host := pkg.GetEnv()["DB_HOST"]
	user := pkg.GetEnv()["DB_USERNAME"]
	password := pkg.GetEnv()["DB_PASSWORD"]
	dbName := pkg.GetEnv()["DB_NAME"]
	port := pkg.GetEnv()["DB_PORT"]
	sslMode := pkg.GetEnv()["DB_SSL_MODE"]

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}
	fmt.Println("Connected to database successfully!")

	return db, nil
}

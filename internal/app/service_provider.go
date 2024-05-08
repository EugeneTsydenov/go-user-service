package app

import "fmt"

type serviceProvider struct {
	userService        string
	userRepository     string
	grpcConfig         string
	userImplementation string
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

//Service provider getters:

func (provider *serviceProvider) UserService() string {
	return provider.userService
}

func (provider *serviceProvider) UserRepository() string {
	return provider.userRepository
}

func (provider *serviceProvider) GrpcConfig() string {
	return provider.grpcConfig
}

func (provider *serviceProvider) UserImplementation() string {
	return provider.userImplementation
}

//Service provider setters:

func (provider *serviceProvider) SetUserService(userService string) error {
	if userService == "" {
		return fmt.Errorf("%v: userService can't be empty", userService)
	}
	return nil
}

func (provider *serviceProvider) SetUserRepository(userRepository string) error {
	if userRepository == "" {
		return fmt.Errorf("%v: userRepository can't be empty", userRepository)
	}
	return nil
}

func (provider *serviceProvider) SetUserImplementation(userImplementation string) error {
	if userImplementation == "" {
		return fmt.Errorf("%v: userImplementation can't be empty", userImplementation)
	}
	return nil
}

func (provider *serviceProvider) SetGrpcConfig(grpcConfig string) error {
	if grpcConfig == "" {
		return fmt.Errorf("%v: grpcConfig can't be empty", grpcConfig)
	}
	return nil
}

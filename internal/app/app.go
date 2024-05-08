package app

import (
	"context"
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
	"github.com/EugeneTsydenov/go-user-service/pkg"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type App struct {
	grpcServer      *grpc.Server
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) Serve() error {
	servicePort := pkg.GetEnv()["SERVICE_PORT"]
	listener, err := net.Listen("tcp", servicePort)
	if err != nil {
		log.Fatalf("Error listening on port %s", servicePort)
	}
	grpcServer := app.GrpcServer()
	fmt.Printf("Grpc server listening on port %s", servicePort)
	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	provider := newServiceProvider()
	err := app.SetServiceProvider(provider)
	if err != nil {
		return err
	}
	err = app.initGrpcServer(ctx)
	if err != nil {
		return err
	}
	err = app.initEnv(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) initGrpcServer(_ context.Context) error {
	err := app.SetGrpcServer(grpc.NewServer())
	if err != nil {
		return err
	}
	reflection.Register(app.grpcServer)
	proto.RegisterUserServiceServer(app.grpcServer, app.serviceProvider.UserImplementation())
	return nil
}

func (app *App) initEnv(_ context.Context) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

//App getters:

func (app *App) GrpcServer() *grpc.Server {
	return app.grpcServer
}

func (app *App) ServiceProvider() *serviceProvider {
	return app.serviceProvider
}

//App setters:

func (app *App) SetGrpcServer(server *grpc.Server) error {
	if server == nil {
		return fmt.Errorf("%v: userService can't be null", server)
	}
	app.grpcServer = server
	return nil
}

func (app *App) SetServiceProvider(serviceProvider *serviceProvider) error {
	if serviceProvider == nil {
		return fmt.Errorf("%v: serviceProvider can't be null", serviceProvider)
	}
	app.serviceProvider = serviceProvider
	return nil
}

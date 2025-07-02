package app

type App struct {
	GrpcServer *gRPCServer
}

func NewApp(addr string) *App {
	logger := initLogger()
	grpcServer := newGRPCServer(addr, logger)

	return &App{
		GrpcServer: grpcServer,
	}
}

func (a *App) Run() error {
	return a.GrpcServer.Run()
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

package server

type Server interface {
	Start() error
	// TODO: Shutdown(ctx context.Context)
}

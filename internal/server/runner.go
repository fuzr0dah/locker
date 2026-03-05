package server

import "context"

type Runner struct {
	servers []Server
	errChan chan error
}

func NewRunner(servers ...Server) *Runner {
	return &Runner{
		servers: servers,
		errChan: make(chan error, len(servers)),
	}
}

func (r *Runner) Start() {
	for _, srv := range r.servers {
		go func(s Server) {
			if err := s.Start(); err != nil {
				r.errChan <- err
			}
		}(srv)
	}
}

func (r *Runner) Wait() <-chan error {
	return r.errChan
}

func (r *Runner) Shutdown(ctx context.Context) {
	for _, srv := range r.servers {
		srv.Shutdown(ctx)
	}
}

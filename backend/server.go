package backend

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server
type Server struct {
	Addr            string
	FrontendHandler http.Handler

	router     chi.Router
	httpServer http.Server

	chatService *chatService

	cancel context.CancelFunc
}

func (srv *Server) Init() error {
	var router = chi.NewRouter()

	*srv = Server{
		Addr:            srv.Addr,
		FrontendHandler: srv.FrontendHandler,

		router: router,
		httpServer: http.Server{
			Addr:    srv.Addr,
			Handler: router,
		},

		chatService: &chatService{srv},
	}

	if err := srv.setupRoutes(router); err != nil {
		return err
	}

	return nil
}

func (srv *Server) setupRoutes(r chi.Router) error {
	r.Use(middleware.Logger)

	r.Mount("/debug", middleware.Profiler())

	r.Mount("/", srv.FrontendHandler)

	var api = chi.NewRouter()
	r.Mount("/api", http.StripPrefix("/api", api))

	api.Mount(srv.chatService.asHandler())

	return nil
}

func (srv *Server) Run(ctx context.Context) error {
	ctx, srv.cancel = context.WithCancel(ctx)

	var (
		httpServer = &srv.httpServer
		httpErrc   = make(chan error, 1)
	)

	httpServer.BaseContext = func(net.Listener) context.Context { return ctx }

	go func() {
		defer close(httpErrc)
		httpErrc <- httpServer.ListenAndServe()
	}()

	for {
		select {
		case <-ctx.Done():
			return nil

		case err := <-httpErrc:
			return fmt.Errorf("http server: %w", err)
		}
	}

	return nil
}

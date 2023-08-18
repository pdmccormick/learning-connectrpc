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

	subs map[chan<- string]struct{}

	broadcastc chan string
	subc       chan chan<- string
	unsubc     chan chan<- string
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

		subs: make(map[chan<- string]struct{}),

		broadcastc: make(chan string, 10),
		subc:       make(chan chan<- string, 1),
		unsubc:     make(chan chan<- string, 1),
	}

	if err := srv.chatService.Init(); err != nil {
		return err
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

		case msg := <-srv.broadcastc:
			for ch, _ := range srv.subs {
				select {
				case ch <- msg:
				default:
					// backlogged
				}
			}

		case ch := <-srv.subc:
			srv.subs[ch] = struct{}{}

		case ch := <-srv.unsubc:
			delete(srv.subs, ch)
			close(ch)
		}
	}

	return nil
}

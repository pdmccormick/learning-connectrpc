package main

import (
	"context"
	"flag"
	"log"

	"github.com/pdmccormick/learning-connectrpc/backend"
	"github.com/pdmccormick/learning-connectrpc/frontend"
)

func main() {
	var (
		httpFlag   = flag.String("http", ":8080", "run server at http://`addr:port`")
		htdocsFlag = flag.String("htdocs", "", "override bundled frontend assets with local path")
	)

	flag.Parse()

	var (
		ctx      = context.Background()
		frontend = frontend.Handler{
			PathOverride: *htdocsFlag,
		}
		srv = backend.Server{
			Addr:            *httpFlag,
			FrontendHandler: &frontend,
		}
	)

	if err := frontend.Init(); err != nil {
		log.Fatalf("Frontend Init: %s", err)
	}

	if err := srv.Init(); err != nil {
		log.Fatalf("Server Init: %s", err)
	}

	if err := srv.Run(ctx); err != nil {
		log.Fatalf("Server Run: %s", err)
	}
}

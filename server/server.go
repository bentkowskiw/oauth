package server

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run starts the http server
func Run(cfg configer, h http.Handler, closers []Closer) {
	// Create new server
	u := cfg.BindURL()

	// Initialize http server
	srv, err := newServer(*u, h, closers, cfg)
	if err != nil {
		log.Println("NewServer", err)
		os.Exit(0)
	}
	// Start http server
	srv.run(*u)
}

// newServer configured to serve API
func newServer(host url.URL, r http.Handler, closers []Closer, cfg configer) (*Server, error) {
	addr := host.Host

	srv := http.Server{
		Addr:        addr,
		Handler:     r,
		IdleTimeout: time.Second, // request idle timeout, this parameter avoid too many goroutines open
	}

	return &Server{
		configer: cfg,
		Server:   &srv,
		closers:  closers,
	}, nil
}

// Run starts the Listener
func (srv *Server) run(hostname url.URL) {
	defer func() {
		for _, closer := range srv.closers {
			closer.Close()
		}
	}()

	go func() {
		log.Printf("Starting API on %s", srv.Addr)
		err := srv.ListenAndServeTLS(srv.CertPath(), srv.KeyPath())
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 10)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	err := srv.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}

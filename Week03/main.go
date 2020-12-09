package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	var srv1 = http.Server{Addr: ":8080"}
	var srv2 = http.Server{Addr: ":8081"}

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		log.Println("server 1 is running")
		if err := srv1.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		log.Println("server 2 is running")
		//return errors.New("server 2 error")
		if err := srv2.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	select {
	case <-quit:
		log.Println("receive signal")
	case <-ctx.Done():
		log.Println("context done")
	}


	log.Println("shutting down server...")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv1.Shutdown(timeoutCtx); err != nil {
		log.Fatal("force shutdown server2 failed", err)
	}

	if err := srv2.Shutdown(timeoutCtx); err != nil {
		log.Fatal("force shutdown server2 failed", err)
	}

	if err := g.Wait(); err != nil {
		log.Println(err)
	}

	log.Println("server existing")
}

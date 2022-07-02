package week03

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestServerLife(t *testing.T) {
	g, ctx := errgroup.WithContext(context.Background())

	// app server
	g.Go(func() error {
		return serveApp(ctx)
	})

	// debug server
	g.Go(func() error {
		return serveDebug(ctx)
	})

	time.Sleep(time.Minute * 1)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGINT)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sig := <-c:
				// ignore sigquit
				fmt.Printf("signal get: %v\n", sig.String())
				if sig.String() == "quit" {
					fmt.Printf("signal get: %v, continue\n", sig)
					continue
				}

				// return err, shutdown server
				return fmt.Errorf("recieve signal %v", sig)
			}
		}
	})
	if err := g.Wait(); err != nil {
		log.Fatal("Error: ", err)
	}
	time.Sleep(time.Second)
}

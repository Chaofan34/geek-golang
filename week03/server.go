package week03

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func serveApp(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Welcome app server!")
	})
	return serve("0.0.0.0:8080", mux, ctx)
}

func serveDebug(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Welcome debug server!")
	})
	return serve("0.0.0.0:8081", mux, ctx)
}

func serve(addr string, handler http.Handler, ctx context.Context) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-ctx.Done()
		log.Printf("stop %v server!!!\n", addr)
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("shutdown %v err: v%\n", addr, err)
		}
	}()

	// # mock start failed
	if strings.Contains(addr, "8080") {
		return fmt.Errorf("server %v start err", addr)
	}

	time.Sleep(time.Second)
	log.Printf("serve listenAndServer!!!!, %v\n", addr)
	return s.ListenAndServe()
}

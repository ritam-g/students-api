package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/ritam-g/students-api/internal"
)

func main() {
	// load config

	cfg := config.MustLoad()
	//db setup
	//router setup
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hellow this is great api "))
	})
	// server setup

	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	fmt.Printf("server  serted on port no %s", cfg.HttpServer.Address)
	done := make(chan os.Signal,1)

	signal.Notify(done , os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("fails to log server", err.Error())
		}
	}()

	<- done

	slog.Info("shutting down the server ")

	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err:= server.Shutdown(ctx)

	if err!= nil{
		slog.Error("faild to shutdown server ",slog.String("Error ",err.Error()))
	}

	slog.Info("server shut down gracfully ")
}

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

	"github.com/shivang2003/students-api/internal/config"
	"github.com/shivang2003/students-api/internal/http/handlers/student"
	"github.com/shivang2003/students-api/storage/sqlite"
)

func main() {

	//load the config
	cfg := config.MustLoad()

	//db setup

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal("failed to initialize storage REASON: ", err.Error())
	}

	slog.Info("Storage initialized successfully", slog.String("env", cfg.Env))

	//setup router
	router := http.NewServeMux()

	//route apis
	//dependency injection can be done here by passing the required dependencies to the handler functions

	router.HandleFunc("POST /api/v1/students", student.New(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	//graceful shutdown

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("Started server at ", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Printf("Server started at %s", cfg.HTTPServer.Addr)

	go func() { //separate goroutine to start the server and listen for incoming requests which allows the main goroutine to listen for shutdown signals and perform graceful shutdown when needed
		err := server.ListenAndServe() //this
		if err != nil {
			log.Fatal("failed to start server", err.Error())
		}
	}()

	<-done

	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", "error", err.Error())
	}

	slog.Info("Server stopped gracefully")
}

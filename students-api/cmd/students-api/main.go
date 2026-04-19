package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shivang2003/students-api/internal/config"
)

func main() {

	//load the config
	cfg := config.MustLoad()

	//db setup

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	//setup server
	server := http.Server{
		Addr:     cfg.HTTPServer.Addr,
		Handler: router,
	}
	fmt.Printf("Server started at %s", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal("failed to start server")
	}

	
}

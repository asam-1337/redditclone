package main

import (
	"github.com/asam-1337/reddit-clone.git/internal/controller/http/v1"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
	"github.com/asam-1337/reddit-clone.git/internal/service"
	"github.com/asam-1337/reddit-clone.git/pkg/httpserver"
	"log"
	"sync"
)

func main() {
	mu := &sync.Mutex{}
	repo := repository.NewRepository(mu)
	services := service.NewService(repo)
	handlers := v1.NewHandler(services)

	srv := new(httpserver.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured in starting server: %s", err.Error())
	}

}

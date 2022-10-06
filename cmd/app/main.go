package main

import (
	"github.com/asam-1337/reddit-clone.git/internal/controller/http/v1"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
	"github.com/asam-1337/reddit-clone.git/internal/service"
	"github.com/asam-1337/reddit-clone.git/pkg/httpserver"
	"github.com/spf13/viper"
	"log"
	"sync"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured in initializing config: %s", err.Error())
	}

	mu := &sync.Mutex{}
	repo := repository.NewRepository(mu)
	services := service.NewService(repo)
	handlers := v1.NewHandler(services)

	srv := new(httpserver.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured in starting server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

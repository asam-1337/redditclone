package main

import (
	"github.com/asam-1337/reddit-clone.git/internal/controller/http/v1"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
	"github.com/asam-1337/reddit-clone.git/internal/service"
	"github.com/asam-1337/reddit-clone.git/pkg/httpserver"
<<<<<<< HEAD
=======
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
>>>>>>> 9095c13b73e4758d1574c5bcebdee4ce847e73b0
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

func main() {
	if err := initConfig(); err != nil {
<<<<<<< HEAD
		log.Fatalf("error occured in initializing config: %s", err.Error())
=======
		log.Fatalf("error occured in read config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured in read .env file: %s", err.Error())
>>>>>>> 9095c13b73e4758d1574c5bcebdee4ce847e73b0
	}

	mu := &sync.Mutex{}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error occured in db open: %s", err.Error())
	}

	repo := repository.NewRepository(mu, db)
	services := service.NewService(repo)
	handlers := v1.NewHandler(services)

	srv := new(httpserver.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured in starting server: %s", err.Error())
	}

}

func initConfig() error {
<<<<<<< HEAD
	viper.AddConfigPath("configs")
=======
	viper.AddConfigPath("./config")
>>>>>>> 9095c13b73e4758d1574c5bcebdee4ce847e73b0
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

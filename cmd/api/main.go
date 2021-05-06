package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/innovember/real-time-forum/internal/mwares"
	userDelivery "github.com/innovember/real-time-forum/internal/user/delivery"
	userRepo "github.com/innovember/real-time-forum/internal/user/repository"
	userUsecase "github.com/innovember/real-time-forum/internal/user/usecases"

	"github.com/innovember/real-time-forum/config"
	"github.com/innovember/real-time-forum/pkg/database"
)

func main() {
	log.Println("Server is starting...")

	config, err := config.LoadConfig("./config/config.json")
	if err != nil {
		log.Fatalln("config error: ", err)
	}

	if !database.FileExist(filepath.Join(config.GetDBPath(), config.GetDBFilename())) {
		err := database.CreateDir(config.GetDBPath())
		if err != nil {
			log.Fatal("dbDir err: ", err)
		}
	}
	dbConn, err := database.GetDBInstance(config.GetDBDriver(), config.GetProdDBConnString())
	if err != nil {
		log.Fatal("dbConn err: ", err)
	}
	defer dbConn.Close()
	if err := database.UploadSchemesToDB(dbConn, config.GetDBSchemesDir()); err != nil {
		log.Fatal("upload schemes err: ", err)
	}

	userRepository := userRepo.NewUserDBRepository(dbConn)
	userUsecase := userUsecase.NewUserUsecase(userRepository)

	mux := http.NewServeMux()
	mm := mwares.NewMiddlewareManager()

	userHandler := userDelivery.NewUserHandler(userUsecase)
	userHandler.Configure(mux, mm)
	log.Println("Server is listening", config.GetLocalServerPath())
	err = http.ListenAndServe(config.GetPort(), mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

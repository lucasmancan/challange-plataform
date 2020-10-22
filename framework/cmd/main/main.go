package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasmancan/challange-plataform/application/repositories"
	"github.com/lucasmancan/challange-plataform/framework/cmd/server"
	"github.com/lucasmancan/challange-plataform/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	myRouter := mux.NewRouter().StrictSlash(true)

	userRepository := repositories.UserRepositoryDb{Db: db}

	userController := server.UserController{UserRepositoy: &userRepository}
	transferController := server.TransferController{TransferRepository: &userRepository}

	myRouter.HandleFunc("/users/{token}", userController.FindByToken).Methods("GET")
	myRouter.HandleFunc("/users", userController.Save).Methods("POST")

	myRouter.HandleFunc("/transfer/{id}", transferController.FindById).Methods("GET")
	myRouter.HandleFunc("/transfer", transferController.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

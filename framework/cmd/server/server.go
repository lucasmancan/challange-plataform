package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasmancan/challange-plataform/application/repositories"
	"github.com/lucasmancan/challange-plataform/domain"
	"github.com/lucasmancan/challange-plataform/framework/utils"
)

type UserController interface {
	Save(reponse http.ResponseWriter, request *http.Request)
	FindByToken(reponse http.ResponseWriter, request *http.Request)
}

type UserControllerImpl struct {
	repo *repositories.UserRepositoryDb
}

func (controller UserControllerImpl) save(reponse http.ResponseWriter, request *http.Request) {

	var user *domain.User
	result, err := controller.repo.Insert(user)

	if err != nil {
		fmt.Fprintf(reponse, "Error")
	}

	fmt.Fprintf(reponse, "Update User Endpoint Hit: "+result.Email)

}

func (controller UserControllerImpl) findByToken(reponse http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	result, err := controller.repo.FindByToken(vars["token"])

	if err != nil {
		fmt.Fprintf(reponse, "Error")
	}

	fmt.Fprintf(reponse, "Update User Endpoint Hit: "+result.Email)

}

func handleRequests(repo *repositories.UserRepositoryDb) {
	myRouter := mux.NewRouter().StrictSlash(true)

	controller := UserControllerImpl{repo: repo}

	myRouter.HandleFunc("/users/{token}", controller.findByToken).Methods("GET")
	myRouter.HandleFunc("/users", controller.save).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	db := utils.ConnectDB()

	userRepo := repositories.UserRepositoryDb{Db: db}
	handleRequests(&userRepo)
}

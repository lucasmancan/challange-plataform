package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasmancan/challange-plataform/application/repositories"
	"github.com/lucasmancan/challange-plataform/domain"
)

type TransferControllerInterface interface {
	Save(reponse http.ResponseWriter, request *http.Request)
	FindByToken(reponse http.ResponseWriter, request *http.Request)
}

type TransferController struct {
	TransferRepository *repositories.UserRepositoryDb
}

func (controller TransferController) Create(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(domain.Response{Message: "Error decoding request body" + err.Error(), Status: 400})
		return
	}

	result, err := controller.TransferRepository.Insert(&user)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(domain.Response{Message: err.Error(), Status: 400})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (controller TransferController) FindById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	result, err := controller.TransferRepository.FindByToken(vars["token"])

	if err != nil {
		json.NewEncoder(w).Encode(domain.Response{Message: "Error" + err.Error(), Status: 400})
	}

	json.NewEncoder(w).Encode(result)
}

package main

import (
	"log"

	"github.com/lucasmancan/challange-plataform/application/repositories"
	"github.com/lucasmancan/challange-plataform/domain"
	"github.com/lucasmancan/challange-plataform/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Lucas",
		Email:    "lucasfmancan@gmail.com",
		Password: "123456",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}

	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Print(result)

}

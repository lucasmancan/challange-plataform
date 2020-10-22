package repositories

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/lucasmancan/challange-plataform/domain"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	FindByToken(token string) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	log.Printf(" Error trying to insert %v", user)

	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during user validation %v", err)
	}
	return user, nil

	// err = repo.Db.Create(user).Error

	// if err != nil {
	// 	log.Fatalf("Error to persist user: %v", err)
	// 	return user, err
	// }

	// return user, nil
}

func (repo UserRepositoryDb) FindByToken(token string) (*domain.User, error) {

	var user *domain.User = &domain.User{Name: "Lucas"}

	log.Print(token)

	// repo.Db.Where("token = ?", token).First(&user)

	return user, nil
}

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

	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during user validation %v", err)
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persist user: %v", err)
		return user, err
	}

	return user, nil
}

func (repo UserRepositoryDb) FindByToken(token string) (*domain.User, error) {

	var user *domain.User

	log.Print(token)
	repo.Db.Where("token = ?", token).First(&user)

	return user, nil
}

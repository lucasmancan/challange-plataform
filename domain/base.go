package domain

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string     `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Timer `json:"createdAt" gorm:"type:datetime"`
	UpdatedAt time.Timer `json:"updatedAt" gorm:"type:datetime"`
	DeletedAt time.Timer `json:"deletedAt" gorm:"type:datetime" sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OauthUser struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey;type:char(36)"`
	FirstName *string        `json:"firstName"`
	LastName  *string        `json:"lastName"`
	Email     string         `json:"email" gorm:"unique;not null;"`
	Language  string         `json:"language" gorm:"default:EN;not null;type:varchar(4)"`
	Password  string         `json:"password" gorm:"not null;"`
	RoleID    uuid.UUID      `json:"-"`
	Role      AuthRole       `json:"role" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (user *OauthUser) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	user.Id = uuid
	return nil
}

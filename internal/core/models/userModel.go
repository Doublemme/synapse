package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OauthUser struct {
	Id        uuid.UUID      `gorm:"primaryKey;type:char(36)" json:"id"`
	FirstName *string        `json:"firstName"`
	LastName  *string        `json:"lastName"`
	Email     string         `gorm:"unique;not null;" json:"email"`
	Language  string         `gorm:"default:EN;not null;type:varchar(4)" json:"language"`
	Password  string         `gorm:"not null;" json:"password"`
	RoleID    uuid.UUID      `json:"-"`
	Role      AuthRole       `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL" json:"role"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (user *OauthUser) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	user.Id = uuid
	return nil
}

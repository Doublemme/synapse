package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OauthUser struct {
	Id        uuid.UUID `gorm:"primaryKey;type:char(36)"`
	FirstName *string
	LastName  *string
	Email     string `gorm:"unique;not null;"`
	Language  string `gorm:"default:EN;not null;type:varchar(4)"`
	Password  string `gorm:"not null;"`
	RoleID    uuid.UUID
	Role      AuthRole `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *OauthUser) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	user.Id = uuid
	return nil
}

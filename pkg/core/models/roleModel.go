package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRole struct {
	Id          string `gorm:"primaryKey;type:char(36)"`
	Name        string
	Description string       `gorm:"type:text"`
	Actions     []AuthAction `gorm:"many2many:role_actions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (role *AuthRole) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	role.Id = uuid.String()

	return nil
}

type AuthModule struct {
	Id          string         `gorm:"primaryKey;type:char(36)"`
	Name        string         `gorm:"unique"`
	Description string         `gorm:"type:text"`
	Resources   []AuthResource `gorm:"foreignKey:ModuleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (module *AuthModule) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	module.Id = uuid.String()

	return nil
}

type AuthResource struct {
	Id          string `gorm:"primaryKey;type:char(36)"`
	ModuleId    string
	Name        string       `gorm:"unique"`
	Description string       `gorm:"type:text"`
	Actions     []AuthAction `gorm:"foreignKey:ResourceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (resource *AuthResource) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	resource.Id = uuid.String()

	return nil
}

type AuthAction struct {
	Id          string `gorm:"primaryKey;type:char(36)"`
	ResourceId  string
	Name        string `gorm:"unique"`
	Description string `gorm:"type:text"`
}

func (action *AuthAction) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	action.Id = uuid.String()

	return nil
}

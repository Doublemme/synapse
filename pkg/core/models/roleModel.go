package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRole struct {
	Id          uuid.UUID `gorm:"primaryKey;type:char(36);"`
	Name        string
	Description string       `gorm:"type:text"`
	Actions     []AuthAction `gorm:"many2many:role_actions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *gorm.DeletedAt `gorm:"index"`
}

func (role *AuthRole) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	role.Id = uuid

	return nil
}

type AuthModule struct {
	Id          uuid.UUID      `gorm:"primaryKey;type:char(36);"`
	Name        string         `gorm:"unique"`
	Description string         `gorm:"type:text"`
	Resources   []AuthResource `gorm:"foreignKey:ModuleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (module *AuthModule) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	module.Id = uuid

	return nil
}

type AuthResource struct {
	Id          uuid.UUID `gorm:"primaryKey;type:char(36);"`
	ModuleId    uuid.UUID
	Name        string       `gorm:"unique"`
	Description string       `gorm:"type:text"`
	Actions     []AuthAction `gorm:"foreignKey:ResourceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (resource *AuthResource) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	resource.Id = uuid

	return nil
}

type AuthAction struct {
	Id          uuid.UUID `gorm:"primaryKey;type:char(36);"`
	ResourceId  uuid.UUID
	Name        string
	Description string `gorm:"type:text"`
}

func (action *AuthAction) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	action.Id = uuid

	return nil
}

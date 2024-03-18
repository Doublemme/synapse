package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRole struct {
	Id          uuid.UUID       `json:"id" gorm:"primaryKey;type:char(36);"`
	Name        string          `json:"name"`
	Description string          `json:"description" gorm:"type:text"`
	Actions     []*AuthAction   `json:"actions" gorm:"many2many:role_actions;"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
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
	Id          uuid.UUID `json:"id" gorm:"primaryKey;type:char(36);"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description" gorm:"type:text"`
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
	Id          uuid.UUID  `json:"id" gorm:"primaryKey;type:char(36);"`
	ModuleId    uuid.UUID  `json:"moduleId"`
	Module      AuthModule `json:"-" gorm:"foreignKey:ModuleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string     `json:"name" gorm:"unique"`
	Description string     `json:"description" gorm:"type:text"`
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
	Id          uuid.UUID    `json:"id" gorm:"primaryKey;type:char(36);"`
	ResourceId  uuid.UUID    `json:"resourceId"`
	Resource    AuthResource `json:"-" gorm:"foreignKey:ResourceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string       `json:"name"`
	Description string       `json:"description" gorm:"type:text"`
	Roles       []*AuthRole  `json:"-" gorm:"many2many:role_actions;"`
}

func (action *AuthAction) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	action.Id = uuid

	return nil
}

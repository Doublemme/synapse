package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRole struct {
	Id          uuid.UUID       `gorm:"primaryKey;type:char(36);" json:"id"`
	Name        string          `json:"name"`
	Description string          `gorm:"type:text" json:"description"`
	Actions     []*AuthAction   `gorm:"many2many:role_actions;" json:"actions"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deletedAt"`
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
	Id          uuid.UUID `gorm:"primaryKey;type:char(36);" json:"id"`
	Name        string    `gorm:"unique" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
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
	Id          uuid.UUID  `gorm:"primaryKey;type:char(36);" json:"id"`
	ModuleId    uuid.UUID  `json:"moduleId"`
	Module      AuthModule `json:"-" gorm:"foreignKey:ModuleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string     `gorm:"unique" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
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
	Id          uuid.UUID    `gorm:"primaryKey;type:char(36);" json:"id"`
	ResourceId  uuid.UUID    `json:"resourceId"`
	Resource    AuthResource `json:"-" gorm:"foreignKey:ResourceId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string       `json:"name"`
	Description string       `gorm:"type:text" json:"description"`
	Roles       []*AuthRole  `gorm:"many2many:role_actions;" json:"-"`
}

func (action *AuthAction) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV7()
	if err != nil {
		return err
	}

	action.Id = uuid

	return nil
}

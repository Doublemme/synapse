package helpers

import (
	"encoding/json"
	"os"

	"github.com/doublemme/synapse/pkg/core/models"
	"github.com/doublemme/synapse/pkg/synapse/types"
	"gorm.io/gorm"
)

// Load the acl configuration from the given files
func UnmarshalAcl(acl *[]types.AclModule, configFiles ...string) error {

	for _, aclConf := range configFiles {
		moduleAcl := make([]types.AclModule, 0)

		file, err := os.OpenFile(aclConf, 0, os.ModeAppend)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&moduleAcl)
		if err != nil {
			return err
		}

		*acl = append(*acl, moduleAcl...)
	}

	return nil
}

func SyncAclToDb(db *gorm.DB, acls *[]types.AclModule) error {
	// Loop all the acl modules given
	for _, acl := range *acls {
		module := models.AuthModule{
			Name:        acl.Name,
			Description: acl.Description,
		}

		// loop throught the resources
		for _, res := range acl.Resources {
			resource := models.AuthResource{
				ModuleId:    module.Id,
				Name:        res.Name,
				Description: res.Description,
			}
			// loop throught the resource.actions and create them
			for _, act := range res.Actions {
				action := models.AuthAction{
					Name:        act.Name,
					Description: act.Description,
				}
				resource.Actions = append(resource.Actions, action)
			}
			// Append the resource to the module resources list
			module.Resources = append(module.Resources, resource)
		}

		// Create the module and all the resources and actions
		if err := db.Model(&models.AuthModule{}).Create(&module).Error; err != nil {
			return err
		}

	}

	return nil
}

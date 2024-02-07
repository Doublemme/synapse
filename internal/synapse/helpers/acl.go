package helpers

import (
	"encoding/json"
	"io"

	"github.com/doublemme/synapse/internal/core/models"
	"github.com/doublemme/synapse/internal/synapse/types"
	"gorm.io/gorm"
)

// Load the modules for the access control list from the reader into the &acl
func UnmarshalAcl(acl *[]types.AclModule, reader io.Reader) error {

	moduleAcl := make([]types.AclModule, 0)

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&moduleAcl)
	if err != nil {
		return err
	}

	*acl = append(*acl, moduleAcl...)

	return nil
}

// Synchronize the ACL modules provided to the database
//
// return an error if something in the insert operation goes wrong
//
// @param db *gorm.DB - A pointer to the gorm DB instance
//
// @param acls *[]AclModules - A pointer to a slice of Acl modules
func SyncAcl(db *gorm.DB, modules *[]types.AclModule) error {

	// Loop all the modules and insert it if it's not found in db
	for _, m := range *modules {
		module, exist := getModuleByName(m.Name, db)
		if !exist {

			module = models.AuthModule{

				Name:        m.Name,
				Description: m.Description,
			}
			if err := db.Create(&module).Error; err != nil {
				return err
			}
		}

		// loop throught the resources
		for _, res := range m.Resources {

			resource, exist := getResourceByName(module.Id.String(), res.Name, db)
			if !exist {
				resource = models.AuthResource{
					ModuleId:    module.Id,
					Name:        res.Name,
					Description: res.Description,
				}

				if err := db.Create(&resource).Error; err != nil {
					return err
				}
			}

			// loop throught the resource actions
			for _, act := range res.Actions {

				action, exist := getActionByName(resource.Id.String(), act.Name, db)
				if !exist {

					action = models.AuthAction{
						ResourceId:  resource.Id,
						Name:        act.Name,
						Description: act.Description,
					}

					if err := db.Create(&action).Error; err != nil {
						return err
					}
				}
			}

		}

	}
	//TODO: Implement deletions of removed modules, resources or actions
	return nil
}

// Search for the module with the given name and return it if exist
func getModuleByName(moduleName string, db *gorm.DB) (models.AuthModule, bool) {
	exist := false
	var model models.AuthModule

	res := db.Model(&models.AuthModule{}).Where("name = ?", moduleName).Take(&model)
	if res.RowsAffected == 1 {
		exist = true
	}

	return model, exist
}

// Search for a resource filtered by moduleId and the resource name and return it if exists
func getResourceByName(moduleId string, resName string, db *gorm.DB) (models.AuthResource, bool) {
	exist := false
	var resource models.AuthResource

	res := db.Model(&models.AuthResource{}).Where("name = ? AND module_id = ?", resName, moduleId).Take(&resource)
	if res.RowsAffected == 1 {
		exist = true
	}

	return resource, exist
}

// Search for an resource action filtered by resource id and action name and return it if exists
func getActionByName(resourceId string, actionName string, db *gorm.DB) (models.AuthAction, bool) {
	exist := false
	var action models.AuthAction

	res := db.Model(&models.AuthAction{}).Where("name = ? AND resource_id = ?", actionName, resourceId).Take(&action)
	if res.RowsAffected == 1 {
		exist = true
	}

	return action, exist
}

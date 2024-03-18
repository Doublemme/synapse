package helpers

import (
	"bytes"
	"testing"

	"github.com/doublemme/synapse/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalAcl(t *testing.T) {
	expectedRes := []types.AclModule{{Name: "core", Description: "Core module resources", Resources: []types.AclResource{{Name: "general", Description: "Grant the access to the application general settings", Actions: []types.AclAction{{Name: "read", Description: "Allow the user to read the application's general settings"}, {Name: "write", Description: "Allow the user to edit the application's general settings"}}}}}}
	var aclModules []types.AclModule

	var buffer bytes.Buffer
	_, err := buffer.Write([]byte(`[{"name": "core","description": "Core module resources","resources": [{"name": "general","description": "Grant the access to the application general settings","actions":[{"name": "read","description": "Allow the user to read the application's general settings"},{"name": "write","description": "Allow the user to edit the application's general settings"}]}]}]`))
	assert.ErrorIs(t, err, nil, "error writing string to the buffer")

	assert.Nil(t, UnmarshalAcl(&aclModules, &buffer), "error unmarshaling acl modules")
	assert.Equal(t, expectedRes, aclModules)
}

package rbacframe_test

import (
    "encoding/json"
    "github.com/go-chassis/sc-client/rbacframe"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestRoleStruct(t *testing.T)  {
    perms := []byte(`{"resources": ["service", "instance"], "verbs": ["get", "list", "create"]}`)
    var p rbacframe.Permission
    t.Run("perm struct test", func(t *testing.T) {
        err := json.Unmarshal(perms, &p)
        assert.NoError(t, err)
    })

    role := []byte(`{"id": "123456", "name": "tester", "perms": [{"resources": ["service", "instance"], "verbs": ["get", "list", "create"]}]}`)
    var r rbacframe.Role
    t.Run("role struct test", func(t *testing.T) {
        err := json.Unmarshal(role, &r)
        assert.NoError(t, err)
    })

    roles := []byte(`{"data": [{"id": "123456", "name": "tester", "perms": [{"resources": ["service", "instance"], "verbs": ["get", "list", "create"]}]}]}`)
    var rs rbacframe.RoleResponse
    t.Run("roles struct test", func(t *testing.T) {
        err := json.Unmarshal(roles, &rs)
        assert.NoError(t, err)
    })
}

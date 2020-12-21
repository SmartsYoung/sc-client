package rbacframe_test

import (
    "encoding/json"
    "github.com/go-chassis/sc-client/rbacframe"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAccountStruct(t *testing.T) {
    account := []byte(`{"id": "123456", "name": "dev_test", "password": "Complicate_1", "roles": ["admin", "developer"],
                        "tokenExpirationTime":"30s", "currentPassword":"Complicate_1", "status":"UP"}`)
    var res rbacframe.Account
    t.Run("account test", func(t *testing.T) {
        err := json.Unmarshal(account, &res)
        assert.NoError(t, err)
    })

    accounts := []byte(`{"total": 1, "data": [{"id": "123456", "name": "dev_test", "password": "Complicate_1", "roles": ["admin", "developer"], 
                        "tokenExpirationTime":"30s", "currentPassword":"Complicate_1", "status":"UP"}]}`)
    var res1 rbacframe.AccountResponse
    t.Run("accounts response test", func(t *testing.T) {
        err := json.Unmarshal(accounts, &res1)
        assert.NoError(t, err)
    })

    token := []byte(`{"token": "eyTest.ptVZ123456"}`)
    var to rbacframe.Token 
    t.Run("token test", func(t *testing.T) {
        err := json.Unmarshal(token, &to)
        assert.NoError(t, err)
    })
}


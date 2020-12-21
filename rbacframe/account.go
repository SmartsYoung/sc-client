package rbacframe

type AccountResponse struct {
	Total    int64      `json:"total"`
	Accounts []*Account `json:"data,omitempty"`
}

type Account struct {
	ID                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	Password            string   `json:"password,omitempty"`
	Roles               []string `json:"roles,omitempty"`
	TokenExpirationTime string   `json:"tokenExpirationTime,omitempty"`
	CurrentPassword     string   `json:"currentPassword,omitempty"`
	Status              string   `json:"status,omitempty"`
}

type Token struct {
	TokenStr string `json:"token,omitempty"`
}

package rbacframe

type RoleResponse struct {
	Roles []*Role `json:"data,omitempty"`
}

type Role struct {
	ID    string        `json:"id,omitempty"`
	Name  string        `json:"name,omitempty"`
	Perms []*Permission `json:"perms,omitempty"`
}

type Permission struct {
	Resources []string `json:"resources,omitempty"`
	Verbs     []string `json:"verbs,omitempty"`
}

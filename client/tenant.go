package client

import "encoding/json"

type TenantResource struct {
	Client *Client
}

type Tenant struct {
	ID     string `json:"id"`
	Tenant string `json:"tenant"`
}

// Create tenant
func (r *TenantResource) Create(s *Tenant) (*Tenant, error) {
	c := *r.Client
	j, err := c.Create("/tenants", s)
	if err != nil {
		return nil, err
	}

	tenant := &Tenant{}
	if err := json.Unmarshal(j, tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

// Get tenant
func (r *TenantResource) Get(name string) (*Tenant, error) {
	c := *r.Client
	j, err := c.Get("/tenants", name)
	if err != nil {
		return nil, err
	}

	tenant := &Tenant{}
	if err := json.Unmarshal(j, tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

// All tenants
func (r *TenantResource) All() (*[]Tenant, error) {
	c := *r.Client
	j, err := c.All("/tenants")
	if err != nil {
		return nil, err
	}

	tenants := &[]Tenant{}
	if err := json.Unmarshal(j, tenants); err != nil {
		return nil, err
	}

	return tenants, nil
}

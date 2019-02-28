package roles

import "github.com/gophercloud/gophercloud"

const (
	ExtPath  = "OS-KSADM"
	RolePath = "roles"
	UserPath = "users"
)

// resourceURL generates URL for listing resources.
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(ExtPath, RolePath, id)
}

// rootURL generates URL for listing all available global roles.
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(ExtPath, RolePath)
}

// userRoleURL generates URL for getting role of user.
func userRoleURL(c *gophercloud.ServiceClient, tenantID, userID, roleID string) string {
	return c.ServiceURL("tenants", tenantID, UserPath, userID, RolePath, ExtPath, roleID)
}

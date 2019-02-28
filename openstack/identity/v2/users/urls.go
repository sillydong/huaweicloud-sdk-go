package users

import "github.com/gophercloud/gophercloud"

const (
	tenantPath = "tenants"
	userPath   = "users"
	rolePath   = "roles"
)

// ResourceURL generates URL for listing resource of a single user
func ResourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(userPath, id)
}

// rootURL generates URL for listing exist users
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(userPath)
}

// listRolesURL generates URL for list roles of user
func listRolesURL(c *gophercloud.ServiceClient, tenantID, userID string) string {
	return c.ServiceURL(tenantPath, tenantID, userPath, userID, rolePath)
}

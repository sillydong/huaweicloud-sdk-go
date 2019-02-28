package tenants

import "github.com/gophercloud/gophercloud"

// listURL generates URL for listing tenants.
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}

// getURL generates URL for getting tenant.
func getURL(client *gophercloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

// createURL generates URL for createing tenant.
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}

// deleteURL generates URL for deleteing tenant.
func deleteURL(client *gophercloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

// updateURL generates URL for updating a tenant.
func updateURL(client *gophercloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

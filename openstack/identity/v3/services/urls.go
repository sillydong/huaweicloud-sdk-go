package services

import "github.com/gophercloud/gophercloud"

// listURL generate url to querying services
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("services")
}

// createURL generate url to create service details
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("services")
}

// serviceURL generate url to show service details
func serviceURL(client *gophercloud.ServiceClient, serviceID string) string {
	return client.ServiceURL("services", serviceID)
}

// updateURL generate url to update service details
func updateURL(client *gophercloud.ServiceClient, serviceID string) string {
	return client.ServiceURL("services", serviceID)
}

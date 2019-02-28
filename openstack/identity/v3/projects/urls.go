package projects

import "github.com/gophercloud/gophercloud"

// listURL generate url to querying information about a specified project
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("projects")
}

// getURL generate url to show project details
func getURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

// createURL generate url to add project
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("projects")
}

// deleteURL generate url for deleting project
func deleteURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

// updateURL generate url to update project
func updateURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

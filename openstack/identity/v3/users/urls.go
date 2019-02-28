package users

import "github.com/gophercloud/gophercloud"

// listURL generate url to query a user list
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("users")
}

// getURL generate url to query detailed information about a specified user
func getURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

// createURL generate url to create a user under a tenant
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("users")
}

// updateURL generate url to modify user information under a tenant
func updateURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

// deleteURL generate url to delete a specified user
func deleteURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

// updatePasswdURL generate url to change the password for a user
func updatePasswdURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "password")
}

// listGroupsURL generate url to query the information about the user group
func listGroupsURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "groups")
}

// listProjectsURL generate url to query the project information
func listProjectsURL(client *gophercloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "projects")
}

// listInGroupURL generate url to list users in group
func listInGroupURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID, "users")
}

// operateOnGroupUserURL generate url to modify group users
func operateOnGroupUserURL(client *gophercloud.ServiceClient,
	groupID string, userID string) string {

	return client.ServiceURL("groups", groupID, "users", userID)
}

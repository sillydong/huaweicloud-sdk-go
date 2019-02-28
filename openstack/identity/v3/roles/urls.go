package roles

import "github.com/gophercloud/gophercloud"

const (
	rolePath = "roles"
)

// listURL generate url to list roles
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

// getURL generate url to show role details
func getURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

// createURL generate url for creating role
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

// updateURL generate url for updating role
func updateURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

// deleteURL generate url for deleting role
func deleteURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

// listAssignmentsURL generate url for listing role assignments
func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

// assignURL generate url to assign or unassign role to group on domain or
// project
func assignURL(client *gophercloud.ServiceClient, targetType, targetID, actorType, actorID, roleID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath, roleID)
}

// checkRoleOfURL generate url to check whether group has role assignment on
// domain or project
func checkRoleOfURL(client *gophercloud.ServiceClient,
	targetType, targetID, actorType, actorID, roleID string) string {

	return client.ServiceURL(
		targetType, targetID,
		actorType, actorID,
		"roles", roleID,
	)
}

// listRolesOfURL generate url to list rolesto group on domain or project
func listRolesOfURL(client *gophercloud.ServiceClient,
	targetType, targetID, actorType, actorID string) string {

	return client.ServiceURL(
		targetType, targetID,
		actorType, actorID,
		"roles",
	)
}

package attachinterfaces

import "github.com/gophercloud/gophercloud"

// listInterfaceURL generate URL to list port interfaces
func listInterfaceURL(client *gophercloud.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}

// getInterfaceURL generate URL to show port interface details
func getInterfaceURL(client *gophercloud.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}

// createInterfaceURL generate URL to attach interface
func createInterfaceURL(client *gophercloud.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}

// deleteInterfaceURL generate URL to detach interface
func deleteInterfaceURL(client *gophercloud.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}

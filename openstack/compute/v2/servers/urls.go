package servers

import "github.com/gophercloud/gophercloud"

// createURL generate url to create server
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("servers")
}

// listURL generate url for querying ECSes
func listURL(client *gophercloud.ServiceClient) string {
	return createURL(client)
}

// listDetailURL generate url to list servers details
func listDetailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("servers", "detail")
}

// deleteURL generate url to delete server
func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

// getURL generate url to get one server detail
func getURL(client *gophercloud.ServiceClient, id string) string {
	return deleteURL(client, id)
}

// updateURL generate url for update server
func updateURL(client *gophercloud.ServiceClient, id string) string {
	return deleteURL(client, id)
}

// actionURL generate action URL for server starting, rebooting and stopping
func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

// metadatumURL generate URL to get or set metadata items
func metadatumURL(client *gophercloud.ServiceClient, id, key string) string {
	return client.ServiceURL("servers", id, "metadata", key)
}

// metadataURL generate URL to list or update metadatas
func metadataURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "metadata")
}

// listAddressesURL generate url for server ips
func listAddressesURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "ips")
}

// listAddressesByNetworkURL generate url for listing server ips by network
func listAddressesByNetworkURL(client *gophercloud.ServiceClient, id, network string) string {
	return client.ServiceURL("servers", id, "ips", network)
}

// passwordURL generate for updating server password
func passwordURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "os-server-password")
}

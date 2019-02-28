package endpoints

import "github.com/gophercloud/gophercloud"

// listURL generate url to querying endpoints
func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("endpoints")
}

// endpointURL generate url to show endpoint details
func endpointURL(client *gophercloud.ServiceClient, endpointID string) string {
	return client.ServiceURL("endpoints", endpointID)
}

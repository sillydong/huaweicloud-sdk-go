package availabilityzones

import "github.com/gophercloud/gophercloud"

// listURL generate url to list availability zones
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}

// listDetailURL generate url to list availablity zone detail
func listDetailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone", "detail")
}

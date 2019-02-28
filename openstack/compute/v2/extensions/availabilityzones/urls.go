package availabilityzones

import "github.com/gophercloud/gophercloud"

// listURL generate url to list availability zones
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}

// func listDetailURL(c *golangsdk.ServiceClient) string {
//	return c.ServiceURL("os-availability-zone", "detail")
// }

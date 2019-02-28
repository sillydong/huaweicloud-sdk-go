package availabilityzones

import "github.com/gophercloud/gophercloud"

// listURL generates URL for list avaliabilityzones
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}

package quotasets

import "github.com/gophercloud/gophercloud"

const resourcePath = "os-quota-sets"

// getURL generates URL for getting  quotasets
func getURL(c *golangsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}

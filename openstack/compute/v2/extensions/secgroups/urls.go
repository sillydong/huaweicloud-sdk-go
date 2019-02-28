package secgroups

import "github.com/gophercloud/gophercloud"

const (
	secgrouppath = "os-security-groups"
	rulepath     = "os-security-group-rules"
)

// resourceURL generate a url for specified secgrouppath
func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(secgrouppath, id)
}

// rootURL generate a url for secgrouppath
func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(secgrouppath)
}

// listByServerURL generate a url for server's secgrouppath
func listByServerURL(c *gophercloud.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, secgrouppath)
}

// rootRuleURL generate a url for root rulepath
func rootRuleURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rulepath)
}

// resourceRuleURL generate a url for resource rulepath
func resourceRuleURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rulepath, id)
}

// serverActionURL generate a url to do some actions of server
func serverActionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("servers", id, "action")
}

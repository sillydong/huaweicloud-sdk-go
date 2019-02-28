package members

import (
	"github.com/gophercloud/gophercloud"
)

func imageMembersURL(c *gophercloud.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID, "members")
}

// listMembersURL generate a URL for listing all members of an image
func listMembersURL(c *gophercloud.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

// createMemberURL generate a URL for adding a member to an image
func createMemberURL(c *gophercloud.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

// imageMemberURL generate a URL for  getting a image member
func imageMemberURL(c *gophercloud.ServiceClient, imageID string, memberID string) string {
	return c.ServiceURL("images", imageID, "members", memberID)
}

// getMemberURL generate a URL to get the detail of an image member
func getMemberURL(c *gophercloud.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

// updateMemberURL generate a URL to update the status of an image member
func updateMemberURL(c *gophercloud.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

// deleteMemberURL generate a URL to remove an image member
func deleteMemberURL(c *gophercloud.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

// getMemberSchemas generate a url to get member schemas
func getMemberSchemas(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("schemas", "member")
}

// getMembersSchemas generate a url to get member schemas
func getMembersSchemas(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("schemas", "members")
}

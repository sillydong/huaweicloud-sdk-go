package bandwidths

import "github.com/huaweicloud/huaweicloud-sdk-go"

func UpdateURL(c *gophercloud.ServiceClient, ID string) string {
	return c.ServiceURL("bandwidths", ID)
}

package apiversions

import (
	"strings"

	"github.com/huaweicloud/huaweicloud-sdk-go"
)

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}

func apiInfoURL(c *gophercloud.ServiceClient, version string) string {
	return c.Endpoint + strings.TrimRight(version, "/") + "/"
}

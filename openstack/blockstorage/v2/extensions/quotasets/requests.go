package quotasets

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
)

// GetUsage returns detailed public data about a previously created QuotaSet.
func GetUsage(client *gophercloud.ServiceClient, projectID string) (r GetUsageResult) {
	u := fmt.Sprintf("%s?usage=true", getURL(client, projectID))
	_, r.Err = client.Get(u, &r.Body, nil)
	return
}

package availablezones

import (
	"github.com/gophercloud/gophercloud"
)

// Get available zones
func Get(client *golangsdk.ServiceClient) (r GetResult) {
	_, r.Err = client.Get(getURL(client), &r.Body, nil)
	return
}

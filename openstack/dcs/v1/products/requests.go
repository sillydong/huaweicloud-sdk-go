package products

import (
	"github.com/gophercloud/gophercloud"
)

// Get products
func Get(client *golangsdk.ServiceClient) (r GetResult) {
	_, r.Err = client.Get(getURL(client), &r.Body, nil)
	return
}

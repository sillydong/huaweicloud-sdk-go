package extensions

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(
		c,
		ListExtensionURL(c),
		func(r pagination.PageResult) pagination.Page {
			return ExtensionPage{pagination.SinglePageBase(r)}
		},
	)
}

package testing

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/vpc/v1/security/groups"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
	"github.com/huaweicloud/huaweicloud-sdk-go/testhelper/client"
)

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1/" + "57e98940a77f4bb988a21a7d0603a626/"
	return sc
}

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)
	sc := ServiceClient()

	actual, err := groups.List(sc, groups.ListOpts{}).AllPages()
	resp, err := groups.ExtractGroups(actual)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, &resp)
}
func TestListEachPage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)
	sc := ServiceClient()
	count := 0
	err := groups.List(sc, groups.ListOpts{Limit: 1}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		data, paErr := groups.ExtractGroups(page)

		if paErr != nil {
			t.Errorf("Failed to extract groups: %v", paErr)
			return false, paErr
		}
		th.CheckDeepEquals(t, &ListResponse, &data)

		return true, nil

	})
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
	th.AssertNoErr(t, err)
}

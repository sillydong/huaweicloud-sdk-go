package testing

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/vpc/v2.0/publicips"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
	"github.com/huaweicloud/huaweicloud-sdk-go/testhelper/client"
)

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/" + "128a7bf965154373a7b73c89eb6b65aa/"
	return sc
}
func TestCreateOndemand(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleOndemandSuccessfully(t)
	sc := ServiceClient()

	actual, err := publicips.Create(sc, OndmandOpts)
	th.AssertNoErr(t, err)

	if on, ok := actual.(publicips.PostPaid); ok {
		th.CheckDeepEquals(t, on.Type, OndmandOpts.PublicIP.Type)
		th.CheckDeepEquals(t, on.BandwidthSize, OndmandOpts.Bandwidth.Size)
		th.CheckDeepEquals(t, on.IPVersion, OndmandOpts.PublicIP.IPVersion)

	}
}

func TestCreateWithBSSInfo(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleWithBSSInfoSuccessfully(t)
	sc := ServiceClient()

	actual, err := publicips.Create(sc, BSSOpts)
	th.AssertNoErr(t, err)
	if order, ok := actual.(publicips.PrePaid); ok {
		th.CheckDeepEquals(t, order, CreateResultForBSS)
		th.CheckDeepEquals(t, order, CreateResultForBSS)
	}
}

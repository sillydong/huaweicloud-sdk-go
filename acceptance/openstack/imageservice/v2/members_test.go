// +build acceptance imageservice images

package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/members"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestListImageMembers(t *testing.T) {

	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	image, err := CreateEmptyImage(t, client)
	th.AssertNoErr(t, err)
	defer DeleteImage(t, client, image)

	// FIXME broken, huaweicloud return 400 error
	createResult := members.Create(client, image.ID, "dummymemberid")
	th.AssertNoErr(t, createResult.Err)

	err = members.List(client, image.ID).EachPage(func(page pagination.Page) (bool, error) {
		members, err := members.ExtractMembers(page)
		if err != nil {
			t.Fatalf("Unable to extract members: %v", err)
		}

		for _, member := range members {
			tools.PrintResource(t, member)
		}

		return true, nil
	})
	th.AssertNoErr(t, err)

	detailsResult := members.Get(client, image.ID, "dummymemberid")
	member, err := detailsResult.Extract()
	th.AssertNoErr(t, err)
	tools.PrintResource(t, member)

	deleteResult := members.Delete(client, image.ID, "dummymemberid")
	th.AssertNoErr(t, deleteResult.Err)
}

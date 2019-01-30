// +build acceptance identity

package v2

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/clients"
	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/tools"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/identity/v2/extensions"
)

func TestExtensionsList(t *testing.T) {
	client, err := clients.NewIdentityV2Client()
	if err != nil {
		t.Fatalf("Unable to create an identity client: %v", err)
	}

	allPages, err := extensions.List(client).AllPages()
	if err != nil {
		t.Fatalf("Unable to list extensions: %v", err)
	}

	allExtensions, err := extensions.ExtractExtensions(allPages)
	if err != nil {
		t.Fatalf("Unable to extract extensions: %v", err)
	}

	for _, extension := range allExtensions {
		tools.PrintResource(t, extension)
	}
}

func TestExtensionsGet(t *testing.T) {
	client, err := clients.NewIdentityV2Client()
	if err != nil {
		t.Fatalf("Unable to create an identity client: %v", err)
	}

	extension, err := extensions.Get(client, "OS-KSCRUD").Extract()
	if err != nil {
		t.Fatalf("Unable to get extension OS-KSCRUD: %v", err)
	}

	tools.PrintResource(t, extension)
}

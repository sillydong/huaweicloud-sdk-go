package testing

import (
	"fmt"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/auth/token"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/identity/v2/tokens"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
	"github.com/huaweicloud/huaweicloud-sdk-go/testhelper/client"
)

func tokenPost(t *testing.T, options token.TokenOptions, requestJSON string) tokens.CreateResult {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleTokenPost(t, requestJSON)

	return tokens.Create(client.ServiceClient(), options)
}

func tokenPostErr(t *testing.T, options token.TokenOptions, expectedErr error) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleTokenPost(t, "")

	actualErr := tokens.Create(client.ServiceClient(), options).Err
	th.CheckDeepEquals(t, expectedErr, actualErr)
}

func TestCreateWithPassword(t *testing.T) {
	options := token.TokenOptions{
		Username: "me",
		Password: "swordfish",
	}

	IsSuccessful(t, tokenPost(t, options, `
    {
      "auth": {
        "passwordCredentials": {
          "username": "me",
          "password": "swordfish"
        }
      }
    }
  `))
}

func TestCreateTokenWithTenantID(t *testing.T) {
	options := token.TokenOptions{
		Username: "me",
		Password: "opensesame",
		TenantID: "fc394f2ab2df4114bde39905f800dc57",
	}

	IsSuccessful(t, tokenPost(t, options, `
    {
      "auth": {
        "tenantId": "fc394f2ab2df4114bde39905f800dc57",
        "passwordCredentials": {
          "username": "me",
          "password": "opensesame"
        }
      }
    }
  `))
}

func TestCreateTokenWithTenantName(t *testing.T) {
	options := token.TokenOptions{
		Username:   "me",
		Password:   "opensesame",
		TenantName: "demo",
	}

	IsSuccessful(t, tokenPost(t, options, `
    {
      "auth": {
        "tenantName": "demo",
        "passwordCredentials": {
          "username": "me",
          "password": "opensesame"
        }
      }
    }
  `))
}

func TestRequireUsername(t *testing.T) {
	options := token.TokenOptions{
		Password: "thing",
	}

	message := fmt.Sprintf(gophercloud.CE_MissingInputMessage, "Username")
	err := gophercloud.NewSystemCommonError(gophercloud.CE_MissingInputCode, message)
	tokenPostErr(t, options, err)
}

func tokenGet(t *testing.T, tokenId string) tokens.GetResult {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleTokenGet(t, tokenId)
	return tokens.Get(client.ServiceClient(), tokenId)
}

func TestGetWithToken(t *testing.T) {
	GetIsSuccessful(t, tokenGet(t, "db22caf43c934e6c829087c41ff8d8d6"))
}

package main

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/auth/aksk"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/vpc/v1/vpcs"
)

func main() {
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
		DomainID:         "{domainID}",
	}

	//初始化provider client。
	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}
	//初始化服务 client
	client, err_client := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})
	if err_client != nil {
		fmt.Println("Failed to get the NewVPCV1 client: ", err_client)
		return
	}

	CreateVPC(client)
	GetVPC(client)
	UpdateVPC(client)
	ListVPC(client)
	DeleteVPC(client)

}

func CreateVPC(sc *gophercloud.ServiceClient) {

	resp, err := vpcs.Create(sc, vpcs.CreateOpts{
		Name: "ABC",
		Cidr: "192.168.0.0/16",
	}).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)

}

func UpdateVPC(sc *gophercloud.ServiceClient) {

	resp, err := vpcs.Update(sc, "463497ec-7a31-4c82-91e7-360243e54be0", vpcs.UpdateOpts{
		Name: "ABC-back",
		Cidr: "192.168.0.0/24",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)

}

func GetVPC(sc *gophercloud.ServiceClient) {
	resp, err := vpcs.Get(sc, "463497ec-7a31-4c82-91e7-360243e54be0").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("vpc name is:", resp.Name)
	fmt.Println("vpc Id is:", resp.ID)
	fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
	fmt.Println("vpc Status is:", resp.Status)
	fmt.Println("vpc Cidr is:", resp.Cidr)
	fmt.Println("vpc Routes is:", resp.Routes)

}

func ListVPC(sc *gophercloud.ServiceClient) {

	allpages, err := vpcs.List(sc, vpcs.ListOpts{
		//Limit: 2,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	vpcList, err := vpcs.ExtractVpcs(allpages)
	for _, resp := range vpcList {

		fmt.Println("vpc name is:", resp.Name)
		fmt.Println("vpc Id is:", resp.ID)
		fmt.Println("vpc EnterpriseProjectId is:", resp.EnterpriseProjectId)
		fmt.Println("vpc Status is:", resp.Status)
		fmt.Println("vpc Cidr is:", resp.Cidr)
		fmt.Println("vpc Routes is:", resp.Routes)
	}

}

func DeleteVPC(sc *gophercloud.ServiceClient) {

	resp := vpcs.Delete(sc, "5aaaf1cc-9138-4958-955d-4cd6193ff9ff")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete VPC success!")
}

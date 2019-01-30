package main

import (
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/auth/aksk"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/bss/v1/account"
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

	// 初始化服务的client
	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	//初始化查询参数
	reqTmp := account.ResourceDailyOpts{
		StartTime:           "2018-06-01",
		EndTime:             "2018-06-30",
		PayMethod:           "0",
		CloudServiceType:    "hws.service.type.ebs",
		RegionCode:          "cn-xianhz-1",
		ResourceId:          "",
		EnterpriseProjectId: "",
	}

	//根据查询参数获取消费汇总列表
	rspTmp, err := account.ListResourceDaily(sc, reqTmp)
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//打印返回参数，如currency，totalRecord，totalAmount，dailyRecords等
	fmt.Println("Succeed to get the ResourceDaily List!")
	fmt.Println("totalRecord:", rspTmp.TotalRecord)
	fmt.Println("currency:", rspTmp.Currency)
	fmt.Println("totalAmount:", rspTmp.TotalAmount)
}

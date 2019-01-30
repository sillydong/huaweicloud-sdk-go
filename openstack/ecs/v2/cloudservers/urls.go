package cloudservers

import "github.com/huaweicloud/huaweicloud-sdk-go"

func resetPwdURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("servers", serverID, "os-reset-password")
}

func changeURL(sc *gophercloud.ServiceClient, serverID string) string {
	return sc.ServiceURL("cloudservers", serverID, "changeos")
}

package metricdata

import "github.com/huaweicloud/huaweicloud-sdk-go"

// batch query metric data url
func batchQueryMetricDataURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("batch-query-metric-data")
}

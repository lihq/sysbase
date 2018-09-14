package main

import (
	"fmt"

	"github.com/MonitorMetrics/base/cpu"
	"github.com/MonitorMetrics/base/disk"
	"github.com/MonitorMetrics/base/mem"
	"github.com/MonitorMetrics/base/net"
)

func main() {
	all := []map[string]interface{}{}

	result, err := metricsCPU.Gets()
	if err != nil {
	} else {
		all = append(all, result...)
	}

	result, _ = metricsDisk.Gets()
	if err != nil {
	} else {
		all = append(all, result...)
	}

	result, _ = metricsMem.Gets()
	if err != nil {
	} else {
		all = append(all, result...)
	}

	ipaddrs, _ := metricsNet.GetLANIpAddrs()
	fmt.Println(ipaddrs)
	for _, item := range all {
		fmt.Println(fmt.Sprintf("%s %v %v", item["k"], item["v"], item["t"]))
	}

}

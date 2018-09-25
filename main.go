package main

import (
	"fmt"
	"log"

	"github.com/MonitorMetrics/base/cpu"
	"github.com/MonitorMetrics/base/disk"
	"github.com/MonitorMetrics/base/mem"
	"github.com/MonitorMetrics/base/net"
	"github.com/MonitorMetrics/base/top"

	"github.com/MonitorMetrics/base/models"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	all := []datapoint.DataPoint{}

	result, err := metricsCPU.Gets()
	if err != nil {
	} else {
		all = append(all, result...)
	}

	result, err = metricsDisk.Gets()
	if err != nil {
	} else {
		all = append(all, result...)
	}

	result, err = metricsMem.Gets()
	if err != nil {
	} else {
		all = append(all, result...)
	}

	result, err = metricsTop.Gets()
	if err != nil {
		log.Println("metricsTop.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	ipaddrs, err := metricsNet.GetLANIpAddrs()
	if err != nil {
		log.Println("metricsNet.GetLANIpAddrs failed", err)
	} else {
		fmt.Println(ipaddrs)
		for _, item := range all {
			fmt.Println(item)
		}
	}

}

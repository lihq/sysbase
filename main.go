// Example collect all system base monitor metrics.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/OpenCCTV/sys_base/cpu"
	"github.com/OpenCCTV/sys_base/disk"
	"github.com/OpenCCTV/sys_base/mem"
	"github.com/OpenCCTV/sys_base/net"
	"github.com/OpenCCTV/sys_base/top"

	"github.com/OpenCCTV/sys_base/models"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	all := []datapoint.DataPoint{}

	result, err := metricsCPU.Gets()
	if err != nil {
		log.Fatalln("metricsCPU.Gets", err)
	} else {
		all = append(all, result...)
	}

	result, err = metricsDisk.Gets()
	if err != nil {
		log.Fatalln("metricsCPU.Gets", err)
	} else {
		all = append(all, result...)
	}

	result, err = metricsMem.Gets()
	if err != nil {
		log.Fatalln("metricsMem.Gets", err)
	} else {
		all = append(all, result...)
	}

	result, err = metricsTop.Gets()
	if err != nil {
		log.Println("metricsTop.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	metricsNet.Gets()
	time.Sleep(time.Duration(1) * time.Second)
	result, err = metricsNet.Gets()
	if err != nil {
		log.Println("metricsNet.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	for _, item := range all {
		fmt.Println(item)
	}

	ipaddrs, err := metricsNet.GetLANIpAddrs()
	if err != nil {
		log.Println("metricsNet.GetLANIpAddrs failed", err)
	} else {
		fmt.Println(ipaddrs)
	}

	ipaddrs, err = metricsNet.GetWANIpAddrs()
	if err != nil {
		log.Println("metricsNet.GetWANIpAddrsfailed", err)
	} else {
		fmt.Println(ipaddrs)
	}

}

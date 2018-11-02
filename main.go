// Example collect all system base monitor metrics.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/OpenCCTV/sysbase/cpu"
	"github.com/OpenCCTV/sysbase/disk"
	"github.com/OpenCCTV/sysbase/mem"
	"github.com/OpenCCTV/sysbase/net"
	"github.com/OpenCCTV/sysbase/top"

	"github.com/OpenCCTV/sysbase/models"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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

	ipaddr, err := metricsNet.GetFirstIP()
	if err != nil {
		log.Println("metricsNet.GetFirstIP failed", err)
	} else {
		fmt.Println(ipaddr)
	}

}

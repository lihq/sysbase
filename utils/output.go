package utils

import (
	"log"

	"github.com/OpenCCTV/sysbase/cpu"
	"github.com/OpenCCTV/sysbase/disk"
	"github.com/OpenCCTV/sysbase/mem"
	"github.com/OpenCCTV/sysbase/models"
	"github.com/OpenCCTV/sysbase/net"
	"github.com/OpenCCTV/sysbase/top"
)

func GetMergeAll() []datapoint.DataPoint {
	var result []datapoint.DataPoint
	var err error
	all := []datapoint.DataPoint{}

	result, err = metricsCPU.Gets()
	if err != nil {
		log.Println("metricsCPU.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	result, err = metricsDisk.Gets()
	if err != nil {
		log.Println("metricsDisk.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	result, err = metricsMem.Gets()
	if err != nil {
		log.Println("metricsMem.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	result, err = metricsNet.Gets()
	if err != nil {
		log.Println("metricsNet.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	result, err = metricsTop.Gets()
	if err != nil {
		log.Println("metricsTop.Gets failed", err)
	} else {
		all = append(all, result...)
	}

	return all
}

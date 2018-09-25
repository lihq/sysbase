package metricsCPU

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/MonitorMetrics/base/models"
)

func Gets() (result []datapoint.DataPoint, err error) {
	points := []datapoint.DataPoint{}

	if runtime.GOOS != "linux" {
		err = errors.New(fmt.Sprintf("platform %s not support", runtime.GOOS))
		return points, err
	}

	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return points, err
	}

	fields := strings.Fields(string(b))

	var p datapoint.DataPoint
	now := time.Now()

	if loadMin1, err := strconv.ParseFloat(fields[0], 64); err != nil {
		return points, err
	} else {
		p = datapoint.DataPoint{}
		p.ContentType = datapoint.ContentTypeGauge
		p.Metric = "load.1min"
		p.Value = loadMin1
		p.Timestamp = now
		points = append(points, p)
	}

	if loadMin5, err := strconv.ParseFloat(fields[1], 64); err != nil {
		return points, err
	} else {
		p = datapoint.DataPoint{}
		p.ContentType = datapoint.ContentTypeGauge
		p.Metric = "load.5min"
		p.Value = loadMin5
		p.Timestamp = now
		points = append(points, p)
	}

	if loadMin15, err := strconv.ParseFloat(fields[2], 64); err != nil {
		return points, err
	} else {
		p = datapoint.DataPoint{}
		p.ContentType = datapoint.ContentTypeGauge
		p.Metric = "load.15min"
		p.Value = loadMin15
		p.Timestamp = now
		points = append(points, p)
	}

	return points, nil

}

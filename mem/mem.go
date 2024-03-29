// package metricsMem implements collect memory related monitor metrics.
package metricsMem

import (
	"strconv"
	"strings"

	"github.com/OpenCCTV/sysbase/helpers"
	"github.com/OpenCCTV/sysbase/models"
)

type StatMem struct {
	MemTotal       int64
	MemUsedPercent int64

	SwapTotal       int64
	SwapUsedPercent int64
}

func ParseOutputFree(output string) (StatMem, error) {
	var sm StatMem

	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		// head total used free
		if len(fields) < 4 {
			continue
		}

		section := fields[0]
		totalStr := fields[1]
		usedStr := fields[2]
		//free := fields[3]

		keyPrefix := ""

		if section == "Mem:" {
			keyPrefix = "mem"
		} else if section == "Swap:" {
			keyPrefix = "swap"
		} else {
			continue
		}

		usedPercent := int64(-1)

		total, err := strconv.ParseInt(totalStr, 10, 64)
		if err != nil {
			return sm, err
		}
		used, err := strconv.ParseInt(usedStr, 10, 64)
		if err != nil {
			return sm, err
		}

		if total > 0 {
			usedPercent = int64(float64(used) / float64(total) * 100)
		} else {
			usedPercent = 0
		}

		if keyPrefix == "mem" {
			sm.MemTotal = total
			sm.MemUsedPercent = usedPercent
		} else if keyPrefix == "swap" {
			sm.SwapTotal = total
			sm.SwapUsedPercent = usedPercent
		}
	}

	return sm, nil
}

// Gets returns memory total bytes and used percent from `free -b`.
func Gets() (result []datapoint.DataPoint, err error) {
	points := []datapoint.DataPoint{}
	cmd := `free -b`

	timeoutInSeconds := 1
	out, err := helpers.ExecCommand(cmd, timeoutInSeconds)
	if err != nil {
		return points, err
	}

	var p datapoint.DataPoint
	var _total string

	sm, err := ParseOutputFree(string(out))
	if err != nil {
		return points, err
	}

	p = datapoint.DataPoint{}
	p.Metric = "mem.total"
	p.ContentType = datapoint.ContentTypeGauge
	p.Value = sm.MemTotal
	points = append(points, p)

	p = datapoint.DataPoint{}
	p.Metric = "mem.used.percent"
	p.ContentType = datapoint.ContentTypeGauge
	p.Value = sm.MemUsedPercent
	_total = helpers.Format4HumanSize(float64(sm.MemTotal))
	p.Tags = map[string]interface{}{
		"total": _total,
	}
	points = append(points, p)

	p = datapoint.DataPoint{}
	p.Metric = "swap.total"
	p.ContentType = datapoint.ContentTypeGauge
	p.Value = sm.SwapTotal
	points = append(points, p)

	p = datapoint.DataPoint{}
	p.Metric = "swap.used.percent"
	p.ContentType = datapoint.ContentTypeGauge
	p.Value = sm.SwapUsedPercent
	_total = helpers.Format4HumanSize(float64(sm.SwapTotal))
	p.Tags = map[string]interface{}{
		"total": _total,
	}
	points = append(points, p)

	return points, nil

}

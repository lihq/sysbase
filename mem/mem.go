// package metricsMem implements collect memory related monitor metrics.
package metricsMem

import (
	"strconv"
	"strings"

	"github.com/MonitorMetrics/base/helpers"
	"github.com/MonitorMetrics/base/models"
)

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

	for _, line := range strings.Split(string(out), "\n") {
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
			continue
		}
		used, err := strconv.ParseInt(usedStr, 10, 64)
		if err != nil {
			continue
		}

		if total > 0 {
			usedPercent = int64(float64(used) / float64(total) * 100)
		}

		p = datapoint.DataPoint{}
		p.Metric = keyPrefix + ".total"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = total
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = keyPrefix + ".used.percent"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = usedPercent
		_total := helpers.Format4HumanSize(float64(total))
		p.Tags = map[string]interface{}{
			"total": _total,
		}
		points = append(points, p)

	}

	return points, nil

}

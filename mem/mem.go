package metricsMem

import (
	"strconv"
	"strings"

	"github.com/MonitorMetrics/base/helpers"
)

func Gets() (result []map[string]interface{}, err error) {
	m := []map[string]interface{}{}
	cmd := `free -b`

	timeoutInSeconds := 1
	out, err := helpers.ExecCommand(cmd, timeoutInSeconds)
	if err != nil {
		return m, err
	}

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

		m = append(m, map[string]interface{}{
			"k": keyPrefix + ".total",
			"v": total,
		})

		m = append(m, map[string]interface{}{
			"k": keyPrefix + ".used.percent",
			"v": usedPercent,
		})

	}

	return m, nil

}

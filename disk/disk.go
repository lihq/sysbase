package metricsDisk

import (
	"log"
	"strconv"
	"strings"

	"github.com/MonitorMetrics/base/helpers"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

func Gets() (result []map[string]interface{}, err error) {
	m := []map[string]interface{}{}

	cmd := `df -h |grep --color=never '^/dev'`

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
		// Filesystem Size Used Avail Use% Mounted on
		if len(fields) != 6 {
			log.Println("parse df output failed")
			continue
		}

		fs := fields[0]
		size := fields[1]
		//used := fields[2]
		//avail := fields[3]
		usedPercentStr := fields[4]
		mountedOn := fields[5]

		sizeBytes := int64(-1)
		usedPercent := int64(-1)

		if strings.HasSuffix(size, "M") {
			prefix := strings.TrimRight(size, "M")
			sizeFloat, err := strconv.ParseFloat(prefix, 32)
			if err != nil {
				continue
			}
			sizeBytes = int64(sizeFloat) * MB

		} else if strings.HasSuffix(size, "G") {
			prefix := strings.TrimRight(size, "G")
			sizeFloat, err := strconv.ParseFloat(prefix, 32)
			if err != nil {
				continue
			}
			sizeBytes = int64(sizeFloat) * GB

		} else if strings.HasSuffix(size, "T") {
			prefix := strings.TrimRight(size, "T")
			sizeFloat, err := strconv.ParseFloat(prefix, 32)
			if err != nil {
				continue
			}
			sizeBytes = int64(sizeFloat) * TB

		} else {
			log.Println("parse df size failed")
			continue
		}

		if strings.HasSuffix(usedPercentStr, "%") {
			prefix := strings.TrimRight(usedPercentStr, "%")
			usedPercentUint, err := strconv.ParseUint(prefix, 10, 32)
			if err != nil {
			}

			usedPercent = int64(usedPercentUint)
		}

		m = append(m, map[string]interface{}{
			"k": "disk.total",
			"v": sizeBytes,
			"t": map[string]interface{}{
				"src": fs,
				"dst": mountedOn,
			},
		})

		m = append(m, map[string]interface{}{
			"k": "disk.used.percent",
			"v": usedPercent,
			"t": map[string]interface{}{
				"src": fs,
				"dst": mountedOn,
			},
		})

	}

	return m, nil
}

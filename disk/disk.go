package metricsDisk

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/MonitorMetrics/base/helpers"
	"github.com/MonitorMetrics/base/models"
)

func Gets() (result []datapoint.DataPoint, err error) {
	points := []datapoint.DataPoint{}

	cmd := `df -h |grep --color=never '^/dev'`

	timeoutInSeconds := 1
	out, err := helpers.ExecCommand(cmd, timeoutInSeconds)
	if err != nil {
		return points, err
	}

	var p datapoint.DataPoint
	now := time.Now()

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

		sizeBytes, err := helpers.ParseSize(size)
		if err != nil {
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

		p = datapoint.DataPoint{}
		p.Metric = "disk.total"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = sizeBytes
		p.Timestamp = now
		p.Tags = map[string]interface{}{
			"src": fs,
			"dst": mountedOn,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "disk.used.percent"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = usedPercent
		p.Timestamp = now
		p.Tags = map[string]interface{}{
			"src":   fs,
			"dst":   mountedOn,
			"total": size,
		}
		points = append(points, p)
	}

	return points, nil
}

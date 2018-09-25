package metricsDisk

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/MonitorMetrics/base/helpers"
	"github.com/MonitorMetrics/base/models"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
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
			"src": fs,
			"dst": mountedOn,
		}
		points = append(points, p)
	}

	return points, nil
}

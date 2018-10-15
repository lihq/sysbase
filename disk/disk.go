// package metricsDisk implements collect harddisk related monitor metrics.
package metricsDisk

import (
	"time"

	"github.com/MonitorMetrics/base/helpers"
	"github.com/MonitorMetrics/base/models"
)

// Gets returns disk total bytes and used percent from `df -h` and %util from `iostat`(sysstat).
func Gets() (result []datapoint.DataPoint, err error) {
	points := []datapoint.DataPoint{}
	now := time.Now()

	var cmd string
	var out []byte

	timeoutInSeconds := 1
	cmd = `df -h |grep --color=never '^/dev'`
	out, err = helpers.ExecCommand(cmd, timeoutInSeconds)
	if err != nil {
		return points, err
	}
	dfs := ParseOutputDf(string(out))
	var p datapoint.DataPoint

	for _, df := range dfs {
		p = datapoint.DataPoint{}
		p.Metric = "disk.total"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = df.SizeInBytes
		p.Timestamp = now
		p.Tags = map[string]interface{}{
			"src": df.Filesystem,
			"dst": df.MountedOn,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "disk.used.percent"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = df.UsedPercent
		p.Timestamp = now
		p.Tags = map[string]interface{}{
			"src":   df.Filesystem,
			"dst":   df.MountedOn,
			"total": df.Size,
		}
		points = append(points, p)
	}

	cmd = `iostat -xdm`
	out, err = helpers.ExecCommand(cmd, timeoutInSeconds)
	iostats := ParseOutputIostat(string(out))
	for _, iostat := range iostats {
		p = datapoint.DataPoint{}
		p.Metric = "disk.util"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = iostat.Util
		p.Timestamp = now
		p.Tags = map[string]interface{}{
			"device": iostat.Device,
		}
		points = append(points, p)

	}

	return points, nil
}

package metricsTop

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/MonitorMetrics/base/helpers"
	"github.com/MonitorMetrics/base/models"
)

const (
	columnsTopOuput = 12
	linesHeader     = 6
	linesTopN       = 10
)

func Gets() ([]datapoint.DataPoint, error) {
	points := []datapoint.DataPoint{}

	timeoutInSeconds := 4

	cmdStr := fmt.Sprintf(`top -c -b -d 3 -n 1 | head -n %d`, linesHeader+linesTopN+1)

	out, err := helpers.ExecCommand(cmdStr, timeoutInSeconds)
	if err != nil {
		log.Println("helpers.ExecCommand failed", err)
		return points, err
	}

	points = parseOutput(out)
	return points, nil

}

func parseOutput(output []byte) []datapoint.DataPoint {
	points := []datapoint.DataPoint{}

	lines := strings.Split(string(output), "\n")

	now := time.Now()
	var p datapoint.DataPoint
	for lineno, line := range lines {
		if lineno < linesHeader+1 {
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Index(line, "grep") != -1 {
			continue
		}

		columns := Fields(line)

		if len(columns) != columnsTopOuput {
			log.Println("got unexpected top output", fmt.Sprintf("-%#v-", line))
			continue
		}

		percentCpuI := columns[8]
		percentCpu, errCpu := strconv.ParseFloat(percentCpuI, 32)

		percentMemI := columns[9]
		percentMem, errMem := strconv.ParseFloat(percentMemI, 32)
		if errCpu != nil || errMem != nil {
			log.Println("parse percentCpu or percentMem failed", percentCpuI, percentMemI, errCpu, errMem)
			continue
		}

		procCMD := columns[11]
		procName := parseProcName(procCMD)

		p = datapoint.DataPoint{}
		p.Metric = "top.cpu"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = percentCpu
		p.Timestamp = now
		p.Tags = map[string]interface{}{
			"cmd":  procCMD,
			"proc": procName,
			"no":   lineno - linesHeader,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "top.mem"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = percentMem
		p.Timestamp = now
		p.Tags = map[string]interface{}{
			"cmd":  procCMD,
			"proc": procName,
			"no":   lineno - linesHeader,
		}

		points = append(points, p)

	}

	return points

}

func parseProcName(cmd string) (proc string) {
	fields := strings.Fields(cmd)
	if len(fields) > 0 {
		s := fields[0]
		if strings.Index(s, "[") == -1 {
			proc = path.Base(s)
		} else {
			proc = s
		}
	}
	return
}

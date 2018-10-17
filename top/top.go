package metricsTop

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/MonitorMetrics/base/helpers"
	"github.com/MonitorMetrics/base/models"
)

const (
	columnsTopOuput = 12
	linesHeader     = 6
	linesTopN       = 10
)

// Gets returns CPU and memory usage from `top`.
func Gets() ([]datapoint.DataPoint, error) {
	points := []datapoint.DataPoint{}

	timeoutInSeconds := 4

	cmdStr := fmt.Sprintf(`top -c -b -d 3 -n 1`)

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

	var p datapoint.DataPoint
	n := 0
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

		// skip kernel related processes
		if strings.Index(procName, "[") != -1 {
			continue
		}

		p = datapoint.DataPoint{}
		p.Metric = "top.cpu"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = percentCpu
		p.Tags = map[string]interface{}{
			"proc": procName,
			//"no":   n,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "top.mem"
		p.ContentType = datapoint.ContentTypeGauge
		p.Value = percentMem
		p.Tags = map[string]interface{}{
			"proc": procName,
			//"no":   n,
		}

		points = append(points, p)

		n += 1
		if n >= linesTopN {
			break
		}

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

package metricsTop

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/MonitorMetrics/base/helpers"
)

const (
	columnsTopOuput = 12
	linesHeader     = 6
	linesTopN       = 10
)

func Gets() ([]map[string]interface{}, error) {
	empty := []map[string]interface{}{}

	timeoutInSeconds := 4

	cmdStr := fmt.Sprintf(`top -c -b -d 3 -n 1 | head -n %d`, linesHeader+linesTopN+1)

	out, err := helpers.ExecCommand(cmdStr, timeoutInSeconds)
	if err != nil {
		log.Println("helpers.ExecCommand failed", err)
		return empty, err
	}
	fmt.Println(string(out))

	m := parseOutput(out)

	return m, nil

}

func parseOutput(output []byte) []map[string]interface{} {
	m := []map[string]interface{}{}

	lines := strings.Split(string(output), "\n")
	for lineno, line := range lines {
		if lineno < linesHeader-1 {
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

		m = append(m, map[string]interface{}{
			"k": "top.cpu",
			"v": percentCpu,
			"t": map[string]interface{}{
				"cmd":  procCMD,
				"proc": procName,
				"no":   lineno - linesHeader,
			},
		})

		m = append(m, map[string]interface{}{
			"k": "top.mem",
			"v": percentMem,
			"t": map[string]interface{}{
				"cmd":  procCMD,
				"proc": procName,
				"no":   lineno - linesHeader,
			},
		})

	}

	return m

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

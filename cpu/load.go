package metricsCPU

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

func Gets() (result []map[string]interface{}, err error) {
	m := []map[string]interface{}{}

	if runtime.GOOS != "linux" {
		err = errors.New(fmt.Sprintf("platform %s not support", runtime.GOOS))
		return m, err
	}

	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return m, err
	}

	fields := strings.Fields(string(b))

	if loadMin1, err := strconv.ParseFloat(fields[0], 64); err != nil {
		return m, err
	} else {
		item := map[string]interface{}{
			"k": "load.1min",
			"v": loadMin1,
		}
		m = append(m, item)
	}

	if loadMin5, err := strconv.ParseFloat(fields[1], 64); err != nil {
		return m, err
	} else {
		item := map[string]interface{}{
			"k": "load.5min",
			"v": loadMin5,
		}
		m = append(m, item)
	}

	if loadMin15, err := strconv.ParseFloat(fields[2], 64); err != nil {
		return m, err
	} else {
		item := map[string]interface{}{
			"k": "load.15min",
			"v": loadMin15,
		}
		m = append(m, item)
	}

	return m, nil

}

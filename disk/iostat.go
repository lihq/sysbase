// package metricsDisk implements collect harddisk related monitor metrics.
package metricsDisk

import (
	"strconv"
	"strings"
)

type Iostat struct {
	Device string
	Util   uint8
}

func ParseOutputIostat(output string) []Iostat {
	var iostats []Iostat

	output = strings.TrimSpace(output)
	for lineno, line := range strings.Split(output, "\n") {
		// skip header
		if lineno < 3 {
			continue
		}

		line = strings.TrimSpace(line)

		fields := strings.Fields(line)
		if len(fields) != 14 {
			continue
		}
		device := fields[0]
		if device == "" {
			continue
		}
		utilS := fields[13]
		util, err := strconv.ParseFloat(utilS, 32)
		if err != nil {
			continue
		}
		var iostat Iostat
		iostat.Device = device
		iostat.Util = uint8(util)

		iostats = append(iostats, iostat)

	}

	return iostats
}

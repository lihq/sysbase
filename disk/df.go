package metricsDisk

import (
	"log"
	"strconv"
	"strings"

	"github.com/OpenCCTV/sys_base/helpers"
)

type Df struct {
	Filesystem  string
	Size        string
	SizeInBytes float64
	UsedPercent uint8
	MountedOn   string
}

func ParseOutputDf(output string) []Df {
	var dfs []Df
	var err error

	for _, line := range strings.Split(output, "\n") {
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

		var df Df
		df.Filesystem = fs
		df.Size = size
		df.MountedOn = mountedOn

		df.SizeInBytes, err = helpers.ParseSize(size)
		if err != nil {
			log.Println("parse df size failed")
			continue
		}

		if strings.HasSuffix(usedPercentStr, "%") {
			prefix := strings.TrimRight(usedPercentStr, "%")
			usedPercentUint, err := strconv.ParseUint(prefix, 10, 32)
			if err != nil {
				log.Println("parse df used percent failed")
				continue
			}

			df.UsedPercent = uint8(usedPercentUint)
		}

		dfs = append(dfs, df)
	}
	return dfs

}

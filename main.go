// Example collect all system base monitor metrics.
package main

import (
	"encoding/json"
	"log"

	"github.com/OpenCCTV/sysbase/net"

	"github.com/OpenCCTV/sysbase/utils"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	dataPoints := utils.GetMergeAll()

	lans, err := metricsNet.GetLANIpAddrs()
	if err != nil {
		log.Fatalln(err)
	}
	wans, err := metricsNet.GetWANIpAddrs()
	if err != nil {
		log.Fatalln(err)
	}

	result := map[string]interface{}{
		"metrics": dataPoints,
		"lans":    lans,
		"wans":    wans,
	}
	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(string(out))
	}

}

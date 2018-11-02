package metricsNet

import (
	"log"
	"strconv"
	"strings"

	"github.com/OpenCCTV/sys_base/helpers"
)

type NetDev struct {
	Face string

	RxPackets uint64
	RxBytes   uint64
	RxErrs    uint64
	RxDrop    uint64

	TxPackets uint64
	TxBytes   uint64
	TxErrs    uint64
	TxDrop    uint64
}

var LastNetDevs = map[string]NetDev{}

func ParseProcNetDev(output string, facesWhitelist []string) []NetDev {
	var netDevs []NetDev
	var err error

	output = strings.TrimSpace(output)
	for lineno, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if lineno < 2 {
			continue
		}

		fields := strings.Fields(line)
		expectedColumns := 17

		if len(fields) != expectedColumns {
			continue
		}

		face := strings.TrimSuffix(fields[0], ":")
		if face == "lo" {
			continue
		}

		if !helpers.StringInSlice(face, facesWhitelist) {
			continue
		}

		netdev := NetDev{}
		netdev.Face = face

		// RX/in
		netdev.RxBytes, err = strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			log.Println(err)
		}

		netdev.RxPackets, err = strconv.ParseUint(fields[2], 10, 64)
		if err != nil {
			log.Println(err)
		}

		netdev.RxErrs, err = strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			log.Println(err)
		}

		netdev.RxDrop, err = strconv.ParseUint(fields[4], 10, 64)
		if err != nil {
			log.Println(err)
		}

		// TX/out
		netdev.TxBytes, err = strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			log.Println(err)
		}

		netdev.TxPackets, err = strconv.ParseUint(fields[10], 10, 64)
		if err != nil {
			log.Println(err)
		}

		netdev.TxErrs, err = strconv.ParseUint(fields[11], 10, 64)
		if err != nil {
			log.Println(err)
		}

		netdev.TxDrop, err = strconv.ParseUint(fields[12], 10, 64)
		if err != nil {
			log.Println(err)
		}

		netDevs = append(netDevs, netdev)

	}
	return netDevs
}

func ApplyDiff(netDevs *[]NetDev) {
	for idx, netdev := range *netDevs {
		netdevOld, ok := LastNetDevs[netdev.Face]
		if ok {
			(*netDevs)[idx].RxPackets = (*netDevs)[idx].RxPackets - netdevOld.RxPackets
			(*netDevs)[idx].RxBytes = (*netDevs)[idx].RxBytes - netdevOld.RxBytes
			(*netDevs)[idx].RxErrs = (*netDevs)[idx].RxErrs - netdevOld.RxErrs
			(*netDevs)[idx].RxDrop = (*netDevs)[idx].RxDrop - netdevOld.RxDrop

			(*netDevs)[idx].TxPackets = (*netDevs)[idx].TxPackets - netdevOld.TxPackets
			(*netDevs)[idx].TxBytes = (*netDevs)[idx].TxBytes - netdevOld.TxBytes
			(*netDevs)[idx].TxErrs = (*netDevs)[idx].TxErrs - netdevOld.TxErrs
			(*netDevs)[idx].TxDrop = (*netDevs)[idx].TxDrop - netdevOld.TxDrop
		} else {
			(*netDevs)[idx].RxPackets = 0
			(*netDevs)[idx].RxBytes = 0
			(*netDevs)[idx].RxErrs = 0
			(*netDevs)[idx].RxDrop = 0

			(*netDevs)[idx].TxPackets = 0
			(*netDevs)[idx].TxBytes = 0
			(*netDevs)[idx].TxErrs = 0
			(*netDevs)[idx].TxDrop = 0
		}
		LastNetDevs[netdev.Face] = netdev

	}
}

// GetAliveIfaces returns alive network interfaces, it does not contains interface aliases.
func GetAliveIfaces() ([]string, error) {
	result := []string{}
	timeoutInSeconds := 1
	cmd := `ifconfig |grep --color=never -o ^[a-z0-9]*`
	b, err := helpers.ExecCommand(cmd, timeoutInSeconds)
	if err != nil {
		return result, err
	}

	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		result = append(result, line)
	}

	return result, nil
}

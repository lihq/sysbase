package metricsNet

import (
	"io/ioutil"

	"github.com/MonitorMetrics/base/models"
)

func Gets() (result []datapoint.DataPoint, err error) {
	points := []datapoint.DataPoint{}

	b, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		return points, err
	}

	facesWhitelist, err := GetAliveIfaces()
	if err != nil {
		return points, err
	}

	netDevs := ParseProcNetDev(string(b), facesWhitelist)
	ApplyDiff(&netDevs)

	for _, netdev := range netDevs {
		var p datapoint.DataPoint

		p = datapoint.DataPoint{}
		p.Metric = "net.in.packets"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.RxPackets
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "net.in.bytes"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.RxBytes
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "net.in.errs"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.RxErrs
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "net.in.drop"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.RxDrop
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

		// out
		p = datapoint.DataPoint{}
		p.Metric = "net.out.packets"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.TxPackets
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "net.out.bytes"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.TxBytes
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "net.out.errs"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.TxErrs
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

		p = datapoint.DataPoint{}
		p.Metric = "net.out.drop"
		p.ContentType = datapoint.ContentTypeDiff
		p.Value = netdev.TxDrop
		p.Tags = map[string]interface{}{
			"face": netdev.Face,
		}
		points = append(points, p)

	}

	return points, nil
}

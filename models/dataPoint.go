package datapoint

import (
	"encoding/json"
	"time"
)

const (
	ContentTypeGauge    = 1
	ContentTypeDiff     = 2 // diff = currValue - lastValue
	ContentTypeString   = 3 //
	ContentTypeReversed = 4
)

type DataPoint struct {
	Endpoint    string                 `json:"endpoint" db:"endpoint"`
	ContentType uint8                  `json:"contentType" db:"ContentType"`
	Metric      string                 `json:"metric" db:"metric"`
	Value       interface{}            `json:"value" db:"value"`
	Tags        map[string]interface{} `json:"tags" db:"tags"`
	Timestamp   time.Time              `json:"ts" db:"ts"`
}

func (this *DataPoint) Tags2str() (result string) {
	out, err := json.Marshal(this.Tags)
	if err != nil {
	} else {
		result = string(out)
	}
	return result
}

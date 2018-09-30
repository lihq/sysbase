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
	Endpoint    string                 `json:"-" db:"endpoint"`
	ContentType uint8                  `json:"-" db:"ContentType"`
	Metric      string                 `json:"-" db:"metric"`
	Value       interface{}            `json:"value" db:"value"`
	Tags        map[string]interface{} `json:"tags" db:"tags"`
	Timestamp   time.Time              `json:"ts" db:"ts"`
}

func (this *DataPoint) MarshalJSON() ([]byte, error) {
	type Alias DataPoint
	return json.Marshal(&struct {
		*Alias
		Timestamp int64 `json:"ts"`
	}{
		Alias:     (*Alias)(this),
		Timestamp: this.Timestamp.Unix(),
	})
}

func (this *DataPoint) Tags2str() (result string) {
	out, err := json.Marshal(this.Tags)
	if err != nil {
	} else {
		result = string(out)
	}
	return result
}

package metricsMem

import (
	"testing"
)

func TestParseOutputFree(t *testing.T) {
	cases := []struct {
		Sample          string
		MemTotal        int64
		MemUsedPercent  int64
		SwapTotal       int64
		SwapUsedPercent int64
	}{
		{`			  total        used        free      shared  buff/cache   available
Mem:      131771376    11178184    38897980        9140    81695212   118591828
Swap:             0           0           0
`, 131771376, 8, 0, 0},
	}

	for _, c := range cases {
		got, err := ParseOutputFree(c.Sample)
		if err != nil {
			t.Errorf("ParseOutputFree(%v) expected err=nil, got %v", c.Sample, err)
		}

		if got.MemTotal != c.MemTotal ||
			got.MemUsedPercent != c.MemUsedPercent ||
			got.SwapTotal != c.SwapTotal ||
			got.SwapUsedPercent != c.SwapUsedPercent {
			t.Errorf("ParseOutputFree(%v) expected != got %v", c.Sample, got)
		}

	}

}

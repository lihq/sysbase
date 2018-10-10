package helpers

import (
	"math"
	"testing"
)

func TestFormat4HumanSize(t *testing.T) {
	cases := []struct {
		sample   float64
		expected string
	}{
		{float64(0), "0B"},
		{float64(1024), "1KB"},
		{float64(1025), "1KB"},
		{float64(1024)*1024 + 1, "1MB"},
	}

	for _, c := range cases {
		got := Format4HumanSize(c.sample)
		if c.expected != got {
			t.Errorf("Format4HumanSize(%v) expectd %v, got %v", c.sample, c.expected, got)
		}

	}

}

func equalFloat(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func TestParseSize(t *testing.T) {
	cases := []struct {
		sample   string
		expected float64
	}{
		{"1KB", float64(1024)},
		{"1.6T", float64(1.6 * TB)},
		{"69G", float64(69 * GB)},
	}

	for _, c := range cases {
		got, err := ParseSize(c.sample)
		if err != nil {
			t.Errorf("Format4HumanSize(%v) expectd ok, got error %v", c.sample, err)
		}

		if !equalFloat(c.expected, got) {
			t.Errorf("Format4HumanSize(%v) expectd %v, got %v", c.sample, c.expected, got)
		}

	}

}

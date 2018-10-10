package helpers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

// Format4HumanSize converts bytes from float64 into human-read size in string.
func Format4HumanSize(size float64) (result string) {
	suffix := []string{
		"B",
		"KB",
		"MB",
		"GB",
		"TB",
	}

	i := 0
	for size >= 1024 && i < len(suffix)-1 {
		size /= 1024
		i += 1
	}

	// decimal is 0, format x.0 => x
	if int(size*10)%10 == 0 {
		result = fmt.Sprintf("%.0f%s", size, suffix[i])
	} else {
		result = fmt.Sprintf("%.1f%s", size, suffix[i])
	}
	return

}

// ParseSize converts human-read size from string into bytes in float64.
func ParseSize(size string) (sizeBytes float64, err error) {
	if strings.HasSuffix(size, "K") || strings.HasSuffix(size, "KB") {
		prefix := strings.TrimRight(size, "KB")
		prefix = strings.TrimRight(prefix, "K")
		sizeFloat, err := strconv.ParseFloat(prefix, 64)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = float64(sizeFloat) * KB
	} else if strings.HasSuffix(size, "M") || strings.HasSuffix(size, "MB") {
		prefix := strings.TrimRight(size, "MB")
		prefix = strings.TrimRight(prefix, "M")
		sizeFloat, err := strconv.ParseFloat(prefix, 64)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = float64(sizeFloat) * MB

	} else if strings.HasSuffix(size, "G") || strings.HasSuffix(size, "GB") {
		prefix := strings.TrimRight(size, "GB")
		prefix = strings.TrimRight(prefix, "G")
		sizeFloat, err := strconv.ParseFloat(prefix, 64)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = float64(sizeFloat) * GB

	} else if strings.HasSuffix(size, "T") || strings.HasSuffix(size, "TB") {
		prefix := strings.TrimRight(size, "TB")
		prefix = strings.TrimRight(prefix, "T")
		sizeFloat, err := strconv.ParseFloat(prefix, 64)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = float64(sizeFloat) * TB
	} else {
		err = errors.New("parse df size failed")
		return sizeBytes, err
	}

	return sizeBytes, nil
}

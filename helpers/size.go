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

func Format4HumanSize(size float64) string {
	suffix := []string{
		"B",
		"KB",
		"MB",
		"GB",
		"TB",
	}

	i := 0
	for size > 1024 && i < len(suffix)-1 {
		size /= 1024
		i += 1
	}

	return fmt.Sprintf("%.0f%s", size, suffix[i])
}

func ParseSize(size string) (sizeBytes int64, err error) {
	if strings.HasSuffix(size, "K") || strings.HasSuffix(size, "KB") {
		prefix := strings.TrimRight(size, "KB")
		prefix = strings.TrimRight(prefix, "K")
		sizeFloat, err := strconv.ParseFloat(prefix, 32)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = int64(sizeFloat) * KB
	} else if strings.HasSuffix(size, "M") || strings.HasSuffix(size, "MB") {
		prefix := strings.TrimRight(size, "MB")
		prefix = strings.TrimRight(prefix, "M")
		sizeFloat, err := strconv.ParseFloat(prefix, 32)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = int64(sizeFloat) * MB

	} else if strings.HasSuffix(size, "G") || strings.HasSuffix(size, "GB") {
		prefix := strings.TrimRight(size, "GB")
		prefix = strings.TrimRight(prefix, "G")
		sizeFloat, err := strconv.ParseFloat(prefix, 32)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = int64(sizeFloat) * GB

	} else if strings.HasSuffix(size, "T") || strings.HasSuffix(size, "TB") {
		prefix := strings.TrimRight(size, "TB")
		prefix = strings.TrimRight(prefix, "T")
		sizeFloat, err := strconv.ParseFloat(prefix, 32)
		if err != nil {
			return sizeBytes, err
		}
		sizeBytes = int64(sizeFloat) * TB
	} else {
		err = errors.New("parse df size failed")
		return sizeBytes, err
	}

	return sizeBytes, nil
}

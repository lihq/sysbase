package metricsTop

import (
	"strings"
	"unicode"
)

// Fields customs strings.Fields for parse linux top output does not break command column into parts.
func Fields(line string) []string {
	// default columns on linux: PID USER PR NI VIRT RES SHR S %CPU %MEM TIME+ COMMAND
	// USER: Effective User Name
	// PR: Priority
	// NI: Nice Value
	// VIRT: Virtual Image (KiB)
	// RES: Resident Size (KiB)
	// S: Process Status
	// %CPU: CPU Usage
	// %MEM: Memory Usage (RES)
	// TIME+: CPU Time, hundredths
	// COMMAND: Command Name/Line <= 12th column

	idx := 0
	fields := []string{}
	for cur, chr := range line {
		if unicode.IsSpace(chr) {
			field := strings.TrimSpace(line[idx:cur])
			if field != "" {
				fields = append(fields, field)
				idx = cur
			}
		} else {
			continue
		}

		if len(fields) == 11 {
			break
		}
	}

	field := strings.TrimSpace(line[idx:])
	fields = append(fields, field)

	return fields
}

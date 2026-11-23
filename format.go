package code

import (
	"fmt"
)

func formatSize(size int64, humanReadable bool) string {
	if !humanReadable {
		return fmt.Sprintf("%dB", size)
	}
	
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
    value := float64(size)
    unitIndex := 0

	for unitIndex < len(units)-1 && value >= 1024 {
        value /= 1024
        unitIndex++
    }

    if unitIndex == 0 {
        return fmt.Sprintf("%.0fB", value)
    }

    return fmt.Sprintf("%.1f%s", value, units[unitIndex])
}

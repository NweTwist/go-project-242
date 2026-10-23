package goproject242

import (
	"fmt"
)

func FormatSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.1fKB", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2fMB", float64(size)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2fGB", float64(size)/(1024*1024*1024))
	}
}

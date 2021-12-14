package main

import (
	"fmt"
	"time"
)

func formatDate(t time.Time) string {
	suffix := "th"
	switch t.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}
	return t.Format("Monday 2" + suffix + " January")
}
func main() {
	for _, n := range []int{1, 2, 3, 4, 12, 13, 21, 22} {
		t := time.Date(2015, 3, n, 1, 1, 1, 1, time.UTC)
		fmt.Println(formatDate(t))
	}
}

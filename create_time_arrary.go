package main

import (
	"fmt"
	"time"
)

func main() {
	// https://blog.boot.dev/golang/golang-date-time/
	start := time.Date(2022, 11, 9, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, 11, 22, 0, 0, 0, 0, time.UTC)
	difference := end.Sub(start)

	s := make([]time.Time, 0)

	for difference != 0 {
		s = append(s, start)
		start = start.Add(24 * time.Hour)
		difference = end.Sub(start)
	}

	for _, t := range s {
		fmt.Println(t)
	}

}

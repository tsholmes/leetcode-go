package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31"))
}

func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse(time.RFC3339, date1+`T00:00:00Z`)
	d2, _ := time.Parse(time.RFC3339, date2+`T00:00:00Z`)
	ds := int(math.Round(float64(d2.Sub(d1)) / float64(24*time.Hour)))
	if ds < 0 {
		ds = -ds
	}
	return ds
}

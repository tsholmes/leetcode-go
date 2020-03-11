package main

import "math"

func main() {

}

func findPoisonedDuration(timeSeries []int, duration int) int {
	timeSeries = append(timeSeries, math.MaxInt64)

	last := timeSeries[0]
	total := 0

	for _, t := range timeSeries {
		dur := t - last
		if dur > duration {
			dur = duration
		}
		total += dur
		last = t
	}
	return total
}

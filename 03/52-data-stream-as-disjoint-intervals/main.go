package main

import (
	"math"
	"sort"
)

func main() {

}

type SummaryRanges struct {
	starts []int
	ends   []int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (s *SummaryRanges) AddNum(val int) {
	ri := sort.Search(len(s.ends), func(i int) bool { return s.ends[i] >= val })
	if ri == len(s.ends) {
		s.starts = append(s.starts, val)
		s.ends = append(s.ends, val)
	} else if val >= s.starts[ri] {
		// no-op
	} else {
		// val < s.starts[ri]
		s.starts = append(s.starts, 0)
		s.ends = append(s.ends, 0)
		copy(s.starts[ri+1:], s.starts[ri:])
		copy(s.ends[ri+1:], s.ends[ri:])
		s.starts[ri] = val
		s.ends[ri] = val
	}
	count := 0
	lastEnd := math.MinInt64
	for i := range s.starts {
		if s.starts[i] <= lastEnd+1 {
			s.ends[count-1] = s.ends[i]
		} else {
			s.starts[count], s.ends[count] = s.starts[i], s.ends[i]
			count++
		}
		lastEnd = s.ends[i]
	}
	s.starts = s.starts[:count]
	s.ends = s.ends[:count]
}

func (s *SummaryRanges) GetIntervals() [][]int {
	var res [][]int
	for i := range s.starts {
		res = append(res, []int{s.starts[i], s.ends[i]})
	}
	return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */

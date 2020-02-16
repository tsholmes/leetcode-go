package main

import "fmt"

func main() {
	t := Constructor()

	t.RecordTweet("a", 0)
	t.RecordTweet("a", 60)
	t.RecordTweet("a", 10)
	fmt.Println(t.GetTweetCountsPerFrequency("minute", "a", 0, 59))
	fmt.Println(t.GetTweetCountsPerFrequency("minute", "a", 0, 60))
}

type TweetCounts struct {
	counts map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		counts: map[string][]int{},
	}
}

func (t *TweetCounts) RecordTweet(tweetName string, time int) {
	t.counts[tweetName] = append(t.counts[tweetName], time)
}

func (t *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	var f int
	switch freq {
	case "minute":
		f = 60
	case "hour":
		f = 60 * 60
	case "day":
		f = 60 * 60 * 24
	}

	res := make([]int, (endTime-startTime+f)/f)

	for _, time := range t.counts[tweetName] {
		if time < startTime || time > endTime {
			continue
		}
		idx := (time - startTime) / f
		res[idx]++
	}

	return res
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */

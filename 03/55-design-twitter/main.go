package main

func main() {

}

type Twitter struct {
	follows map[[2]int]bool
	tweets  [][2]int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{follows: map[[2]int]bool{}}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.tweets = append(t.tweets, [2]int{userId, tweetId})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	var res []int
	for i := len(t.tweets) - 1; i >= 0 && len(res) < 10; i-- {
		tw := t.tweets[i]
		if tw[0] == userId || t.follows[[2]int{userId, tw[0]}] {
			res = append(res, tw[1])
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	t.follows[[2]int{followerId, followeeId}] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	t.follows[[2]int{followerId, followeeId}] = false
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

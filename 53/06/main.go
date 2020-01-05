package main

import "fmt"

func main() {
	fmt.Println(minInsertions("mbadm"))
	fmt.Println(minInsertions("leetcode"))
}

func minInsertions(s string) int {
	fwd := []byte(s)
	back := []byte(s)
	for i := 0; i < len(back)/2; i++ {
		j := len(back) - i - 1
		back[i], back[j] = back[j], back[i]
	}
	N := len(s)

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i == 0 || j == 0 {
				if fwd[i] == back[j] {
					dp[i][j] = 1
				}
				if i > 0 && dp[i-1][j] == 1 {
					dp[i][j] = 1
				}
				if j > 0 && dp[i][j-1] == 1 {
					dp[i][j] = 1
				}
				continue
			}
			add := 0
			if fwd[i] == back[j] {
				add = 1
			}
			max := dp[i-1][j-1] + add
			if dp[i-1][j] > max {
				max = dp[i-1][j]
			}
			if dp[i][j-1] > max {
				max = dp[i][j-1]
			}
			dp[i][j] = max
		}
	}

	return len(s) - dp[N-1][N-1]
}

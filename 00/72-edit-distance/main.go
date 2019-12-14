package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				min := dp[i-1][j] + 1
				if v := dp[i][j-1] + 1; v < min {
					min = v
				}
				rep := dp[i-1][j-1]
				if word1[i-1] != word2[j-1] {
					rep++
				}
				if rep < min {
					min = rep
				}
				dp[i][j] = min
			}
		}
	}

	return dp[n][m]
}

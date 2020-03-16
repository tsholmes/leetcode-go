package main

func main() {

}

func minDistance(word1 string, word2 string) int {
	N := lcsubseqN(word1, word2)
	return len(word1) + len(word2) - N - N
}

func lcsubseqN(a string, b string) int {
	N, M := len(a), len(b)
	if N == 0 || M == 0 {
		return 0
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
		for j := 0; j < M; j++ {
			add := 0
			if a[i] == b[j] {
				add = 1
			}
			if i == 0 || j == 0 {
				dp[i][j] = add
			} else {
				dp[i][j] = add + dp[i-1][j-1]
			}
			if i > 0 && dp[i-1][j] > dp[i][j] {
				dp[i][j] = dp[i-1][j]
			}
			if j > 0 && dp[i][j-1] > dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[N-1][M-1]
}

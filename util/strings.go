package util

// This is meant to be copied into a solution

func reverse(s string) string {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, s[len(s)-i-1])
	}
	return string(res)
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

func lcsubseq(a string, b string) string {
	N, M := len(a), len(b)
	if N == 0 || M == 0 {
		return ""
	}

	dp := make([][]int, N)
	state := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
		state[i] = make([]int, M)
		for j := 0; j < M; j++ {
			add := 0
			state[i][j] = 1
			if i > 0 {
				state[i][j] = -1
			}
			if a[i] == b[j] {
				add = 1
				state[i][j] = 0
			}
			if i == 0 || j == 0 {
				dp[i][j] = add
			} else {
				dp[i][j] = add + dp[i-1][j-1]
			}
			if i > 0 && dp[i-1][j] > dp[i][j] {
				dp[i][j] = dp[i-1][j]
				state[i][j] = -1
			}
			if j > 0 && dp[i][j-1] > dp[i][j] {
				dp[i][j] = dp[i][j-1]
				state[i][j] = 1
			}
		}
	}

	var res []byte
	i, j := N-1, M-1
	for i >= 0 && j >= 0 {
		switch state[i][j] {
		case 0:
			res = append(res, a[i])
			i--
			j--
		case -1:
			i--
		case 1:
			j--
		}
	}
	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

package main

import "fmt"

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
	fmt.Println(isInterleave("a", "", "a"))
	fmt.Println(isInterleave("", "a", "a"))
	fmt.Println(isInterleave("", "a", "a"))
	fmt.Println(isInterleave("", "", "a"))
	fmt.Println(isInterleave("aaaaaaaaaaaaa", "aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	memo := map[[3]int]bool{}
	var search func(int, int, int) bool
	search = func(i, j, k int) bool {
		if k == len(s3) {
			return i == len(s1) && j == len(s2)
		}
		if i == len(s1) {
			return s2[j:] == s3[k:]
		}
		if j == len(s2) {
			return s1[i:] == s3[k:]
		}
		key := [3]int{i, j, k}
		res, ok := memo[key]
		if ok {
			return res
		}
		if s1[i] == s3[k] && search(i+1, j, k+1) {
			res = true
		} else if s2[j] == s3[k] && search(i, j+1, k+1) {
			res = true
		}
		memo[key] = res
		return res
	}
	return search(0, 0, 0)
}

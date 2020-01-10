package main

import "fmt"

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
	fmt.Println(removeDuplicateLetters("cruaebrnuzdmpfivugqejkspqvxxgnjixjtoboexjwcywzwptiahdbxkmhccsdnlmrmldwoxnurnlaiyzshimpzbmunvwhfkcvbeeorioxoxommgkjablxuibuxbuhhclgjwsgecuhvqscwutbownyjckhqlhjrdmtkozdwuewsxpupwhjeywznccjdeiisirvkvfroiyhhwuynmhwsdzmwauezxbssaxefktyufjnysvcmxrqxunoipqrbjxnxdwmeebpgucfxvvaansdpfetpipqynomtwkloczuepklwmhawfgovewnvxeqyghndlyoqxvoxwozfzprqwvcewvzjykyohfmywymudenrxwcoxrbsgctenzjxhqwtghlpnhkrjkxualiarouhscitxpmgabllajoqipvslibzxioocvvpdlwxvbvspezufenplebnajqsyixar"))
}

func removeDuplicateLetters(s string) string {
	N := len(s)
	after := make([]int, N)
	var all int
	for i := N - 1; i >= 0; i-- {
		j := s[i] - 'a'
		all |= 1 << j
		after[i] = all
	}

	var res []byte
	pos := 0
	for all > 0 {
		for j := 0; j < 26; j++ {
			bit := 1 << uint(j)
			if all&bit == 0 {
				continue
			}
			found := false
			for i := pos; i < N && all&^after[i] == 0; i++ {
				if int(s[i]-'a') == j {
					pos = i + 1
					found = true
					break
				}
			}
			if found {
				res = append(res, byte(j)+'a')
				all = all &^ bit
				break
			}
		}
	}

	return string(res)
}

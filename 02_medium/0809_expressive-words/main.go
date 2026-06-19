package main

// LeetCode #809: Expressive Words
// https://leetcode.com/problems/expressive-words/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ExpressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	fmt.Println(ExpressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
	fmt.Println(ExpressiveWords("abcd", []string{"abc"}))
}

func ExpressiveWords(s string, words []string) int {
	type group struct {
		char byte
		cnt  int
	}

	var encode func(string) []group
	encode = func(str string) []group {
		var groups []group
		for i := 0; i < len(str); {
			j := i
			for j < len(str) && str[j] == str[i] {
				j++
			}
			groups = append(groups, group{str[i], j - i})
			i = j
		}
		return groups
	}

	sGroups := encode(s)
	count := 0

	for _, word := range words {
		wGroups := encode(word)
		if len(wGroups) != len(sGroups) {
			continue
		}
		ok := true
		for i := range sGroups {
			if sGroups[i].char != wGroups[i].char {
				ok = false
				break
			}
			if sGroups[i].cnt < 3 && sGroups[i].cnt != wGroups[i].cnt {
				ok = false
				break
			}
			if sGroups[i].cnt >= 3 && wGroups[i].cnt > sGroups[i].cnt {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}

	return count
}

package main

// LeetCode #3692: Majority Frequency Characters
// https://leetcode.com/problems/majority-frequency-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MajorityFrequencyCharacters("aaabbbccdddde"))
	fmt.Println(MajorityFrequencyCharacters("abcd"))
	fmt.Println(MajorityFrequencyCharacters("pfpfgi"))
}

// Time: O(n)
// Space: O(1)
func MajorityFrequencyCharacters(s string) string {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	freq := make([]int, len(s)+1)
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			freq[cnt[i]]++
		}
	}

	bestSize := 0
	bestFreq := 0
	for k := 1; k <= len(s); k++ {
		if freq[k] > bestSize || (freq[k] == bestSize && k > bestFreq) {
			bestSize = freq[k]
			bestFreq = k
		}
	}

	res := make([]byte, 0)
	for i := 0; i < 26; i++ {
		if cnt[i] == bestFreq {
			res = append(res, byte('a'+i))
		}
	}
	return string(res)
}

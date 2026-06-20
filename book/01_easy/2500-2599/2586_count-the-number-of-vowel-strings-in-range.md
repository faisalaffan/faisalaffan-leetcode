# 2586 — Count The Number Of Vowel Strings In Range

## Deskripsi

**Soal:** [2586. Count The Number Of Vowel Strings In Range](https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2586: Count the Number of Vowel Strings in Range
// https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountTheNumberOfVowelStringsInRange([]string{"are", "amy", "u"}, 0, 2)) // 2
	fmt.Println(CountTheNumberOfVowelStringsInRange([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4)) // 3
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func CountTheNumberOfVowelStringsInRange(words []string, left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if len(words[i]) > 0 && isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			count++
		}
	}
	return count
}
```

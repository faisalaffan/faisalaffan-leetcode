# 2942 — Find Words Containing Character

## Deskripsi

**Soal:** [2942. Find Words Containing Character](https://leetcode.com/problems/find-words-containing-character/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * m) where m is max word length  
**Kompleksitas Ruang:** O(1) excluding output

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2942: Find Words Containing Character
// https://leetcode.com/problems/find-words-containing-character/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findWordsContaining
	fmt.Println(FindWordsContainingCharacter([]string{"leet", "code"}, 'e')) // [0, 1]
	fmt.Println(FindWordsContainingCharacter([]string{"abc", "bcd", "aaaa", "cbc"}, 'a')) // [0, 2]
	fmt.Println(FindWordsContainingCharacter([]string{"abc", "bcd", "aaaa", "cbc"}, 'z')) // []
}

// Time: O(n * m) where m is max word length | Space: O(1) excluding output
// LeetCode submission name: findWordsContaining
func FindWordsContainingCharacter(words []string, x byte) []int {
	result := []int{}
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if word[j] == x {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
```

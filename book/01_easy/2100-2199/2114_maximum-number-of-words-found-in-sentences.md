# 2114 — Maximum Number Of Words Found In Sentences

## Deskripsi

**Soal:** [2114. Maximum Number Of Words Found In Sentences](https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2114: Maximum Number of Words Found in Sentences
// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})) // 6
	fmt.Println(MaximumNumberOfWordsFoundInSentences([]string{"please wait", "continue to fight", "continue to win"}))                              // 3
}

// Time: O(n * m), Space: O(1)
func MaximumNumberOfWordsFoundInSentences(sentences []string) int {
	maxWords := 0
	for _, s := range sentences {
		count := strings.Count(s, " ") + 1
		if count > maxWords {
			maxWords = count
		}
	}
	return maxWords
}
```

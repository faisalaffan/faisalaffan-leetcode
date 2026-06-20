# 1002 — Find Common Characters

## Deskripsi

**Soal:** [1002. Find Common Characters](https://leetcode.com/problems/find-common-characters/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * m) where n = len(words), m = avg word length. Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1002: Find Common Characters
// https://leetcode.com/problems/find-common-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"})) // [e l l]
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))     // [c o]
}

// commonChars returns the common characters (with multiplicity) across all words.
// Time: O(n * m) where n = len(words), m = avg word length. Space: O(1).
func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	// Initialize with counts from first word
	freq := [26]int{}
	for _, c := range words[0] {
		freq[c-'a']++
	}

	for i := 1; i < len(words); i++ {
		currFreq := [26]int{}
		for _, c := range words[i] {
			currFreq[c-'a']++
		}
		for i := 0; i < 26; i++ {
			if currFreq[i] < freq[i] {
				freq[i] = currFreq[i]
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < freq[i]; j++ {
			result = append(result, string(rune('a'+i)))
		}
	}
	return result
}
```

# 2287 — Rearrange Characters To Make Target String

## Deskripsi

**Soal:** [2287. Rearrange Characters To Make Target String](https://leetcode.com/problems/rearrange-characters-to-make-target-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2287: Rearrange Characters to Make Target String
// https://leetcode.com/problems/rearrange-characters-to-make-target-string/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(RearrangeCharactersToMakeTargetString("ilovecodingonleetcode", "code")) // 2
	fmt.Println(RearrangeCharactersToMakeTargetString("abcba", "abc"))                  // 1
}

func RearrangeCharactersToMakeTargetString(s string, target string) int {
	sCount := [26]int{}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		sCount[s[i]-'a']++
	}

	tCount := [26]int{}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(target); i++ {
		tCount[target[i]-'a']++
	}

	maxCopies := len(s)
	for i := 0; i < 26; i++ {
		if tCount[i] > 0 {
			copies := sCount[i] / tCount[i]
			if copies < maxCopies {
				maxCopies = copies
			}
		}
	}
	return maxCopies
}
```

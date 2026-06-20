# 1880 — Check If Word Equals Summation Of Two Words

## Deskripsi

**Soal:** [1880. Check If Word Equals Summation Of Two Words](https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func IsSumEqual(firstWord string, secondWord string, targetWord string) bool`

## Solusi Go

```go
package main

// LeetCode #1880: Check if Word Equals Summation of Two Words
// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return wordValue(firstWord)+wordValue(secondWord) == wordValue(targetWord)
}

func wordValue(s string) int {
	val := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		val = val*10 + int(s[i]-'a')
	}
	return val
}

func main() {
	fmt.Println(IsSumEqual("acb", "cba", "cdb"))
	fmt.Println(IsSumEqual("aaa", "a", "aab"))
	fmt.Println(IsSumEqual("aaa", "a", "aaaa"))
}
```

# 0557 — Reverse Words In A String Iii

## Deskripsi

**Soal:** [0557. Reverse Words In A String Iii](https://leetcode.com/problems/reverse-words-in-a-string-iii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func ReverseWordsInAStringIii(s string) string`

## Solusi Go

```go
package main

// LeetCode #557: Reverse Words in a String III
// https://leetcode.com/problems/reverse-words-in-a-string-iii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseWordsInAStringIii(s string) string {
	b := []byte(s)
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			for lo, hi := start, i-1; lo < hi; lo, hi = lo+1, hi-1 {
				b[lo], b[hi] = b[hi], b[lo]
			}
			start = i + 1
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseWordsInAStringIii("Let's take LeetCode contest"))
	fmt.Println(ReverseWordsInAStringIii("Mr Ding"))
}
```

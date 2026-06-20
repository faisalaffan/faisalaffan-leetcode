# 1897 — Redistribute Characters To Make All Strings Equal

## Deskripsi

**Soal:** [1897. Redistribute Characters To Make All Strings Equal](https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * len), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MakeEqual(words []string) bool`

## Solusi Go

```go
package main

// LeetCode #1897: Redistribute Characters to Make All Strings Equal
// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
// Difficulty: Easy

import "fmt"

// Time: O(n * len), Space: O(1)
func MakeEqual(words []string) bool {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
	for _, w := range words {
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(w); i++ {
			freq[w[i]-'a']++
		}
	}
	n := len(words)
	for _, count := range freq {
		if count%n != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(MakeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Println(MakeEqual([]string{"ab", "a"}))
}
```

# 2053 — Kth Distinct String In An Array

## Deskripsi

**Soal:** [2053. Kth Distinct String In An Array](https://leetcode.com/problems/kth-distinct-string-in-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2053: Kth Distinct String in an Array
// https://leetcode.com/problems/kth-distinct-string-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(KthDistinctStringInAnArray([]string{"d", "b", "c", "b", "c", "a"}, 2)) // "a"
	fmt.Println(KthDistinctStringInAnArray([]string{"aaa", "aa", "a"}, 1))              // "aaa"
	fmt.Println(KthDistinctStringInAnArray([]string{"a", "b", "a"}, 3))                 // ""
}

// Time: O(n), Space: O(n)
func KthDistinctStringInAnArray(arr []string, k int) string {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[string]int)
	for _, s := range arr {
		freq[s]++
	}

	idx := 1
	for _, s := range arr {
		if freq[s] == 1 {
			if idx == k {
				return s
			}
			idx++
		}
	}
	return ""
}
```

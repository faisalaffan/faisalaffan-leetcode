# 3121 — Count The Number Of Special Characters Ii

## Deskripsi

**Soal:** [3121. Count The Number Of Special Characters Ii](https://leetcode.com/problems/count-the-number-of-special-characters-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfSpecialChars(word string) int`

## Solusi Go

```go
package main

// LeetCode #3121: Count the Number of Special Characters II
// https://leetcode.com/problems/count-the-number-of-special-characters-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func numberOfSpecialChars(word string) int {
  // Membuat slice untuk menyimpan hasil
	firstLower := make([]int, 26)
  // Membuat slice untuk menyimpan hasil
	lastUpper := make([]int, 26)
  // Iterasi seluruh elemen
	for i := range firstLower {
		firstLower[i] = -1
		lastUpper[i] = -1
	}

	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			idx := ch - 'a'
			if firstLower[idx] == -1 {
				firstLower[idx] = i
			}
		} else {
			idx := ch - 'A'
			lastUpper[idx] = i
		}
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if firstLower[i] != -1 && lastUpper[i] != -1 && firstLower[i] > lastUpper[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))  // Expected: 3
	fmt.Println(numberOfSpecialChars("abc"))       // Expected: 0
	fmt.Println(numberOfSpecialChars("AbBCab"))    // Expected: 0
}
```

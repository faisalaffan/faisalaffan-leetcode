# 2083 — Substrings That Begin And End With The Same Letter

## Deskripsi

**Soal:** [2083. Substrings That Begin And End With The Same Letter](https://leetcode.com/problems/substrings-that-begin-and-end-with-the-same-letter/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfSubstrings(s string) int64`

## Solusi Go

```go
package main

// LeetCode #2083: Substrings That Begin and End With the Same Letter
// https://leetcode.com/problems/substrings-that-begin-and-end-with-the-same-letter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfSubstrings(s string) int64 {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int64, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	var result int64 = 0
	for _, f := range freq {
		result += f * (f + 1) / 2
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfSubstrings("abc"))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", numberOfSubstrings("abacaba"))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", numberOfSubstrings("aa"))
	// Expected: 3
}
```

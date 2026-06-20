# 3639 — Minimum Time To Activate String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTimeToActivateString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3639: Minimum Time to Activate String
// https://leetcode.com/problems/minimum-time-to-activate-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func minimumTimeToActivateString(s string) int {
	n := len(s)
	// dp[i] = min time to activate prefix of length i
  // Alokasi slice integer
	dp := make([]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = i // worst case: type each character
	}

	for i := 1; i <= n; i++ {
		// Option 1: type the current character (1 sec)
		if dp[i] > dp[i-1]+1 {
			dp[i] = dp[i-1] + 1
		}

		// Option 2: try to activate from a matching prefix
		// Look for a substring in already activated prefix that matches s[i-1:]
		for j := 1; j < i; j++ {
			k := 0
			for i-1+k < n && j-1+k < i-1 && s[i-1+k] == s[j-1+k] {
				k++
			}
			if k > 0 {
				cost := dp[i-1] + 1 // 1 second to activate the matching substring
				if dp[i+k-1] > cost {
					dp[i+k-1] = cost
				}
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(minimumTimeToActivateString("abcabc"))
	fmt.Println(minimumTimeToActivateString("aaaa"))
	fmt.Println(minimumTimeToActivateString("abacaba"))
}
```

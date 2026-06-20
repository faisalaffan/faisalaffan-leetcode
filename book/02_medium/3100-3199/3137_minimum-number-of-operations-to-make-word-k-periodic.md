# 3137 — Minimum Number Of Operations To Make Word K Periodic

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperationsToMakeKPeriodic(word string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n / k)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3137: Minimum Number of Operations to Make Word K-Periodic
// https://leetcode.com/problems/minimum-number-of-operations-to-make-word-k-periodic/
// Difficulty: Medium
// Time: O(n) | Space: O(n / k)

import "fmt"

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	maxFreq := 0

	for i := 0; i < n; i += k {
		sub := word[i : i+k]
		freq[sub]++
		if freq[sub] > maxFreq {
			maxFreq = freq[sub]
		}
	}

	return n/k - maxFreq
}

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcodeleet", 4)) // Expected: 1
	fmt.Println(minimumOperationsToMakeKPeriodic("abcabcabc", 3))    // Expected: 0
	fmt.Println(minimumOperationsToMakeKPeriodic("aabbccddee", 5))   // Expected: 1
}
```

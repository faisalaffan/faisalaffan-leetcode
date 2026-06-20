# 3085 — Minimum Deletions To Make String K Special

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumDeletions(word string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * 26) = O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3085: Minimum Deletions to Make String K-Special
// https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/
// Difficulty: Medium
// Time: O(n * 26) = O(n) | Space: O(26) = O(1)

import "fmt"

func minimumDeletions(word string, k int) int {
  // Alokasi slice integer
	freq := make([]int, 26)
	for _, ch := range word {
		freq[ch-'a']++
	}

	ans := len(word)
	for _, minFreq := range freq {
		if minFreq == 0 {
			continue
		}
		ops := 0
		for _, f := range freq {
			if f < minFreq {
				ops += f
			} else if f > minFreq+k {
				ops += f - (minFreq + k)
			}
		}
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumDeletions("aabcaba", 2)) // Expected: 2
	fmt.Println(minimumDeletions("dabdcbdcdcd", 2)) // Expected: 2
	fmt.Println(minimumDeletions("aaabaaa", 2)) // Expected: 0
}
```

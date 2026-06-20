# 0744 — Find Smallest Letter Greater Than Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func nextGreatestLetter(letters []byte, target byte) byte`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #744: Find Smallest Letter Greater Than Target
// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))) // 'c'
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))) // 'f'
	fmt.Println(string(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))) // 'x'
}

// nextGreatestLetter finds the smallest letter in letters that is greater than target.
// Time: O(log n). Space: O(1).
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[l%len(letters)]
}
```

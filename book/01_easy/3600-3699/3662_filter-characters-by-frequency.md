# 3662 — Filter Characters By Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FilterCharactersByFrequency(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3662: Filter Characters by Frequency
// https://leetcode.com/problems/filter-characters-by-frequency/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(FilterCharactersByFrequency("aadbbcccca", 3))
	fmt.Println(FilterCharactersByFrequency("xyz", 2))
}

// Time: O(n)
// Space: O(1)
func FilterCharactersByFrequency(s string, k int) string {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	res := make([]byte, 0, len(s))
	for _, ch := range s {
		if cnt[ch-'a'] < k {
			res = append(res, byte(ch))
		}
	}
	return string(res)
}
```

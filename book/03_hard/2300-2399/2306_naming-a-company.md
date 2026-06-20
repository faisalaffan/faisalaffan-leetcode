# 2306 — Naming A Company

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func distinctNames(ideas []string) int64
```

> **💡 Hint:** Group ideas by their first character. For each group, store the

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2306: Naming a Company
// https://leetcode.com/problems/naming-a-company/
// Difficulty: Hard
//
// Approach: Group ideas by their first character. For each group, store the
// suffix (everything after first char) in a set. For each pair of groups (i,j)
// with i != j, count common suffixes. Valid pairs from this group pair =
// (len(group[i]) - common) * (len(group[j]) - common). Sum over all pairs.
// Result = sum * 2 (since swap order also counts: "A B" and "B A" are both valid).

import (
	"fmt"
)

func main() {
	// Example 1: ["coffee","donuts","time","toffee"] => 6
	fmt.Println(distinctNames([]string{"coffee", "donuts", "time", "toffee"}))
	// Example 2: ["lack","back"] => 0
	fmt.Println(distinctNames([]string{"lack", "back"}))
	// Edge: single idea
	fmt.Println(distinctNames([]string{"hello"}))
	// Edge: no common suffixes
	fmt.Println(distinctNames([]string{"abc", "def", "ghi"}))
	// Edge: all same first letter
	fmt.Println(distinctNames([]string{"aa", "ab", "ac"}))
}

func distinctNames(ideas []string) int64 {
	// group[first_letter] = set of suffixes
	groups := make([]map[string]bool, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range groups {
		groups[i] = make(map[string]bool)
	}

	for _, idea := range ideas {
		first := idea[0] - 'a'
		suffix := idea[1:]
		groups[first][suffix] = true
	}

	var ans int64
	for i := 0; i < 26; i++ {
		if len(groups[i]) == 0 {
			continue
		}
		for j := i + 1; j < 26; j++ {
			if len(groups[j]) == 0 {
				continue
			}
			// Count common suffixes between groups i and j
			common := 0
			for suffix := range groups[i] {
				if groups[j][suffix] {
					common++
				}
			}
			validI := len(groups[i]) - common
			validJ := len(groups[j]) - common
			ans += int64(validI * validJ * 2)
		}
	}
	return ans
}
```

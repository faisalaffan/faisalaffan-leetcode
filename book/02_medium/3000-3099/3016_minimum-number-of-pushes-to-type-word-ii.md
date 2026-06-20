# 3016 — Minimum Number Of Pushes To Type Word Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumPushes(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3016: Minimum Number of Pushes to Type Word II
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumPushes("abcde"))
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
}

func minimumPushes(word string) int {
  // Alokasi slice integer
	cnt := make([]int, 26)
	for _, ch := range word {
		cnt[ch-'a']++
	}
  // Custom sort dengan comparator
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	ans := 0
	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return ans
}
```

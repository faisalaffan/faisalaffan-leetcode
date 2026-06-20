# 2007 — Find Original Array From Doubled Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindOriginalArrayFromDoubledArray(changed []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2007: Find Original Array From Doubled Array
// https://leetcode.com/problems/find-original-array-from-doubled-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{1, 3, 4, 2, 6, 8}))
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{6, 3, 0, 1}))
	fmt.Println(FindOriginalArrayFromDoubledArray([]int{1}))
}

// Time: O(n log n), Space: O(n)
func FindOriginalArrayFromDoubledArray(changed []int) []int {
	n := len(changed)
	if n%2 == 1 {
		return []int{}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(changed)
	maxVal := changed[n-1]
  // Alokasi slice integer
	cnt := make([]int, maxVal+1)
	for _, x := range changed {
		cnt[x]++
	}

  // Alokasi slice integer
	ans := make([]int, 0, n/2)
	for _, x := range changed {
		if cnt[x] == 0 {
			continue
		}
		if x*2 > maxVal || cnt[x*2] == 0 {
			return []int{}
		}
		ans = append(ans, x)
		cnt[x]--
		cnt[x*2]--
	}

	if len(ans) != n/2 {
		return []int{}
	}
	return ans
}
```

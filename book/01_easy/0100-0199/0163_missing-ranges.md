# 0163 — Missing Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMissingRanges(nums []int, lower int, upper int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #163: Missing Ranges
// https://leetcode.com/problems/missing-ranges/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"strconv"
)

// Time: O(n) | Space: O(1) excluding output
func FindMissingRanges(nums []int, lower int, upper int) []string {
	var res []string
	addRange := func(lo, hi int) {
		if lo > hi {
			return
		}
		if lo == hi {
			res = append(res, strconv.Itoa(lo))
		} else {
			res = append(res, strconv.Itoa(lo)+"->"+strconv.Itoa(hi))
		}
	}
	prev := lower - 1
	for i := 0; i <= len(nums); i++ {
		var curr int
		if i < len(nums) {
			curr = nums[i]
		} else {
			curr = upper + 1
		}
		if curr-prev > 1 {
			addRange(prev+1, curr-1)
		}
		prev = curr
	}
	return res
}

func main() {
	fmt.Println(FindMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99))
	fmt.Println(FindMissingRanges([]int{-1}, -1, -1))
}
```

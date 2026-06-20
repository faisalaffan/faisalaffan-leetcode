# 1985 — Find The Kth Largest Integer In The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheKthLargestIntegerInTheArray(nums []string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) (ignoring sort space)  
**Kompleksitas Ruang:** O(1) (ignoring sort space)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1985: Find the Kth Largest Integer in the Array
// https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"3", "6", "7", "10"}, 4))
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"2", "21", "12", "1"}, 3))
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"0", "0"}, 2))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func FindTheKthLargestIntegerInTheArray(nums []string, k int) string {
  // Custom sort dengan comparator
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) != len(nums[j]) {
			return len(nums[i]) > len(nums[j])
		}
		return nums[i] > nums[j]
	})
	return nums[k-1]
}
```

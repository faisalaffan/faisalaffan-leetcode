# 3101 — Count Alternating Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countAlternatingSubarrays(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3101: Count Alternating Subarrays
// https://leetcode.com/problems/count-alternating-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countAlternatingSubarrays(nums []int) int64 {
	var ans int64
	n := len(nums)
	left := 0

	for right := 0; right < n; right++ {
		if right > 0 && nums[right] == nums[right-1] {
			left = right
		}
		ans += int64(right - left + 1)
	}

	return ans
}

func main() {
	fmt.Println(countAlternatingSubarrays([]int{0, 1, 1, 1})) // Expected: 5
	fmt.Println(countAlternatingSubarrays([]int{1, 0, 1, 0})) // Expected: 10
}
```

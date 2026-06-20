# 2393 — Count Strictly Increasing Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countIncreasing(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2393: Count Strictly Increasing Subarrays
// https://leetcode.com/problems/count-strictly-increasing-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Track length of current increasing run. Each position adds run_len subarrays.

import "fmt"

func main() {
	fmt.Println(countIncreasing([]int{1, 3, 5, 4, 4, 6})) // 10
	fmt.Println(countIncreasing([]int{1, 2, 3, 4, 5}))    // 15
}

func countIncreasing(nums []int) int64 {
	var ans int64
	run := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			run++
		} else {
			run = 1
		}
		ans += int64(run)
	}
	return ans
}
```

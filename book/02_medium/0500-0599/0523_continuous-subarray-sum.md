# 0523 — Continuous Subarray Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckSubarraySum(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k) where k = min(k, n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #523: Continuous Subarray Sum
// https://leetcode.com/problems/continuous-subarray-sum/
// Difficulty: Medium
// Time: O(n)
// Space: O(k) where k = min(k, n)

import "fmt"

func main() {
	fmt.Println(CheckSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(CheckSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	fmt.Println(CheckSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}

func CheckSubarraySum(nums []int, k int) bool {
	// Map remainder -> first index
  // Membuat map (HashMap) — pencarian O(1)
	remainderMap := make(map[int]int)
	remainderMap[0] = -1
	sum := 0

	for i, num := range nums {
		sum += num
		rem := sum % k
		if rem < 0 {
			rem += k
		}
		if prevIdx, ok := remainderMap[rem]; ok {
			if i-prevIdx >= 2 {
				return true
			}
		} else {
			remainderMap[rem] = i
		}
	}

	return false
}
```

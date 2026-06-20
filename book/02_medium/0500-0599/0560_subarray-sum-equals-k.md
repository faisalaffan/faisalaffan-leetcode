# 0560 — Subarray Sum Equals K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SubarraySum(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #560: Subarray Sum Equals K
// https://leetcode.com/problems/subarray-sum-equals-k/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(SubarraySum([]int{1, 1, 1}, 2))
	fmt.Println(SubarraySum([]int{1, 2, 3}, 3))
}

func SubarraySum(nums []int, k int) int {
	count := 0
	sum := 0
  // Membuat map (HashMap) — pencarian O(1)
	sumMap := make(map[int]int)
	sumMap[0] = 1

	for _, num := range nums {
		sum += num
		if val, ok := sumMap[sum-k]; ok {
			count += val
		}
		sumMap[sum]++
	}

	return count
}
```

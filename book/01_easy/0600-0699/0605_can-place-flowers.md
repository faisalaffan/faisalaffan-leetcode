# 0605 — Can Place Flowers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CanPlaceFlowers(flowerbed []int, n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #605: Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(flowerbed) && count < n; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func main() {
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(CanPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))
}
```

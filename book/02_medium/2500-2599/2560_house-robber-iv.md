# 2560 — House Robber Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCapability(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log max)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2560: House Robber IV
// https://leetcode.com/problems/house-robber-iv/
// Difficulty: Medium
// Time: O(n log max) | Space: O(1)

import "fmt"

func minCapability(nums []int, k int) int {
	canRob := func(cap int) bool {
		count := 0
		i := 0
		for i < len(nums) {
			if nums[i] <= cap {
				count++
				i += 2 // Skip adjacent house
			} else {
				i++
			}
		}
		return count >= k
	}

	left, right := nums[0], nums[0]
	for _, v := range nums {
		if v < left {
			left = v
		}
		if v > right {
			right = v
		}
	}

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if canRob(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCapability([]int{2, 3, 5, 9}, 2))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", minCapability([]int{2, 7, 9, 3, 1}, 2))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", minCapability([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	// Expected: 5
}
```

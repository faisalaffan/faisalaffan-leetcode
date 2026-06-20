# 0042 — Trapping Rain Water

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func trap(height []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #42: Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/
// Difficulty: Hard

import "fmt"

// trap calculates the total amount of water that can be trapped between bars.
// Uses the two-pointer approach.
//
// Complexity: O(n) time, O(1) space
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	total := 0

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}

	return total
}

func main() {
	// Test case from LeetCode
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("Test 1:", trap(height)) // 6

	// Additional test cases
	fmt.Println("Test 2: [4,2,0,3,2,5] ->", trap([]int{4, 2, 0, 3, 2, 5})) // 9
	fmt.Println("Test 3: [1,2,3] ->", trap([]int{1, 2, 3}))                // 0
	fmt.Println("Test 4: [] ->", trap([]int{}))                            // 0
}
```

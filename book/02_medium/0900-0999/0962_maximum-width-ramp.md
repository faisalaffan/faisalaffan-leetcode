# 0962 — Maximum Width Ramp

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxWidthRamp(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #962: Maximum Width Ramp
// https://leetcode.com/problems/maximum-width-ramp/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(n)
func maxWidthRamp(nums []int) int {
	n := len(nums)
  // Alokasi slice integer
	stack := make([]int, 0)

	// Build decreasing stack of indices
	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	ans := 0
	for j := n - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
			width := j - stack[len(stack)-1]
			if width > ans {
				ans = width
			}
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}
```

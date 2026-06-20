# 0503 — Next Greater Element Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NextGreaterElementIi(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #503: Next Greater Element II
// https://leetcode.com/problems/next-greater-element-ii/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(NextGreaterElementIi([]int{1, 2, 1}))
	fmt.Println(NextGreaterElementIi([]int{1, 2, 3, 4, 3}))
}

func NextGreaterElementIi(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	result := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = -1
	}

	stack := []int{}
	// Iterate twice to handle circular array
	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = num
		}
		if i < n {
			stack = append(stack, i)
		}
	}

	return result
}
```

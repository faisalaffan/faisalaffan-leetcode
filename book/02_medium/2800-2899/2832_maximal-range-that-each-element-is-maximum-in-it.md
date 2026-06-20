# 2832 — Maximal Range That Each Element Is Maximum In It

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximalRangeThatEachElementIsMaximumInIt(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2832: Maximal Range That Each Element Is Maximum in It
// https://leetcode.com/problems/maximal-range-that-each-element-is-maximum-in-it/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximalRangeThatEachElementIsMaximumInIt(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	result := make([]int, n)

	// Previous greater element
  // Alokasi slice integer
	prev := make([]int, n)
  // Alokasi slice integer
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prev[i] = stack[len(stack)-1]
		} else {
			prev[i] = -1
		}
		stack = append(stack, i)
	}

	// Next greater element
  // Alokasi slice integer
	next := make([]int, n)
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			next[i] = stack[len(stack)-1]
		} else {
			next[i] = n
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		result[i] = next[i] - prev[i] - 1
	}

	return result
}

func main() {
	fmt.Println(MaximalRangeThatEachElementIsMaximumInIt([]int{1, 5, 4, 3, 6}))
	fmt.Println(MaximalRangeThatEachElementIsMaximumInIt([]int{1, 2, 1}))
}
```

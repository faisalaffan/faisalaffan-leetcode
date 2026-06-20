# 3834 — Merge Adjacent Equal Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MergeAdjacentEqualElements(nums []int) []int64
```

> **💡 Hint:** Use a stack to repeatedly merge leftmost adjacent equal pairs.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3834: Merge Adjacent Equal Elements
// https://leetcode.com/problems/merge-adjacent-equal-elements/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Use a stack to repeatedly merge leftmost adjacent equal pairs.

import "fmt"

func MergeAdjacentEqualElements(nums []int) []int64 {
  // Alokasi slice integer
	stack := make([]int64, 0)

	for _, v := range nums {
		cur := int64(v)
		// While stack top equals cur, merge (pop + double)
		for len(stack) > 0 && stack[len(stack)-1] == cur {
			cur += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, cur)
	}

	return stack
}

func main() {
	// Example 1
	fmt.Println(MergeAdjacentEqualElements([]int{3, 1, 1, 2})) // Expected: [3 4]

	// Example 2
	fmt.Println(MergeAdjacentEqualElements([]int{2, 2, 4})) // Expected: [8]

	// Example 3
	fmt.Println(MergeAdjacentEqualElements([]int{3, 7, 5})) // Expected: [3 7 5]
}
```

# 1504 — Count Submatrices With All Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumSubmat(mat [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** O(R*C), Space: O(C)  
**Kompleksitas Ruang:** O(C)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1504: Count Submatrices With All Ones
// https://leetcode.com/problems/count-submatrices-with-all-ones/
// Difficulty: Medium

import "fmt"

func main() {
	mat1 := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(NumSubmat(mat1))

	mat2 := [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}
	fmt.Println(NumSubmat(mat2))

	fmt.Println(NumSubmat([][]int{{1, 1}, {1, 1}}))
}

func NumSubmat(mat [][]int) int {
	// Time: O(R*C), Space: O(C)
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	rows, cols := len(mat), len(mat[0])
  // Alokasi slice integer
	height := make([]int, cols)
	total := 0

	for r := 0; r < rows; r++ {
		// Update height of consecutive 1s in each column
		for c := 0; c < cols; c++ {
			if mat[r][c] == 1 {
				height[c]++
			} else {
				height[c] = 0
			}
		}

		// Count submatrices ending at row r using monotonic stack
  // Alokasi slice integer
		stack := make([]int, 0)
  // Alokasi slice integer
		sum := make([]int, cols)

		for c := 0; c < cols; c++ {
			// Pop from stack while height[c] <= height[stack top]
			for len(stack) > 0 && height[stack[len(stack)-1]] >= height[c] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				prev := stack[len(stack)-1]
				sum[c] = sum[prev] + height[c]*(c-prev)
			} else {
				sum[c] = height[c] * (c + 1)
			}

			stack = append(stack, c)
			total += sum[c]
		}
	}

	return total
}
```

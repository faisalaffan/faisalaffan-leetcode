# 0946 — Validate Stack Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func validateStackSequences(pushed []int, popped []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #946: Validate Stack Sequences
// https://leetcode.com/problems/validate-stack-sequences/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(n)
func validateStackSequences(pushed []int, popped []int) bool {
  // Alokasi slice integer
	stack := make([]int, 0)
	j := 0
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) > 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
```

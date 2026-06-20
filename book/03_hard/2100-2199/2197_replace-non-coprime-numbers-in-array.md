# 2197 — Replace Non Coprime Numbers In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lcm(a, b int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack, GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2197: Replace Non-Coprime Numbers in Array
// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/
// Difficulty: Hard
//
// Stack merge: for each number, push onto stack. While stack has >= 2 elements
// and gcd(stack[-2], stack[-1]) > 1, pop top two, push lcm.

import (
	"fmt"
)

func main() {
	// Example: [6,4,3,2,7,6,2] => [12,7,6]
	fmt.Println(replaceNonCoprimeNumbers([]int{6, 4, 3, 2, 7, 6, 2}))
	// Example: [2,2,1,1,3,3,3] => [2,1,1,3]
	fmt.Println(replaceNonCoprimeNumbers([]int{2, 2, 1, 1, 3, 3, 3}))
	// Example: [1,1,1,1] => [1,1,1,1]
	fmt.Println(replaceNonCoprimeNumbers([]int{1, 1, 1, 1}))
	// Example: [12,18,6] => [36]
	fmt.Println(replaceNonCoprimeNumbers([]int{12, 18, 6}))
	// Example: single
	fmt.Println(replaceNonCoprimeNumbers([]int{7}))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func replaceNonCoprimeNumbers(nums []int) []int {
  // Alokasi slice integer
	stack := make([]int, 0, len(nums))

	for _, x := range nums {
		stack = append(stack, x)
		for len(stack) >= 2 {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			g := gcd(a, b)
			if g == 1 {
				break
			}
			// Replace with LCM
			stack = stack[:len(stack)-2]
			stack = append(stack, lcm(a, b))
		}
	}

	return stack
}
```

# 1052 — Grumpy Bookstore Owner

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSatisfied(customers []int, grumpy []int, minutes int) int
```

> **💡 Hint:** Sliding window - find best minutes to use secret technique

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1052: Grumpy Bookstore Owner
// https://leetcode.com/problems/grumpy-bookstore-owner/
// Difficulty: Medium
//
// Approach: Sliding window - find best minutes to use secret technique
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)) // 16
	fmt.Println(maxSatisfied([]int{1}, []int{0}, 1)) // 1
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	baseSatisfied := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			baseSatisfied += customers[i]
		}
	}

	extraSatisfied := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			extraSatisfied += customers[i]
		}
	}

	maxExtra := extraSatisfied
	for i := minutes; i < n; i++ {
		if grumpy[i-minutes] == 1 {
			extraSatisfied -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			extraSatisfied += customers[i]
		}
		if extraSatisfied > maxExtra {
			maxExtra = extraSatisfied
		}
	}

	return baseSatisfied + maxExtra
}
```

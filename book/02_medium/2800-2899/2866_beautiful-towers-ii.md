# 2866 — Beautiful Towers Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BeautifulTowersIi(maxHeights []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2866: Beautiful Towers II
// https://leetcode.com/problems/beautiful-towers-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func BeautifulTowersIi(maxHeights []int) int64 {
	n := len(maxHeights)

	// Left to right: sum of heights ending at i with non-decreasing left side
  // Alokasi slice integer
	left := make([]int64, n)
  // Alokasi slice integer
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && maxHeights[stack[len(stack)-1]] > maxHeights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = int64(maxHeights[i]) * int64(i+1)
		} else {
			prev := stack[len(stack)-1]
			left[i] = left[prev] + int64(maxHeights[i])*int64(i-prev)
		}
		stack = append(stack, i)
	}

	// Right to left
  // Alokasi slice integer
	right := make([]int64, n)
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && maxHeights[stack[len(stack)-1]] > maxHeights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = int64(maxHeights[i]) * int64(n-i)
		} else {
			prev := stack[len(stack)-1]
			right[i] = right[prev] + int64(maxHeights[i])*int64(prev-i)
		}
		stack = append(stack, i)
	}

	var best int64
	for i := 0; i < n; i++ {
		total := left[i] + right[i] - int64(maxHeights[i])
		if total > best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(BeautifulTowersIi([]int{5, 3, 4, 1, 1}))
	fmt.Println(BeautifulTowersIi([]int{6, 5, 3, 9, 2, 7}))
}
```

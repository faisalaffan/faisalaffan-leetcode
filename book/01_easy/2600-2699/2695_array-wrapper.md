# 2695 — Array Wrapper

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArrayWrapper(nums []int) *ArrayWrapperVal
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2695: Array Wrapper
// https://leetcode.com/problems/array-wrapper/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Wraps an array with string conversion and addition.

import (
	"fmt"
	"strings"
)

func main() {
	w1 := ArrayWrapper([]int{1, 2})
	w2 := ArrayWrapper([]int{3, 4})
	fmt.Println(ArrayWrapperAdd(w1, w2))
	fmt.Println(ArrayWrapperString(w1))
}

type ArrayWrapperVal struct {
	nums []int
}

func ArrayWrapper(nums []int) *ArrayWrapperVal {
	return &ArrayWrapperVal{nums: nums}
}

func ArrayWrapperAdd(a, b *ArrayWrapperVal) int {
	sum := 0
	for _, v := range a.nums {
		sum += v
	}
	for _, v := range b.nums {
		sum += v
	}
	return sum
}

func ArrayWrapperString(w *ArrayWrapperVal) string {
	strs := make([]string, len(w.nums))
	for i, v := range w.nums {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return "[" + strings.Join(strs, ",") + "]"
}
```

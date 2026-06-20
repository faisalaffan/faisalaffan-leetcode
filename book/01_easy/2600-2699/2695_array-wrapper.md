# 2695 — Array Wrapper

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ArrayWrapper(nums []int) *ArrayWrapperVal`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


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

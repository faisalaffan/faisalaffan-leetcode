# 0702 — Search In A Sorted Array Of Unknown Size

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func search(reader *ArrayReader, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** O(log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #702: Search in a Sorted Array of Unknown Size
// https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/
// Difficulty: Medium [Paid]
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	reader := &ArrayReader{arr: []int{-1, 0, 3, 5, 9, 12}}
	fmt.Println(search(reader, 9))
	fmt.Println(search(reader, 2))
}

type ArrayReader struct {
	arr []int
}

func (r *ArrayReader) get(index int) int {
	if index >= len(r.arr) {
		return 1 << 31 - 1
	}
	return r.arr[index]
}

func search(reader *ArrayReader, target int) int {
	// Find upper bound
	left, right := 0, 1
	for reader.get(right) < target {
		left = right
		right <<= 1
	}

	// Binary search
  // Binary search loop
	for left <= right {
		mid := left + (right-left)/2
		val := reader.get(mid)
		if val == target {
			return mid
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
```

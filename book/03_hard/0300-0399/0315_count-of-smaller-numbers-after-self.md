# 0315 — Count Of Smaller Numbers After Self

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countSmaller(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #315: Count of Smaller Numbers After Self
// https://leetcode.com/problems/count-of-smaller-numbers-after-self/
// Difficulty: Hard

import "fmt"

func countSmaller(nums []int) []int {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return []int{}
	}

	// Pair each number with its original index
	type pair struct {
		val int
		idx int
	}
	arr := make([]pair, n)
	for i, v := range nums {
		arr[i] = pair{val: v, idx: i}
	}

  // Alokasi slice integer
	result := make([]int, n)

	var mergeSort func([]pair) []pair
	mergeSort = func(a []pair) []pair {
		if len(a) <= 1 {
			return a
		}
		mid := len(a) / 2
		left := mergeSort(a[:mid])
		right := mergeSort(a[mid:])

		// Merge while counting
		merged := make([]pair, 0, len(a))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if left[i].val <= right[j].val {
				// All elements already placed from right that are smaller
				result[left[i].idx] += j
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[j])
				j++
			}
		}
		for i < len(left) {
			result[left[i].idx] += j
			merged = append(merged, left[i])
			i++
		}
		for j < len(right) {
			merged = append(merged, right[j])
			j++
		}
		return merged
	}

	mergeSort(arr)
	return result
}

func main() {
	// Example 1
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
	// [2, 1, 1, 0]

	// Example 2
	fmt.Println(countSmaller([]int{-1}))
	// [0]

	// Example 3
	fmt.Println(countSmaller([]int{-1, -1}))
	// [0, 0]
}
```

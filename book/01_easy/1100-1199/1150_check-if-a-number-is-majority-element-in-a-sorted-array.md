# 1150 — Check If A Number Is Majority Element In A Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isMajorityElement(nums []int, target int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1150: Check If a Number Is Majority Element in a Sorted Array
// https://leetcode.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isMajorityElement([]int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5)) // true
	fmt.Println(isMajorityElement([]int{10, 100, 101, 101}, 101))        // false
}

// LeetCode submission: isMajorityElement
func isMajorityElement(nums []int, target int) bool {
	n := len(nums)
	// Binary search for first occurrence
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	first := lo
	last := first + n/2
	return last < n && nums[last] == target
}
```

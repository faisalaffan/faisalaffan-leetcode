# 1063 — Number Of Valid Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func validSubarrays(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1063: Number of Valid Subarrays
// https://leetcode.com/problems/number-of-valid-subarrays/
// Difficulty: Hard [Paid]
//
// Monotonic stack approach. A valid subarray is one where the first element
// is the minimum of that subarray. For each element at index i, we find the
// next smaller element to the right (at index j). All subarrays starting at
// i and ending before j have arr[i] as the minimum, so count += j-i.

import "fmt"

func main() {
	fmt.Println(validSubarrays([]int{1, 4, 2, 5, 3}))
}

func validSubarrays(nums []int) int {
	n := len(nums)
  // Alokasi slice integer
	stack := make([]int, 0)
	count := 0

	for i := 0; i <= n; i++ {
		// Pop elements while current value is smaller than top of stack
		for len(stack) > 0 && (i == n || nums[stack[len(stack)-1]] > nums[i]) {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			count += i - idx
		}
		if i < n {
			stack = append(stack, i)
		}
	}

	return count
}
```

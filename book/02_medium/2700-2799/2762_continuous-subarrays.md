# 2762 — Continuous Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ContinuousSubarrays(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2762: Continuous Subarrays
// https://leetcode.com/problems/continuous-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ContinuousSubarrays(nums []int) int64 {
	n := len(nums)
	var result int64
	left := 0

	// Track min and max using deques via slice
	// minDeque stores indices with increasing values
	// maxDeque stores indices with decreasing values
  // Alokasi slice integer
	minDeque := make([]int, 0)
  // Alokasi slice integer
	maxDeque := make([]int, 0)

	for right := 0; right < n; right++ {
		// Maintain minDeque
		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] > nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, right)

		// Maintain maxDeque
		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] < nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, right)

		// Shrink window if condition violated
		for nums[maxDeque[0]]-nums[minDeque[0]] > 2 {
			if minDeque[0] == left {
				minDeque = minDeque[1:]
			}
			if maxDeque[0] == left {
				maxDeque = maxDeque[1:]
			}
			left++
		}

		result += int64(right - left + 1)
	}

	return result
}

func main() {
	fmt.Println(ContinuousSubarrays([]int{5, 4, 2, 4}))
	fmt.Println(ContinuousSubarrays([]int{1, 2, 3}))
}
```

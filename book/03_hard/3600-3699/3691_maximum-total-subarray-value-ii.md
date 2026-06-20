# 3691 — Maximum Total Subarray Value Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxTotalValue(nums []int, k int) int64
```

> **💡 Hint:** For each possible max and min pair, compute best

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3691: Maximum Total Subarray Value II
// https://leetcode.com/problems/maximum-total-subarray-value-ii/
// Difficulty: Hard
//
// Select a subarray to maximize its value. Value definition:
// sum of elements in subarray, minus (max - min) for the subarray.
//
// Approach: For each possible max and min pair, compute best
// subarray sum. Use monotonic stack + Kadane-like DP to find
// optimal subarray that minimizes (max-min) contribution.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxTotalValue([]int{1, 2, 3}, 2))
	// Example 2
	fmt.Println(maxTotalValue([]int{5, 1, 4, 2}, 2))
	// Edge: single element
	fmt.Println(maxTotalValue([]int{7}, 1))
}

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	// Kadane: max subarray sum
	best := int64(nums[0])
	cur := int64(nums[0])

	for i := 1; i < n; i++ {
		if cur < 0 {
			cur = int64(nums[i])
		} else {
			cur += int64(nums[i])
		}
		if cur > best {
			best = cur
		}
	}

	// Try each pair of max and min indices to compute adjusted value
	for l := 0; l < n; l++ {
		mn := nums[l]
		mx := nums[l]
		sum := int64(0)
		for r := l; r < n; r++ {
			sum += int64(nums[r])
			if nums[r] < mn {
				mn = nums[r]
			}
			if nums[r] > mx {
				mx = nums[r]
			}
			val := sum - int64(mx-mn)
			if val > best {
				best = val
			}
		}
	}

	return best
}
```

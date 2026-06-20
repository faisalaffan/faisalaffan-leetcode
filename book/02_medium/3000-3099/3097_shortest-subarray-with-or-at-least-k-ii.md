# 3097 — Shortest Subarray With Or At Least K Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSubarrayLength(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Bitmask

**Kompleksitas Waktu:** O(n * 32) = O(n)  
**Kompleksitas Ruang:** O(32) = O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3097: Shortest Subarray With OR at Least K II
// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/
// Difficulty: Medium
// Time: O(n * 32) = O(n) | Space: O(32) = O(1)

import "fmt"

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}

	n := len(nums)
	ans := n + 1
  // Alokasi slice integer
	bits := make([]int, 32)
	left := 0
	cur := 0

	for right := 0; right < n; right++ {
		cur |= nums[right]
		for b := 0; b < 32; b++ {
			if nums[right]&(1<<b) != 0 {
				bits[b]++
			}
		}

		for left <= right && cur >= k {
			if right-left+1 < ans {
				ans = right - left + 1
			}

			for b := 0; b < 32; b++ {
				if nums[left]&(1<<b) != 0 {
					bits[b]--
					if bits[b] == 0 {
						cur &^= (1 << b)
					}
				}
			}
			left++
		}
	}

	if ans > n {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSubarrayLength([]int{1, 2, 3}, 3))       // Expected: 1
	fmt.Println(minimumSubarrayLength([]int{1, 2, 3}, 5))       // Expected: 2
	fmt.Println(minimumSubarrayLength([]int{2, 1, 8}, 10))      // Expected: 3
}
```

# 2172 — Maximum And Sum Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumANDSum(nums []int, numSlots int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2172: Maximum AND Sum of Array
// https://leetcode.com/problems/maximum-and-sum-of-array/
// Difficulty: Hard
//
// DP with bitmask. Double each slot (capacity 2 -> 2 slots of capacity 1).
// dp[mask] = max AND sum for the given assignment mask.

import "fmt"

func main() {
	fmt.Println(maximumANDSum([]int{1, 2, 3, 4, 5, 6}, 3)) // 9
	fmt.Println(maximumANDSum([]int{1, 3, 10, 4, 7, 1}, 3)) // 10
	fmt.Println(maximumANDSum([]int{1, 2, 3}, 2))            // 5
	fmt.Println(maximumANDSum([]int{1, 2}, 1))               // 1
}

func maximumANDSum(nums []int, numSlots int) int {
	n := len(nums)
	m := 2 * numSlots // doubled slots
	total := 1 << m

  // Alokasi slice integer
	dp := make([]int, total)
	for i := 1; i < total; i++ {
		dp[i] = -1
	}

	ans := 0
	for mask := 0; mask < total; mask++ {
		if dp[mask] < 0 {
			continue
		}
		idx := popcount(mask)
		if idx >= n {
			if dp[mask] > ans {
				ans = dp[mask]
			}
			continue
		}
		for slot := 0; slot < m; slot++ {
			if mask&(1<<slot) == 0 {
				nm := mask | (1 << slot)
				val := dp[mask] + (nums[idx] & (slot/2 + 1))
				if val > dp[nm] {
					dp[nm] = val
				}
			}
		}
	}
	return ans
}

func popcount(x int) int {
	c := 0
	for x > 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func MaximumAndSumOfArray() any {
	return maximumANDSum([]int{1, 2, 3, 4, 5, 6}, 3)
}
```

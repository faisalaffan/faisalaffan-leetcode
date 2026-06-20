# 1524 — Number Of Sub Arrays With Odd Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumOfSubarrays(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1524: Number of Sub-arrays With Odd Sum
// https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumOfSubarrays([]int{1, 3, 5}))
	fmt.Println(NumOfSubarrays([]int{2, 4, 6}))
	fmt.Println(NumOfSubarrays([]int{1, 2, 3, 4, 5, 6, 7}))
}

func NumOfSubarrays(arr []int) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	oddCount := 0
	evenCount := 1 // prefix sum = 0 is even
	prefixSum := 0
	result := 0

	for _, num := range arr {
		prefixSum += num

		if prefixSum%2 == 0 {
			// Current prefix is even
			result = (result + oddCount) % mod
			evenCount++
		} else {
			// Current prefix is odd
			result = (result + evenCount) % mod
			oddCount++
		}
	}

	return result
}
```

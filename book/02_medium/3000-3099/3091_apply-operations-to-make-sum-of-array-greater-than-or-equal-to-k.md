# 3091 — Apply Operations To Make Sum Of Array Greater Than Or Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minOperations(k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(sqrt(k))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3091: Apply Operations to Make Sum of Array Greater Than or Equal to k
// https://leetcode.com/problems/apply-operations-to-make-sum-of-array-greater-than-or-equal-to-k/
// Difficulty: Medium
// Time: O(sqrt(k)) | Space: O(1)

import "fmt"

func minOperations(k int) int {
	ans := k - 1
	for a := 1; a <= k; a++ {
		b := (k + a - 1) / a
		ops := (a - 1) + (b - 1)
		if ops < ans {
			ans = ops
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations(11))  // Expected: 5
	fmt.Println(minOperations(1))   // Expected: 0
	fmt.Println(minOperations(5))   // Expected: 3
}
```

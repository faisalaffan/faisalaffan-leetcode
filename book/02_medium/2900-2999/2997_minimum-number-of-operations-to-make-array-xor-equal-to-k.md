# 2997 — Minimum Number Of Operations To Make Array Xor Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minOperationsXOR(nums []int, k int) (ans int)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2997: Minimum Number of Operations to Make Array XOR Equal to K
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(minOperationsXOR([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperationsXOR([]int{2, 0, 2, 0}, 0))
}

func minOperationsXOR(nums []int, k int) (ans int) {
	xor := 0
	for _, x := range nums {
		xor ^= x
	}
	return bits.OnesCount(uint(xor ^ k))
}
```

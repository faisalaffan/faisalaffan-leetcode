# 3792 — Sum Of Increasing Product Blocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func sumOfIncreasingProductBlocks(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3792: Sum of Increasing Product Blocks
// https://leetcode.com/problems/sum-of-increasing-product-blocks/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

const mod3792 = 1000000007

func sumOfIncreasingProductBlocks(n int) int {
	ans := 0
	k := 1
	for i := 1; i <= n; i++ {
		prod := 1
		for j := k; j < k+i; j++ {
			prod = (prod * j) % mod3792
		}
		ans = (ans + prod) % mod3792
		k += i
	}
	return ans
}

func main() {
	fmt.Println(sumOfIncreasingProductBlocks(3))
	fmt.Println(sumOfIncreasingProductBlocks(7))
	fmt.Println(sumOfIncreasingProductBlocks(1))
}
```

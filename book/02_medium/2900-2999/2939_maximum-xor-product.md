# 2939 — Maximum Xor Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maximumXorProduct(a int64, b int64, n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2939: Maximum Xor Product
// https://leetcode.com/problems/maximum-xor-product/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maximumXorProduct(12, 5, 4))
	fmt.Println(maximumXorProduct(6, 7, 5))
	fmt.Println(maximumXorProduct(1, 6, 3))
}

func maximumXorProduct(a int64, b int64, n int) int {
	const mod int64 = 1e9 + 7
	ax := (a >> n) << n
	bx := (b >> n) << n
	for i := n - 1; i >= 0; i-- {
		x, y := (a>>i)&1, (b>>i)&1
		if x == y {
			ax |= 1 << i
			bx |= 1 << i
		} else if ax < bx {
			ax |= 1 << i
		} else {
			bx |= 1 << i
		}
	}
	ax %= mod
	bx %= mod
	return int(ax * bx % mod)
}
```

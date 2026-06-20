# 3125 — Maximum Number That Makes Result Of Bitwise And Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxNumber(n int64) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3125: Maximum Number That Makes Result of Bitwise AND Zero
// https://leetcode.com/problems/maximum-number-that-makes-result-of-bitwise-and-zero/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"
import "math/bits"

func maxNumber(n int64) int64 {
	if n <= 0 {
		return 0
	}
	bits := bits.Len64(uint64(n))
	return int64((uint64(1) << (bits - 1)) - 1)
}

func main() {
	fmt.Println(maxNumber(5))  // Expected: 3
	fmt.Println(maxNumber(10)) // Expected: 7
	fmt.Println(maxNumber(1))  // Expected: 0
}
```

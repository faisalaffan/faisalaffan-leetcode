# 3133 — Minimum Array End

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minEnd(n int, x int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3133: Minimum Array End
// https://leetcode.com/problems/minimum-array-end/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func minEnd(n int, x int) int64 {
	n64 := int64(n - 1)
	x64 := int64(x)
	ans := int64(0)

	bitPos := 0
	for n64 > 0 || x64 > 0 {
		if x64&1 == 1 {
			ans |= (int64(1) << bitPos)
		} else {
			ans |= ((n64 & 1) << bitPos)
			n64 >>= 1
		}
		x64 >>= 1
		bitPos++
	}
	return ans
}

func main() {
	fmt.Println(minEnd(3, 4))  // Expected: 6
	fmt.Println(minEnd(2, 7))  // Expected: 15
	fmt.Println(minEnd(1, 5))  // Expected: 5
}
```

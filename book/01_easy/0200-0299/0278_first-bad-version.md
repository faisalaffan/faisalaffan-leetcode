# 0278 — First Bad Version

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func isBadVersion(version int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #278: First Bad Version
// https://leetcode.com/problems/first-bad-version/
// Difficulty: Easy

import "fmt"

var firstBad int

func isBadVersion(version int) bool {
	return version >= firstBad
}

// Time: O(log n) | Space: O(1)
func FirstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	firstBad = 4
	fmt.Println(FirstBadVersion(5))
	firstBad = 1
	fmt.Println(FirstBadVersion(1))
}
```

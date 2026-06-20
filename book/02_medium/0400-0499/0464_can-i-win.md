# 0464 — Can I Win

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CanIWin(maxChoosableInteger int, desiredTotal int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(2^n) where n = maxChoosableInteger (max 20)  |  **Ruang:** O(2^n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #464: Can I Win
// https://leetcode.com/problems/can-i-win/
// Difficulty: Medium
// Time: O(2^n) where n = maxChoosableInteger (max 20)
// Space: O(2^n)

import "fmt"

func main() {
	fmt.Println(CanIWin(10, 11))
	fmt.Println(CanIWin(10, 0))
	fmt.Println(CanIWin(10, 40))
}

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}
	sum := maxChoosableInteger * (maxChoosableInteger + 1) / 2
	if sum < desiredTotal {
		return false
	}

  // HashMap: O(1) lookup
	memo := make(map[int]bool)
	var dfs func(used int, currentTotal int) bool
	dfs = func(used int, currentTotal int) bool {
		if currentTotal >= desiredTotal {
			return false
		}
		if val, ok := memo[used]; ok {
			return val
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			mask := 1 << (i - 1)
			if used&mask == 0 {
				if !dfs(used|mask, currentTotal+i) {
					memo[used] = true
					return true
				}
			}
		}
		memo[used] = false
		return false
	}

	return dfs(0, 0)
}
```

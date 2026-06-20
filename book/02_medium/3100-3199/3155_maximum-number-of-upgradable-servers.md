# 3155 — Maximum Number Of Upgradable Servers

## Deskripsi

**Soal:** [3155. Maximum Number Of Upgradable Servers](https://leetcode.com/problems/maximum-number-of-upgradable-servers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func maxUpgrades(count []int, upgrade []int, sell []int, money []int) []int`

## Solusi Go

```go
package main

// LeetCode #3155: Maximum Number of Upgradable Servers
// https://leetcode.com/problems/maximum-number-of-upgradable-servers/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maxUpgrades(count []int, upgrade []int, sell []int, money []int) []int {
	n := len(count)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		// For each server type, binary search max upgrades
		lo, hi := 0, count[i]
		for lo <= hi {
			mid := (lo + hi) / 2
			cost := mid * upgrade[i]
			revenue := (count[i] - mid) * sell[i]
			if revenue+cost <= money[i]+revenue {
				// mid upgrades possible if: mid*upgrade <= money[i] + sold*revenue from non-upgraded
				// Simplify: mid*upgrade[i] <= money[i] + (count[i]-mid)*sell[i]
				if mid*upgrade[i] <= money[i]+(count[i]-mid)*sell[i] {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			} else {
				hi = mid - 1
			}
		}
		ans[i] = hi
	}
	return ans
}

func main() {
	fmt.Println(maxUpgrades([]int{2, 3}, []int{3, 4}, []int{1, 2}, []int{4, 5})) // Expected: [1 1]
	fmt.Println(maxUpgrades([]int{1, 1}, []int{5, 5}, []int{1, 1}, []int{0, 0})) // Expected: [0 0]
}
```

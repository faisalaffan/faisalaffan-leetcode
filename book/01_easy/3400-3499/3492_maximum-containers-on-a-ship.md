# 3492 — Maximum Containers On A Ship

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MaximumContainersOnAShip(n int, w int, maxWeight int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3492: Maximum Containers on a Ship
// https://leetcode.com/problems/maximum-containers-on-a-ship/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumContainersOnAShip(3, 2, 10))
	fmt.Println(MaximumContainersOnAShip(5, 3, 50))
}

// MaximumContainersOnAShip returns the maximum number of containers that can be loaded on an n x n ship, each weighing w, within maxWeight.
// Time: O(1). Space: O(1).
func MaximumContainersOnAShip(n int, w int, maxWeight int) int {
	maxBySpace := n * n
	maxByWeight := maxWeight / w
	if maxBySpace < maxByWeight {
		return maxBySpace
	}
	return maxByWeight
}
```

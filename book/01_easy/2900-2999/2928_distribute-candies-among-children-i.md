# 2928 — Distribute Candies Among Children I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func DistributeCandiesAmongChildrenI(n int, limit int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(limit^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2928: Distribute Candies Among Children I
// https://leetcode.com/problems/distribute-candies-among-children-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: distributeCandies
	fmt.Println(DistributeCandiesAmongChildrenI(5, 2)) // 3
	fmt.Println(DistributeCandiesAmongChildrenI(3, 3)) // 10
}

// Time: O(limit^2) | Space: O(1)
// LeetCode submission name: distributeCandies
func DistributeCandiesAmongChildrenI(n int, limit int) int {
	ways := 0
	for a := 0; a <= limit && a <= n; a++ {
		for b := 0; b <= limit && a+b <= n; b++ {
			c := n - a - b
			if c <= limit {
				ways++
			}
		}
	}
	return ways
}
```

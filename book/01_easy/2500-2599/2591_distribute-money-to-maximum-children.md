# 2591 — Distribute Money To Maximum Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func DistributeMoneyToMaximumChildren(money int, children int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2591: Distribute Money to Maximum Children
// https://leetcode.com/problems/distribute-money-to-maximum-children/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DistributeMoneyToMaximumChildren(20, 3)) // 1
	fmt.Println(DistributeMoneyToMaximumChildren(16, 2)) // 2
}

func DistributeMoneyToMaximumChildren(money int, children int) int {
	if money < children {
		return -1
	}

	// Give each child 1 dollar first
	money -= children

	// Now we have 7-dollar increments (to make 8) for as many children as possible
	count := money / 7
	money %= 7

	// If count > children, we over-assigned
	if count > children {
		return children - 1
	}

	// If we have 3 children left and money = 3, we can't give it optimally
	remaining := children - count
	if remaining == 0 && money > 0 {
		count--
	} else if remaining == 1 && money == 3 {
		count--
	}

	return count
}
```

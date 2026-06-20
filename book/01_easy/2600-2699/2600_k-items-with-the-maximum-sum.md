# 2600 — K Items With The Maximum Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func KItemsWithTheMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2600: K Items With the Maximum Sum
// https://leetcode.com/problems/k-items-with-the-maximum-sum/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(KItemsWithTheMaximumSum(3, 2, 0, 2)) // 2
	fmt.Println(KItemsWithTheMaximumSum(3, 2, 0, 4)) // 3
}

func KItemsWithTheMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	}
	k -= numOnes
	if k <= numZeros {
		return numOnes
	}
	k -= numZeros
	return numOnes - k
}
```

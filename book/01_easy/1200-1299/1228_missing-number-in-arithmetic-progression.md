# 1228 — Missing Number In Arithmetic Progression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func missingNumber(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1228: Missing Number In Arithmetic Progression
// https://leetcode.com/problems/missing-number-in-arithmetic-progression/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{5, 7, 11, 13}))  // 9
	fmt.Println(missingNumber([]int{15, 13, 12}))    // 14
}

// LeetCode submission: missingNumber
func missingNumber(arr []int) int {
	n := len(arr)
	diff := (arr[n-1] - arr[0]) / n
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] == arr[0]+mid*diff {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[0] + diff*lo
}
```

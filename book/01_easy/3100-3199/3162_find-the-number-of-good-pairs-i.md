# 3162 — Find The Number Of Good Pairs I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindTheNumberOfGoodPairsI(nums1 []int, nums2 []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3162: Find the Number of Good Pairs I
// https://leetcode.com/problems/find-the-number-of-good-pairs-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfPairs
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // 2
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 3}, []int{1, 2, 3}, 1))  // 5
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: numberOfPairs
func FindTheNumberOfGoodPairsI(nums1 []int, nums2 []int, k int) int {
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				count++
			}
		}
	}
	return count
}
```

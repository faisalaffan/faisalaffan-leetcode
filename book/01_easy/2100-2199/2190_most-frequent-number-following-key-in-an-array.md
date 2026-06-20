# 2190 — Most Frequent Number Following Key In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MostFrequentNumberFollowingKeyInAnArray(nums []int, key int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2190: Most Frequent Number Following Key In an Array
// https://leetcode.com/problems/most-frequent-number-following-key-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{1, 100, 200, 1, 100}, 1))    // 100
	fmt.Println(MostFrequentNumberFollowingKeyInAnArray([]int{2, 2, 2, 2, 3}, 2))          // 2
}

// Time: O(n), Space: O(n)
func MostFrequentNumberFollowingKeyInAnArray(nums []int, key int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
  // Linear scan O(n)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			freq[nums[i+1]]++
		}
	}

	maxCount := 0
	result := 0
	for num, count := range freq {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}
	return result
}
```

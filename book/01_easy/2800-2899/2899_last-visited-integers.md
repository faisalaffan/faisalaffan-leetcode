# 2899 — Last Visited Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func LastVisitedIntegers(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2899: Last Visited Integers
// https://leetcode.com/problems/last-visited-integers/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: lastVisitedIntegers
	fmt.Println(LastVisitedIntegers([]int{1, 2, -1, -1, -1})) // [2, 1, -1]
	fmt.Println(LastVisitedIntegers([]int{1, -1, 2, -1, -1})) // [1, 2, 1]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: lastVisitedIntegers
func LastVisitedIntegers(nums []int) []int {
	seen := []int{}
	result := []int{}
	k := 0

	for _, num := range nums {
		if num != -1 {
			seen = append(seen, num)
			k = 0
		} else {
			k++
			if k <= len(seen) {
				result = append(result, seen[len(seen)-k])
			} else {
				result = append(result, -1)
			}
		}
	}
	return result
}
```

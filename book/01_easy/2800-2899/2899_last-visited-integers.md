# 2899 — Last Visited Integers

## Deskripsi

**Soal:** [2899. Last Visited Integers](https://leetcode.com/problems/last-visited-integers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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

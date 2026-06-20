# 1103 — Distribute Candies To People

## Deskripsi

**Soal:** [1103. Distribute Candies To People](https://leetcode.com/problems/distribute-candies-to-people/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(sqrt(candies))  
**Kompleksitas Ruang:** O(numPeople)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1103: Distribute Candies to People
// https://leetcode.com/problems/distribute-candies-to-people/
// Difficulty: Easy
// Time: O(sqrt(candies)) | Space: O(numPeople)

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))  // [1,2,3,1]
	fmt.Println(distributeCandies(10, 3)) // [5,2,3]
}

// LeetCode submission: distributeCandies
func distributeCandies(candies int, numPeople int) []int {
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, numPeople)
	give := 1
	for candies > 0 {
		for i := 0; i < numPeople && candies > 0; i++ {
			if give <= candies {
				ans[i] += give
				candies -= give
			} else {
				ans[i] += candies
				candies = 0
			}
			give++
		}
	}
	return ans
}
```

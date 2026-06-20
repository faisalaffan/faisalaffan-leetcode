# 1103 — Distribute Candies To People

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func distributeCandies(candies int, numPeople int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(sqrt(candies))  |  **Ruang:** O(numPeople)


## 💻 Solusi Go

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
  // Alokasi slice
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

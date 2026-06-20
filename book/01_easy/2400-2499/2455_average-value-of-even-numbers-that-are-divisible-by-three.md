# 2455 — Average Value Of Even Numbers That Are Divisible By Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func AverageValueOfEvenNumbersThatAreDivisibleByThree(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2455: Average Value of Even Numbers That Are Divisible by Three
// https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(AverageValueOfEvenNumbersThatAreDivisibleByThree([]int{1, 3, 6, 10, 12, 15})) // 9
	fmt.Println(AverageValueOfEvenNumbersThatAreDivisibleByThree([]int{1, 2, 4, 7, 10}))       // 0
}

func AverageValueOfEvenNumbersThatAreDivisibleByThree(nums []int) int {
	sum, count := 0, 0
	for _, n := range nums {
		if n%6 == 0 {
			sum += n
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
```

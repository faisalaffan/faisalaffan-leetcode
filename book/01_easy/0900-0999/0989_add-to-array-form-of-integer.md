# 0989 — Add To Array Form Of Integer

## Deskripsi

**Soal:** [0989. Add To Array Form Of Integer](https://leetcode.com/problems/add-to-array-form-of-integer/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(max(n, log k)). Space: O(max(n, log k)).  
**Kompleksitas Ruang:** O(max(n, log k)).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #989: Add to Array-Form of Integer
// https://leetcode.com/problems/add-to-array-form-of-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34)) // [1,2,3,4]
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))   // [4,5,5]
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))   // [1,0,2,1]
}

// addToArrayForm adds an integer to the array-form of a number.
// Time: O(max(n, log k)). Space: O(max(n, log k)).
func addToArrayForm(num []int, k int) []int {
	i := len(num) - 1
  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	carry := 0
	for i >= 0 || k > 0 || carry > 0 {
		digit := carry
		if i >= 0 {
			digit += num[i]
			i--
		}
		if k > 0 {
			digit += k % 10
			k /= 10
		}
		result = append(result, digit%10)
		carry = digit / 10
	}
	// Reverse
	l, r := 0, len(result)-1
	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}
	return result
}
```

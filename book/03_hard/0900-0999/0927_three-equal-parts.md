# 0927 — Three Equal Parts

## Deskripsi

**Soal:** [0927. Three Equal Parts](https://leetcode.com/problems/three-equal-parts/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func threeEqualParts(arr []int) []int`

## Solusi Go

```go
package main

// LeetCode #927: Three Equal Parts
// https://leetcode.com/problems/three-equal-parts/
// Difficulty: Hard
// Partition binary array into 3 parts with equal binary value.
// Leading zeros allowed. Parts must be non-empty.

import "fmt"

func threeEqualParts(arr []int) []int {
	totalOnes := 0
	for _, v := range arr {
		if v == 1 {
			totalOnes++
		}
	}

	if totalOnes%3 != 0 {
		return []int{-1, -1}
	}

	n := len(arr)
	if totalOnes == 0 {
		return []int{0, n - 1}
	}

	onesPerPart := totalOnes / 3

	// Find positions of the first, second, and third part's first 1
	first1 := -1
	second1 := -1
	third1 := -1

	count := 0
	for i, v := range arr {
		if v == 1 {
			count++
			if count == 1 {
				first1 = i
			}
			if count == onesPerPart+1 {
				second1 = i
			}
			if count == 2*onesPerPart+1 {
				third1 = i
			}
		}
	}

	// Now compare each part: arr[first1..second1-1], arr[second1..third1-1], arr[third1..]
	a, b, c := first1, second1, third1
	for c < n {
		if arr[a] != arr[b] || arr[b] != arr[c] {
			return []int{-1, -1}
		}
		a++
		b++
		c++
	}

	return []int{a - 1, b - 1}
}

func main() {
	fmt.Println(threeEqualParts([]int{1,0,1,0,1})) // Expected: [0,3]
	fmt.Println(threeEqualParts([]int{1,1,0,0,1})) // Expected: [0,2]
	fmt.Println(threeEqualParts([]int{1,1,0,1,1})) // Expected: [-1,-1]
	fmt.Println(threeEqualParts([]int{0,0,0,0,0})) // Expected: [0,4]
	fmt.Println(threeEqualParts([]int{1,0,1,0,1,0})) // Expected: [1,4]
}
```

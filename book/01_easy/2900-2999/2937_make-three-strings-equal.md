# 2937 — Make Three Strings Equal

## Deskripsi

**Soal:** [2937. Make Three Strings Equal](https://leetcode.com/problems/make-three-strings-equal/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(min(len(s1), len(s2), len(s3)))  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2937: Make Three Strings Equal
// https://leetcode.com/problems/make-three-strings-equal/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findMinimumOperations
	fmt.Println(MakeThreeStringsEqual("abc", "abb", "ab")) // 2
	fmt.Println(MakeThreeStringsEqual("dac", "bac", "cac")) // -1
	fmt.Println(MakeThreeStringsEqual("a", "a", "a"))       // 0
}

// Time: O(min(len(s1), len(s2), len(s3))) | Space: O(1)
// LeetCode submission name: findMinimumOperations
func MakeThreeStringsEqual(s1 string, s2 string, s3 string) int {
	// Find longest common prefix length
	i := 0
	for i < len(s1) && i < len(s2) && i < len(s3) {
		if s1[i] != s2[i] || s1[i] != s3[i] {
			break
		}
		i++
	}
	if i == 0 {
		return -1
	}
	return (len(s1) - i) + (len(s2) - i) + (len(s3) - i)
}
```

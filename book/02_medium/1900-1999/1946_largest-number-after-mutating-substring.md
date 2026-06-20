# 1946 — Largest Number After Mutating Substring

## Deskripsi

**Soal:** [1946. Largest Number After Mutating Substring](https://leetcode.com/problems/largest-number-after-mutating-substring/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1946: Largest Number After Mutating Substring
// https://leetcode.com/problems/largest-number-after-mutating-substring/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximumNumber("132", []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8}))
	fmt.Println(MaximumNumber("021", []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}))
	fmt.Println(MaximumNumber("5", []int{1, 4, 7, 5, 3, 2, 5, 6, 9, 4}))
}

// Time: O(n), Space: O(n)
func MaximumNumber(num string, change []int) string {
	n := len(num)
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, n)
	i := 0

	// Skip prefix where change doesn't increase value
	for i < n && change[num[i]-'0'] <= int(num[i]-'0') {
		result[i] = num[i]
		i++
	}

	// Mutate while change increases or keeps same value
	for i < n && change[num[i]-'0'] >= int(num[i]-'0') {
		result[i] = byte(change[num[i]-'0'] + '0')
		i++
	}

	// Copy remaining
	for i < n {
		result[i] = num[i]
		i++
	}
	return string(result)
}
```

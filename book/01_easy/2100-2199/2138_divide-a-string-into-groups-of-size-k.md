# 2138 — Divide A String Into Groups Of Size K

## Deskripsi

**Soal:** [2138. Divide A String Into Groups Of Size K](https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2138: Divide a String Into Groups of Size k
// https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DivideAStringIntoGroupsOfSizeK("abcdefghi", 3, 'x')) // ["abc" "def" "ghi"]
	fmt.Println(DivideAStringIntoGroupsOfSizeK("abcdefghij", 3, 'x')) // ["abc" "def" "ghi" "jxx"]
}

// Time: O(n), Space: O(n)
func DivideAStringIntoGroupsOfSizeK(s string, k int, fill byte) []string {
	var result []string
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i += k {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		group := s[i:end]
		if len(group) < k {
			for len(group) < k {
				group += string(fill)
			}
		}
		result = append(result, group)
	}
	return result
}
```

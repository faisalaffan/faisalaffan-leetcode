# 0779 — K Th Symbol In Grammar

## Deskripsi

**Soal:** [0779. K Th Symbol In Grammar](https://leetcode.com/problems/k-th-symbol-in-grammar/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #779: K-th Symbol in Grammar
// https://leetcode.com/problems/k-th-symbol-in-grammar/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(kthGrammar(1, 1))
	fmt.Println(kthGrammar(2, 1))
	fmt.Println(kthGrammar(2, 2))
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	parent := kthGrammar(n-1, (k+1)/2)
	if k%2 == 0 {
		return 1 - parent
	}
	return parent
}
```

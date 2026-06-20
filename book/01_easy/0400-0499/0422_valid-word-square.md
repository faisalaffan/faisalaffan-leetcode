# 0422 — Valid Word Square

## Deskripsi

**Soal:** [0422. Valid Word Square](https://leetcode.com/problems/valid-word-square/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n*m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ValidWordSquare(words []string) bool`

## Solusi Go

```go
package main

// LeetCode #422: Valid Word Square
// https://leetcode.com/problems/valid-word-square/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n*m), Space: O(1)
func ValidWordSquare(words []string) bool {
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) || i >= len(words[j]) || words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crmy", "dtyx"}))
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))
	fmt.Println(ValidWordSquare([]string{"ball", "area", "lead", "lady"}))
}
```

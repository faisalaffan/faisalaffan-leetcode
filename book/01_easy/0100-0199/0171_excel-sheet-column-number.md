# 0171 — Excel Sheet Column Number

## Deskripsi

**Soal:** [0171. Excel Sheet Column Number](https://leetcode.com/problems/excel-sheet-column-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func TitleToNumber(columnTitle string) int`

## Solusi Go

```go
package main

// LeetCode #171: Excel Sheet Column Number
// https://leetcode.com/problems/excel-sheet-column-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func TitleToNumber(columnTitle string) int {
	result := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(columnTitle); i++ {
		result = result*26 + int(columnTitle[i]-'A'+1)
	}
	return result
}

func main() {
	fmt.Println(TitleToNumber("A"))
	fmt.Println(TitleToNumber("AB"))
	fmt.Println(TitleToNumber("ZY"))
}
```

# 1417 — Reformat The String

## Deskripsi

**Soal:** [1417. Reformat The String](https://leetcode.com/problems/reformat-the-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func reformat(s string) string`

## Solusi Go

```go
package main

// LeetCode #1417: Reformat The String
// https://leetcode.com/problems/reformat-the-string/
// Difficulty: Easy
//
// LeetCode submission: func reformat(s string) string

import "fmt"

func main() {
	fmt.Println(ReformatTheString("a0b1c2")) // "a0b1c2"
	fmt.Println(ReformatTheString("leetcode")) // ""
	fmt.Println(ReformatTheString("1229857369")) // ""
}

// Time: O(n), Space: O(n)
func ReformatTheString(s string) string {
  // Membuat slice untuk menyimpan hasil
	letters := make([]byte, 0, len(s))
  // Membuat slice untuk menyimpan hasil
	digits := make([]byte, 0, len(s))
  // Iterasi seluruh elemen
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			letters = append(letters, s[i])
		} else {
			digits = append(digits, s[i])
		}
	}
	if abs(len(letters)-len(digits)) > 1 {
		return ""
	}
  // Membuat slice untuk menyimpan hasil
	res := make([]byte, len(s))
	var first, second []byte
	if len(letters) >= len(digits) {
		first, second = letters, digits
	} else {
		first, second = digits, letters
	}
	idx := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(first); i++ {
		res[idx] = first[i]
		idx++
		if i < len(second) {
			res[idx] = second[i]
			idx++
		}
	}
	return string(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

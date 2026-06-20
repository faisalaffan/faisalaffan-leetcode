# 1119 — Remove Vowels From A String

## Deskripsi

**Soal:** [1119. Remove Vowels From A String](https://leetcode.com/problems/remove-vowels-from-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1119: Remove Vowels from a String
// https://leetcode.com/problems/remove-vowels-from-a-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeVowels("leetcodeisacommunityforcoders")) // "ltcdscmmntyfrcdrs"
	fmt.Println(removeVowels("aeiou"))                         // ""
}

// LeetCode submission: removeVowels
func removeVowels(s string) string {
  // Membuat slice untuk menyimpan hasil
	ans := make([]byte, 0, len(s))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
```

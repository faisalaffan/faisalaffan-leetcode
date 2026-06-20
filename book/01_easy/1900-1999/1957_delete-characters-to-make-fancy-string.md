# 1957 — Delete Characters To Make Fancy String

## Deskripsi

**Soal:** [1957. Delete Characters To Make Fancy String](https://leetcode.com/problems/delete-characters-to-make-fancy-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1957: Delete Characters to Make Fancy String
// https://leetcode.com/problems/delete-characters-to-make-fancy-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DeleteCharactersToMakeFancyString("leeetcode"))     // "leetcode"
	fmt.Println(DeleteCharactersToMakeFancyString("aaabaaaa"))      // "aabaa"
	fmt.Println(DeleteCharactersToMakeFancyString("aab"))           // "aab"
}

// Time: O(n), Space: O(n)
func DeleteCharactersToMakeFancyString(s string) string {
  // Membuat slice untuk menyimpan hasil
	result := make([]byte, 0, len(s))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		n := len(result)
		if n >= 2 && result[n-1] == s[i] && result[n-2] == s[i] {
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
```

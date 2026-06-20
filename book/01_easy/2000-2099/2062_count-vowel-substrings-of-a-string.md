# 2062 — Count Vowel Substrings Of A String

## Deskripsi

**Soal:** [2062. Count Vowel Substrings Of A String](https://leetcode.com/problems/count-vowel-substrings-of-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2062: Count Vowel Substrings of a String
// https://leetcode.com/problems/count-vowel-substrings-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountVowelSubstringsOfAString("aeiouu"))    // 2
	fmt.Println(CountVowelSubstringsOfAString("unicornarihan")) // 0
	fmt.Println(CountVowelSubstringsOfAString("cuaieuouac"))    // 7
}

// Time: O(n^2), Space: O(1)
func CountVowelSubstringsOfAString(word string) int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word); i++ {
  // Membuat map untuk pencarian O(1): key → value
		vowelSet := make(map[byte]bool)
		for j := i; j < len(word); j++ {
			if !isVowel(word[j]) {
				break
			}
			vowelSet[word[j]] = true
			if len(vowelSet) == 5 {
				count++
			}
		}
	}
	return count
}
```

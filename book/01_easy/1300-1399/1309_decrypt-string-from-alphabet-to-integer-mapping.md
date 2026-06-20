# 1309 — Decrypt String From Alphabet To Integer Mapping

## Deskripsi

**Soal:** [1309. Decrypt String From Alphabet To Integer Mapping](https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func freqAlphabets(s string) string`

## Solusi Go

```go
package main

// LeetCode #1309: Decrypt String from Alphabet to Integer Mapping
// https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// Difficulty: Easy
//
// LeetCode submission: func freqAlphabets(s string) string

import "fmt"

func main() {
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("10#11#12"))       // "jkab"
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("1326#"))           // "acz"
	fmt.Println(DecryptStringFromAlphabetToIntegerMapping("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#")) // "abcdefghijklmnopqrstuvwxyz"
}

// Time: O(n), Space: O(n)
func DecryptStringFromAlphabetToIntegerMapping(s string) string {
  // Membuat slice untuk menyimpan hasil
	res := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		if i+2 < len(s) && s[i+2] == '#' {
			num := (s[i]-'0')*10 + (s[i+1] - '0')
			res = append(res, byte('a'+num-1))
			i += 3
		} else {
			num := s[i] - '0'
			res = append(res, byte('a'+num-1))
			i++
		}
	}
	return string(res)
}
```

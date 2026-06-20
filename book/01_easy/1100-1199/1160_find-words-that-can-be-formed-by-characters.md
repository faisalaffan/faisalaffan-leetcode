# 1160 — Find Words That Can Be Formed By Characters

## Deskripsi

**Soal:** [1160. Find Words That Can Be Formed By Characters](https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * k) where k is max word length  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1160: Find Words That Can Be Formed by Characters
// https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
// Difficulty: Easy
// Time: O(n * k) where k is max word length | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach")) // 6
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr")) // 10
}

// LeetCode submission: countCharacters
func countCharacters(words []string, chars string) int {
  // Membuat slice untuk menyimpan hasil
	ch := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(chars); i++ {
		ch[chars[i]-'a']++
	}
	ans := 0
	for _, w := range words {
  // Membuat slice untuk menyimpan hasil
		need := make([]int, 26)
		ok := true
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(w); i++ {
			idx := w[i] - 'a'
			need[idx]++
			if need[idx] > ch[idx] {
				ok = false
				break
			}
		}
		if ok {
			ans += len(w)
		}
	}
	return ans
}
```

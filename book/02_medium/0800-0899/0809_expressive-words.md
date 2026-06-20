# 0809 — Expressive Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ExpressiveWords(s string, words []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #809: Expressive Words
// https://leetcode.com/problems/expressive-words/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ExpressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	fmt.Println(ExpressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
	fmt.Println(ExpressiveWords("abcd", []string{"abc"}))
}

func ExpressiveWords(s string, words []string) int {
	type group struct {
		char byte
		cnt  int
	}

	var encode func(string) []group
	encode = func(str string) []group {
		var groups []group
  // Linear scan O(n)
		for i := 0; i < len(str); {
			j := i
			for j < len(str) && str[j] == str[i] {
				j++
			}
			groups = append(groups, group{str[i], j - i})
			i = j
		}
		return groups
	}

	sGroups := encode(s)
	count := 0

	for _, word := range words {
		wGroups := encode(word)
		if len(wGroups) != len(sGroups) {
			continue
		}
		ok := true
  // Range loop
		for i := range sGroups {
			if sGroups[i].char != wGroups[i].char {
				ok = false
				break
			}
			if sGroups[i].cnt < 3 && sGroups[i].cnt != wGroups[i].cnt {
				ok = false
				break
			}
			if sGroups[i].cnt >= 3 && wGroups[i].cnt > sGroups[i].cnt {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}

	return count
}
```

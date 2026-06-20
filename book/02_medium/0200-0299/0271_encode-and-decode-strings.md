# 0271 — Encode And Decode Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #271: Encode and Decode Strings
// https://leetcode.com/problems/encode-and-decode-strings/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct{}

func (codec *Codec) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

func (codec *Codec) Decode(s string) []string {
	result := []string{}
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		start := j + 1
		result = append(result, s[start:start+length])
		i = start + length
	}
	return result
}

func main() {
	codec := &Codec{}
	encoded := codec.Encode([]string{"hello", "world", "leet", "code"})
	fmt.Println(encoded)
	fmt.Println(codec.Decode(encoded))

	encoded2 := codec.Encode([]string{""})
	fmt.Println(codec.Decode(encoded2))

	encoded3 := codec.Encode([]string{"a", "", "b"})
	fmt.Println(codec.Decode(encoded3))
}
```

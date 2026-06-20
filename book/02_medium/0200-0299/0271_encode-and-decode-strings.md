# 0271 — Encode And Decode Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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

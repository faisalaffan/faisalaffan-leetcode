# 0722 — Remove Comments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeComments(source []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #722: Remove Comments
// https://leetcode.com/problems/remove-comments/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	source := []string{
		"/*Test program */",
		"int main()",
		"{ ",
		"  // variable declaration ",
		"int a, b, c;",
		"/* This is a test",
		"   multiline  ",
		"   comment for ",
		"   testing */",
		"a = b + c;",
		"}",
	}
	result := removeComments(source)
	for _, line := range result {
		fmt.Println(line)
	}
}

func removeComments(source []string) []string {
	result := make([]string, 0)
	inBlock := false
	var current strings.Builder

	for _, line := range source {
		i := 0
		n := len(line)

		if !inBlock {
			current.Reset()
		}

		for i < n {
			if !inBlock && i+1 < n && line[i] == '/' && line[i+1] == '*' {
				inBlock = true
				i += 2
			} else if inBlock && i+1 < n && line[i] == '*' && line[i+1] == '/' {
				inBlock = false
				i += 2
			} else if !inBlock && i+1 < n && line[i] == '/' && line[i+1] == '/' {
				break
			} else if !inBlock {
				current.WriteByte(line[i])
				i++
			} else {
				i++
			}
		}

		if !inBlock && current.Len() > 0 {
			result = append(result, current.String())
		}
	}

	return result
}
```

# 0809 — Expressive Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ExpressiveWords(s string, words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Loop linear O(n): iterasi setiap elemen
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
  // Range loop: iterasi dengan indeks + nilai
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

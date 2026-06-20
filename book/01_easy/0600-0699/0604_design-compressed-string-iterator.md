# 0604 — Design Compressed String Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(compressedString string) StringIterator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #604: Design Compressed String Iterator
// https://leetcode.com/problems/design-compressed-string-iterator/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"strconv"
)

// StringIterator iterates over a compressed string.
type StringIterator struct {
	chars []byte
	counts []int
	index  int
}

// Constructor creates a new StringIterator from compressed string.
func Constructor(compressedString string) StringIterator {
	var chars []byte
	var counts []int
	i := 0
	for i < len(compressedString) {
		c := compressedString[i]
		i++
		numStart := i
		for i < len(compressedString) && compressedString[i] >= '0' && compressedString[i] <= '9' {
			i++
		}
		count, _ := strconv.Atoi(compressedString[numStart:i])
		chars = append(chars, c)
		counts = append(counts, count)
	}
	return StringIterator{chars: chars, counts: counts, index: 0}
}

// Next returns the next character or ' ' if exhausted.
// Time: O(1), Space: O(1)
func (it *StringIterator) Next() byte {
	if !it.HasNext() {
		return ' '
	}
	c := it.chars[it.index]
	it.counts[it.index]--
	if it.counts[it.index] == 0 {
		it.index++
	}
	return c
}

// HasNext returns true if there are more characters.
// Time: O(1), Space: O(1)
func (it *StringIterator) HasNext() bool {
	return it.index < len(it.chars)
}

func main() {
	obj := Constructor("L1e2t1C1o1d1e1")
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c ", obj.Next())
	fmt.Printf("%c\n", obj.Next())
	fmt.Println(obj.HasNext())
}
```

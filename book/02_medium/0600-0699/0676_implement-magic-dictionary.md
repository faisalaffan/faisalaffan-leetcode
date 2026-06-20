# 0676 — Implement Magic Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L) for buildDict, O(26 * L) for search  
**Kompleksitas Ruang:** O(n * L)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #676: Implement Magic Dictionary
// https://leetcode.com/problems/implement-magic-dictionary/
// Difficulty: Medium
// Time: O(n * L) for buildDict, O(26 * L) for search
// Space: O(n * L)

import "fmt"

func main() {
	md := MagicDictionary{}
	md.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(md.Search("hello"))
	fmt.Println(md.Search("hhllo"))
	fmt.Println(md.Search("hell"))
	fmt.Println(md.Search("leetcoded"))
}

type MagicDictionary struct {
	words []string
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	m.words = dictionary
}

func (m *MagicDictionary) Search(searchWord string) bool {
	for _, word := range m.words {
		if len(word) != len(searchWord) {
			continue
		}
		diff := 0
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diff++
			}
			if diff > 1 {
				break
			}
		}
		if diff == 1 {
			return true
		}
	}
	return false
}
```

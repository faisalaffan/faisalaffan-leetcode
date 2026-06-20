# 0211 — Design Add And Search Words Data Structure

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() WordDictionary
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) for add, O(26^m) worst case for search with wildcards, Space: O(total chars)  
**Kompleksitas Ruang:** O(total chars)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #211: Design Add and Search Words Data Structure
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
// Difficulty: Medium
// Time: O(n) for add, O(26^m) worst case for search with wildcards, Space: O(total chars)

import "fmt"

type WordDictionaryNode struct {
	children [26]*WordDictionaryNode
	isEnd    bool
}

type WordDictionary struct {
	root *WordDictionaryNode
}

func Constructor() WordDictionary {
	return WordDictionary{&WordDictionaryNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionaryNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word, 0)
}

func (this *WordDictionary) search(node *WordDictionaryNode, word string, idx int) bool {
	if node == nil {
		return false
	}
	if idx == len(word) {
		return node.isEnd
	}

	if word[idx] == '.' {
		for i := 0; i < 26; i++ {
			if this.search(node.children[i], word, idx+1) {
				return true
			}
		}
		return false
	}

	child := node.children[word[idx]-'a']
	return this.search(child, word, idx+1)
}

func main() {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad"))
	fmt.Println(wd.Search("bad"))
	fmt.Println(wd.Search(".ad"))
	fmt.Println(wd.Search("b.."))
}
```

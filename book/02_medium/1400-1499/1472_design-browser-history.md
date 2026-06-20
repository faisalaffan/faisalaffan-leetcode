# 1472 — Design Browser History

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewBrowserHistory(homepage string) BrowserHistory
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1472: Design Browser History
// https://leetcode.com/problems/design-browser-history/
// Difficulty: Medium

import "fmt"

type BrowserHistory struct {
	history []string
	current int
}

func main() {
	bh := NewBrowserHistory("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")
	fmt.Println(bh.Back(1)) // "facebook.com"
	fmt.Println(bh.Back(1)) // "google.com"
	fmt.Println(bh.Forward(1)) // "facebook.com"
	bh.Visit("linkedin.com")
	fmt.Println(bh.Forward(2)) // "linkedin.com"
	fmt.Println(bh.Back(2)) // "google.com"
	fmt.Println(bh.Back(7)) // "leetcode.com"

	bh2 := NewBrowserHistory("a.com")
	bh2.Visit("b.com")
	fmt.Println(bh2.Back(1)) // "a.com"
	fmt.Println(bh2.Forward(1)) // "b.com"
}

func NewBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{
		history: []string{homepage},
		current: 0,
	}
}

// Time: O(1)
func (this *BrowserHistory) Visit(url string) {
	this.current++
	this.history = this.history[:this.current]
	this.history = append(this.history, url)
}

// Time: O(1)
func (this *BrowserHistory) Back(steps int) string {
	this.current -= steps
	if this.current < 0 {
		this.current = 0
	}
	return this.history[this.current]
}

// Time: O(1)
func (this *BrowserHistory) Forward(steps int) string {
	this.current += steps
	if this.current >= len(this.history) {
		this.current = len(this.history) - 1
	}
	return this.history[this.current]
}
```

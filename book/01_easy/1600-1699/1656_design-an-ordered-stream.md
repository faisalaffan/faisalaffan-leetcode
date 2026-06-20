# 1656 — Design An Ordered Stream

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n int) OrderedStream
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1656: Design an Ordered Stream
// https://leetcode.com/problems/design-an-ordered-stream/
// Difficulty: Easy

import "fmt"

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{stream: make([]string, n+1), ptr: 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	var result []string
	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		result = append(result, this.stream[this.ptr])
		this.ptr++
	}
	return result
}

func main() {
	os := Constructor(5)
	fmt.Println(os.Insert(3, "ccccc"))
	fmt.Println(os.Insert(1, "aaaaa"))
	fmt.Println(os.Insert(2, "bbbbb"))
	fmt.Println(os.Insert(5, "eeeee"))
	fmt.Println(os.Insert(4, "ddddd"))
}
```

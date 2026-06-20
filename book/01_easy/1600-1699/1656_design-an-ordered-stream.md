# 1656 — Design An Ordered Stream

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor(n int) OrderedStream`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


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

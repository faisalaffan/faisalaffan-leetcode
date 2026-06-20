# 1656 — Design An Ordered Stream

## Deskripsi

**Soal:** [1656. Design An Ordered Stream](https://leetcode.com/problems/design-an-ordered-stream/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(n int) OrderedStream`

## Solusi Go

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

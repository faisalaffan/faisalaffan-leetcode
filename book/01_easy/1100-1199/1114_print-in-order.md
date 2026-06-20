# 1114 — Print In Order

## Deskripsi

**Soal:** [1114. Print In Order](https://leetcode.com/problems/print-in-order/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func NewFoo() *Foo`

## Solusi Go

```go
package main

// LeetCode #1114: Print in Order
// https://leetcode.com/problems/print-in-order/
// Difficulty: Easy (Concurrency)
// Time: O(1) | Space: O(1)

import (
	"fmt"
	"sync"
)

type Foo struct {
	wg1 sync.WaitGroup
	wg2 sync.WaitGroup
}

func NewFoo() *Foo {
	f := &Foo{}
	f.wg1.Add(1)
	f.wg2.Add(1)
	return f
}

func (f *Foo) first() {
	fmt.Print("first")
	f.wg1.Done()
}

func (f *Foo) second() {
	f.wg1.Wait()
	fmt.Print("second")
	f.wg2.Done()
}

func (f *Foo) third() {
	f.wg2.Wait()
	fmt.Print("third")
}

func main() {
	// Test: run in order 1,2,3
	f := NewFoo()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() { f.first(); wg.Done() }()
	go func() { f.second(); wg.Done() }()
	go func() { f.third(); wg.Done() }()
	wg.Wait()
	fmt.Println()
}
```

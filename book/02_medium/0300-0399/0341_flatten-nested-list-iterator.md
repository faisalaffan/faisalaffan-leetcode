# 0341 — Flatten Nested List Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NewInt(val int) *NestedInteger`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n) initialization, O(1) next/hasNext  |  **Ruang:** O(d)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #341: Flatten Nested List Iterator
// https://leetcode.com/problems/flatten-nested-list-iterator/
// Difficulty: Medium
// Time: O(n) initialization, O(1) next/hasNext | Space: O(d)

import "fmt"

type NestedInteger struct {
	integer int
	list    []*NestedInteger
	isInt   bool
}

func NewInt(val int) *NestedInteger {
	return &NestedInteger{integer: val, isInt: true}
}

func NewList(items ...*NestedInteger) *NestedInteger {
	return &NestedInteger{list: items, isInt: false}
}

func (n NestedInteger) IsInteger() bool          { return n.isInt }
func (n NestedInteger) GetInteger() int           { return n.integer }
func (n NestedInteger) GetList() []*NestedInteger { return n.list }

type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := make([]*NestedInteger, 0)
	// Push in reverse order
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	return &NestedIterator{stack: stack}
}

func (it *NestedIterator) Next() int {
	top := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	return top.GetInteger()
}

func (it *NestedIterator) HasNext() bool {
	for len(it.stack) > 0 {
		top := it.stack[len(it.stack)-1]
		if top.IsInteger() {
			return true
		}
		it.stack = it.stack[:len(it.stack)-1]
		list := top.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			it.stack = append(it.stack, list[i])
		}
	}
	return false
}

func main() {
	// Test case 1: [[1,1],2,[1,1]]
	n1 := NewList(NewInt(1), NewInt(1))
	n2 := NewInt(2)
	n3 := NewList(NewInt(1), NewInt(1))
	it1 := Constructor([]*NestedInteger{n1, n2, n3})
	fmt.Print("Test 1: ")
	for it1.HasNext() {
		fmt.Print(it1.Next(), " ")
	}
	fmt.Println()
	// Expected: 1 1 2 1 1

	// Test case 2: [1,[4,[6]]]
	inner := NewList(NewInt(6))
	mid := NewList(NewInt(4), inner)
	it2 := Constructor([]*NestedInteger{NewInt(1), mid})
	fmt.Print("Test 2: ")
	for it2.HasNext() {
		fmt.Print(it2.Next(), " ")
	}
	fmt.Println()
	// Expected: 1 4 6

	// Test case 3: Empty
	it3 := Constructor([]*NestedInteger{})
	fmt.Print("Test 3: ")
	for it3.HasNext() {
		fmt.Print(it3.Next(), " ")
	}
	fmt.Println()
	// Expected: (nothing)
}
```

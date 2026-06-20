# 0385 — Mini Parser

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewInt(val int) *NestedInteger
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #385: Mini Parser
// https://leetcode.com/problems/mini-parser/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

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
func (n *NestedInteger) Add(elem NestedInteger) {
	n.list = append(n.list, &elem)
}
func (n *NestedInteger) SetInteger(val int) {
	n.integer = val
	n.isInt = true
}

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}
	if s[0] != '[' {
		val, _ := strconv.Atoi(s)
		return NewInt(val)
	}

	stack := []*NestedInteger{}
	var current *NestedInteger
	i := 0

	for i < len(s) {
		if s[i] == '[' {
			ni := &NestedInteger{isInt: false, list: []*NestedInteger{}}
			if current != nil {
				stack = append(stack, current)
			}
			current = ni
			i++
		} else if s[i] == ']' {
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				parent.list = append(parent.list, current)
				current = parent
			}
			i++
		} else if s[i] == ',' {
			i++
		} else {
			// Parse integer
			j := i
			for j < len(s) && (s[j] == '-' || (s[j] >= '0' && s[j] <= '9')) {
				j++
			}
			val, _ := strconv.Atoi(s[i:j])
			nestedInt := NewInt(val)
			current.list = append(current.list, nestedInt)
			i = j
		}
	}
	return current
}

func (n *NestedInteger) String() string {
	if n.isInt {
		return strconv.Itoa(n.integer)
	}
	parts := make([]string, len(n.list))
	for i, child := range n.list {
		parts[i] = child.String()
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func main() {
	// Test case 1
	n1 := deserialize("324")
	fmt.Println("Test 1:", n1.String())
	// Expected: 324

	// Test case 2
	n2 := deserialize("[123,[456,[789]]]")
	fmt.Println("Test 2:", n2.String())
	// Expected: [123,[456,[789]]]

	// Test case 3
	n3 := deserialize("[]")
	fmt.Println("Test 3:", n3.String())
	// Expected: []
}
```

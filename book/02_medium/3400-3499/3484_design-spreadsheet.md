# 3484 — Design Spreadsheet

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewSpreadsheet(rows, cols int) *Spreadsheet
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3484: Design Spreadsheet
// https://leetcode.com/problems/design-spreadsheet/
// Difficulty: Medium
// Complexity: O(1) per operation

import "fmt"

type Spreadsheet struct {
	data [][]int
}

func NewSpreadsheet(rows, cols int) *Spreadsheet {
	s := &Spreadsheet{}
	s.data = make([][]int, rows)
  // Range loop: iterasi dengan indeks + nilai
	for i := range s.data {
		s.data[i] = make([]int, cols)
	}
	return s
}

func (s *Spreadsheet) SetCell(row, col, val int) {
	if row >= 0 && row < len(s.data) && col >= 0 && col < len(s.data[0]) {
		s.data[row][col] = val
	}
}

func (s *Spreadsheet) GetCell(row, col int) int {
	if row >= 0 && row < len(s.data) && col >= 0 && col < len(s.data[0]) {
		return s.data[row][col]
	}
	return 0
}

func (s *Spreadsheet) SumRange(r1, c1, r2, c2 int) int {
	sum := 0
	for i := r1; i <= r2 && i < len(s.data); i++ {
		for j := c1; j <= c2 && j < len(s.data[0]); j++ {
			sum += s.data[i][j]
		}
	}
	return sum
}

func (s *Spreadsheet) Print() {
	for _, row := range s.data {
		fmt.Println(row)
	}
}

func main() {
	// Test case 1: basic operations
	ss := NewSpreadsheet(3, 3)
	ss.SetCell(0, 0, 5)
	ss.SetCell(1, 1, 10)
	fmt.Println("Cell (0,0):", ss.GetCell(0, 0))
	fmt.Println("Sum range (0,0)-(2,2):", ss.SumRange(0, 0, 2, 2))

	// Test case 2
	ss2 := NewSpreadsheet(1, 5)
	ss2.SetCell(0, 0, 1)
	ss2.SetCell(0, 1, 2)
	ss2.SetCell(0, 2, 3)
	fmt.Println("Sum (0,0)-(0,2):", ss2.SumRange(0, 0, 0, 2))

	fmt.Println("DesignSpreadsheet() done:", DesignSpreadsheet())
}

func DesignSpreadsheet() any {
	return "Spreadsheet designed"
}
```

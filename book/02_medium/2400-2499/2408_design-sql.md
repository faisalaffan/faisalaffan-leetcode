# 2408 — Design Sql

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() SQL
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(1) per operation if table exists  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2408: Design SQL
// https://leetcode.com/problems/design-sql/
// Difficulty: Medium
// Time: O(1) per operation if table exists | Space: O(n)
// In-memory SQL with tables, rows, and cell access.

import "fmt"

type Table struct {
	columns int
	rows    map[int][]string
	nextID  int
}

type SQL struct {
	tables map[string]*Table
}

func main() {
	sql := Constructor()
	sql.CreateTable("students", 3)
	sql.InsertRow("students", []string{"1", "John", "A"})
	sql.InsertRow("students", []string{"2", "Jane", "B"})
	fmt.Println(sql.SelectCell("students", 1, 2)) // "John" (row 1, col 2)

	sql.CreateTable("courses", 2)
	rowID := sql.InsertRow("courses", []string{"101", "Math"})
	fmt.Println(sql.SelectCell("courses", rowID, 2)) // "Math"
}

func Constructor() SQL {
	return SQL{tables: make(map[string]*Table)}
}

func (s *SQL) CreateTable(name string, columns int) {
	s.tables[name] = &Table{columns: columns, rows: make(map[int][]string), nextID: 1}
}

func (s *SQL) InsertRow(name string, values []string) int {
	t := s.tables[name]
	id := t.nextID
	t.nextID++
	row := make([]string, len(values))
	copy(row, values)
	t.rows[id] = row
	return id
}

func (s *SQL) SelectCell(name string, rowID, col int) string {
	return s.tables[name].rows[rowID][col-1]
}

func (s *SQL) DeleteRow(name string, rowID int) {
	delete(s.tables[name].rows, rowID)
}
```

# 2408 — Design Sql

## Deskripsi

**Soal:** [2408. Design Sql](https://leetcode.com/problems/design-sql/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per operation if table exists  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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

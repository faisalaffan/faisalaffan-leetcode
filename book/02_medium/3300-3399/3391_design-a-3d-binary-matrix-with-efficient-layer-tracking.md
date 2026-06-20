# 3391 — Design A 3D Binary Matrix With Efficient Layer Tracking

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor(n, m, l int) *Matrix3D`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1) set/get, O(n*m) init  Space: O(n*m*l)  |  **Ruang:** O(n*m*l)


## 💻 Solusi Go

```go
package main

// LeetCode #3391: Design a 3D Binary Matrix with Efficient Layer Tracking
// https://leetcode.com/problems/design-a-3d-binary-matrix-with-efficient-layer-tracking/
// Difficulty: Medium [Paid]
// Time: O(1) set/get, O(n*m) init  Space: O(n*m*l)

import "fmt"

type Matrix3D struct {
	data [][][]bool
	n    int
	m    int
	l    int
}

func Constructor(n, m, l int) *Matrix3D {
	mat := &Matrix3D{
		data: make([][][]bool, n),
		n:    n,
		m:    m,
		l:    l,
	}
	for i := 0; i < n; i++ {
		mat.data[i] = make([][]bool, m)
		for j := 0; j < m; j++ {
			mat.data[i][j] = make([]bool, l)
		}
	}
	return mat
}

func (mat *Matrix3D) SetCell(x, y, z int) {
	mat.data[x][y][z] = true
}

func (mat *Matrix3D) UnsetCell(x, y, z int) {
	mat.data[x][y][z] = false
}

func (mat *Matrix3D) GetCell(x, y, z int) bool {
	return mat.data[x][y][z]
}

func (mat *Matrix3D) LargestLayer() int {
	maxCount := 0
	maxLayer := 0
	for k := 0; k < mat.l; k++ {
		count := 0
		for i := 0; i < mat.n; i++ {
			for j := 0; j < mat.m; j++ {
				if mat.data[i][j][k] {
					count++
				}
			}
		}
		if count > maxCount {
			maxCount = count
			maxLayer = k
		}
	}
	return maxLayer
}

func main() {
	mat := Constructor(2, 2, 2)
	mat.SetCell(0, 0, 0)
	mat.SetCell(1, 1, 1)
	fmt.Println(mat.GetCell(0, 0, 0)) // true
	fmt.Println(mat.GetCell(1, 1, 0)) // false
	fmt.Println(mat.LargestLayer())   // 1 (both layers have 1 set)
	mat.UnsetCell(1, 1, 1)
	fmt.Println(mat.LargestLayer()) // 0
}
```

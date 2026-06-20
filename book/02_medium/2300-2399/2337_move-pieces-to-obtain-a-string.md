# 2337 — Move Pieces To Obtain A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func canChange(start string, target string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2337: Move Pieces to Obtain a String
// https://leetcode.com/problems/move-pieces-to-obtain-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func canChange(start string, target string) bool {
	n := len(start)
	i, j := 0, 0

	for i < n || j < n {
		// Skip underscores
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}

		if i == n && j == n {
			return true
		}
		if i == n || j == n {
			return false
		}

		if start[i] != target[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i++
		j++
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println(canChange("_L__R__R_", "L______RR"))
	// Expected: true

	// Test case 2
	fmt.Println(canChange("R_L_", "__LR"))
	// Expected: false

	// Test case 3
	fmt.Println(canChange("_R", "R_"))
	// Expected: false
}
```

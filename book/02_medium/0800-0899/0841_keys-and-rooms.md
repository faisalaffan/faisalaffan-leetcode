# 0841 — Keys And Rooms

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KeysAndRooms(rooms [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n + k) where k = total keys  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #841: Keys and Rooms
// https://leetcode.com/problems/keys-and-rooms/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(KeysAndRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(KeysAndRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	fmt.Println(KeysAndRooms([][]int{{2}, {}, {1}}))
}

// Time: O(n + k) where k = total keys | Space: O(n)
func KeysAndRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	stack := []int{0}
	visited[0] = true
	count := 1

	for len(stack) > 0 {
		room := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, key := range rooms[room] {
			if !visited[key] {
				visited[key] = true
				count++
				stack = append(stack, key)
			}
		}
	}

	return count == n
}
```

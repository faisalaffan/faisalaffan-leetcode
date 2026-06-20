# 2511 — Maximum Enemy Forts That Can Be Captured

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumEnemyFortsThatCanBeCaptured(forts []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2511: Maximum Enemy Forts That Can Be Captured
// https://leetcode.com/problems/maximum-enemy-forts-that-can-be-captured/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MaximumEnemyFortsThatCanBeCaptured([]int{1, 0, 0, -1, 0, 0, 0, 0, 1})) // 4
	fmt.Println(MaximumEnemyFortsThatCanBeCaptured([]int{0, 0, 1, -1}))                 // 0
}

func MaximumEnemyFortsThatCanBeCaptured(forts []int) int {
	maxCap := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(forts); i++ {
		if forts[i] == 1 {
			// Move right
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == -1 {
					if j-i-1 > maxCap {
						maxCap = j - i - 1
					}
					break
				} else if forts[j] == 1 {
					break
				}
			}
		} else if forts[i] == -1 {
			// Move right (capturing from -1 to 1)
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == 1 {
					if j-i-1 > maxCap {
						maxCap = j - i - 1
					}
					break
				} else if forts[j] == -1 {
					break
				}
			}
		}
	}
	return maxCap
}
```

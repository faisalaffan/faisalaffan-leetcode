# 1487 — Making File Names Unique

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetFolderNames(names []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N) average, Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1487: Making File Names Unique
// https://leetcode.com/problems/making-file-names-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetFolderNames([]string{"pes", "fifa", "gta", "pes(2019)"}))
	fmt.Println(GetFolderNames([]string{"gta", "gta(1)", "gta", "avalon"}))
	fmt.Println(GetFolderNames([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece", "onepiece(1)"}))
}

func GetFolderNames(names []string) []string {
	// Time: O(N) average, Space: O(N)
  // Membuat map (HashMap) — pencarian O(1)
	used := make(map[string]int)
	result := make([]string, len(names))

	for i, name := range names {
		if _, exists := used[name]; !exists {
			used[name] = 1
			result[i] = name
			continue
		}

		k := used[name]
		candidate := name + "(" + itoa(k) + ")"
		for {
			if _, exists := used[candidate]; exists {
				k++
				candidate = name + "(" + itoa(k) + ")"
			} else {
				break
			}
		}
		used[name] = k + 1
		used[candidate] = 1
		result[i] = candidate
	}

	return result
}

// Simple int to string for positive ints
func itoa(n int) string {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[pos:])
}
```

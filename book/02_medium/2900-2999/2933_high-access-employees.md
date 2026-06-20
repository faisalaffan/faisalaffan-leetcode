# 2933 — High Access Employees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findHighAccessEmployees(accessTimes [][]string) (ans []string)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2933: High-Access Employees
// https://leetcode.com/problems/high-access-employees/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findHighAccessEmployees([][]string{{"a","0549"},{"b","0456"},{"a","1234"},{"a","1440"}}))
	fmt.Println(findHighAccessEmployees([][]string{{"d","0002"},{"d","0003"},{"d","0004"},{"d","0005"}}))
}

func findHighAccessEmployees(accessTimes [][]string) (ans []string) {
	d := map[string][]int{}
	for _, e := range accessTimes {
		name, s := e[0], e[1]
		h, _ := strconv.Atoi(s[:2])
		m, _ := strconv.Atoi(s[2:])
		t := h*60 + m
		d[name] = append(d[name], t)
	}
	for name, ts := range d {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(ts)
		for i := 2; i < len(ts); i++ {
			if ts[i]-ts[i-2] < 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	return
}
```

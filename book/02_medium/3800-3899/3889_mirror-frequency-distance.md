# 3889 — Mirror Frequency Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MirrorFrequencyDistance(s string) int
```

> **💡 Hint:** Count character frequencies. For each unique char, compute mirror

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3889: Mirror Frequency Distance
// https://leetcode.com/problems/mirror-frequency-distance/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count character frequencies. For each unique char, compute mirror
// char and sum absolute frequency differences, counting each pair once.

import "fmt"

func MirrorFrequencyDistance(s string) int {
  // Alokasi slice integer
	freq := make([]int, 36) // 0-25: letters, 26-35: digits
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			freq[ch-'a']++
		} else {
			freq[26+int(ch-'0')]++
		}
	}

	visited := make([]bool, 36)
	ans := 0
	for i := 0; i < 36; i++ {
		if visited[i] || freq[i] == 0 {
			continue
		}
		// Compute mirror
		var mirror int
		if i < 26 {
			mirror = 25 - i // 'a'->'z', 'b'->'y', etc.
		} else {
			mirror = 26 + (9 - (i - 26)) // '0'->'9', '1'->'8', etc.
		}
		visited[i] = true
		if mirror != i {
			visited[mirror] = true
		}
		diff := freq[i] - freq[mirror]
		if diff < 0 {
			diff = -diff
		}
		ans += diff
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MirrorFrequencyDistance("ab1z9")) // Expected: 3

	// Example 2
	fmt.Println(MirrorFrequencyDistance("4m7n")) // Expected: 2

	// Example 3
	fmt.Println(MirrorFrequencyDistance("byby")) // Expected: 0
}
```

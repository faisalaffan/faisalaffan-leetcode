# 3849 — Maximum Bitwise Xor After Rearrangement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumBitwiseXorAfterRearrangement(s string, t string) string
```

> **💡 Hint:** Count 0s and 1s in t. Greedily match opposite bits for max XOR.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3849: Maximum Bitwise XOR After Rearrangement
// https://leetcode.com/problems/maximum-bitwise-xor-after-rearrangement/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count 0s and 1s in t. Greedily match opposite bits for max XOR.

import "fmt"

func MaximumBitwiseXorAfterRearrangement(s string, t string) string {
	ones, zeros := 0, 0
	for _, ch := range t {
		if ch == '1' {
			ones++
		} else {
			zeros++
		}
	}

	ans := make([]byte, len(s))
	for i, ch := range s {
		if ch == '1' {
			if zeros > 0 {
				ans[i] = '1'
				zeros--
			} else {
				ans[i] = '0'
				ones--
			}
		} else {
			if ones > 0 {
				ans[i] = '1'
				ones--
			} else {
				ans[i] = '0'
				zeros--
			}
		}
	}

	return string(ans)
}

func main() {
	// Example 1
	fmt.Println(MaximumBitwiseXorAfterRearrangement("101", "011")) // Expected: "110"

	// Example 2
	fmt.Println(MaximumBitwiseXorAfterRearrangement("0110", "1110")) // Expected: "1101"

	// Example 3
	fmt.Println(MaximumBitwiseXorAfterRearrangement("0101", "1001")) // Expected: "1111"
}
```

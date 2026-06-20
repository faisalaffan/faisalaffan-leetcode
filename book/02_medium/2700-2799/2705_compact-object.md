# 2705 — Compact Object

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func compact(obj any) any
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2705: Compact Object
// https://leetcode.com/problems/compact-object/
// Difficulty: Medium (JS problem)
// Time: O(n) | Space: O(n)

import "fmt"

func compact(obj any) any {
	switch v := obj.(type) {
	case map[string]any:
  // Membuat map (HashMap) — pencarian O(1)
		result := make(map[string]any)
		for k, val := range v {
			compacted := compact(val)
			if compacted != nil {
				result[k] = compacted
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case []any:
		result := []any{}
		for _, val := range v {
			compacted := compact(val)
			if compacted != nil {
				result = append(result, compacted)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case bool:
		if !v {
			return nil
		}
	case int:
		if v == 0 {
			return nil
		}
	case string:
		if v == "" {
			return nil
		}
	}
	return obj
}

func main() {
	// Test case 1
	input1 := map[string]any{
		"a": nil,
		"b": 1,
		"c": 0,
		"d": "",
		"e": false,
		"f": []any{nil, 1, 0, "", false},
	}
	fmt.Println("Test 1:", compact(input1))

	// Test case 2
	input2 := []any{0, 1, false, true, "", "hello"}
	fmt.Println("Test 2:", compact(input2))

	// Test case 3: empty result
	input3 := map[string]any{"a": false, "b": 0, "c": ""}
	fmt.Println("Test 3:", compact(input3))
}
```

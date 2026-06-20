# 0067 — Add Binary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func AddBinary(a string, b string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(max(n,m))  
**Kompleksitas Ruang:** O(max(n,m))

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #67: Add Binary
// https://leetcode.com/problems/add-binary/
// Difficulty: Easy

import "fmt"

// Time: O(max(n,m)) | Space: O(max(n,m))
func AddBinary(a string, b string) string {
	i, j, carry := len(a)-1, len(b)-1, 0
	res := make([]byte, 0, max(len(a), len(b))+1)
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}
		res = append(res, byte('0'+carry%2))
		carry /= 2
	}
	// reverse
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

func main() {
	fmt.Println(AddBinary("11", "1"))
	fmt.Println(AddBinary("1010", "1011"))
}
```

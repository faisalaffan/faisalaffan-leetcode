# 0751 — Ip To Cidr

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ipToCIDR(ip string, n int) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n) for each block  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #751: IP to CIDR
// https://leetcode.com/problems/ip-to-cidr/
// Difficulty: Medium [Paid]
// Time: O(log n) for each block
// Space: O(1)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := ipToCIDR("255.0.0.7", 10)
	for _, s := range result {
		fmt.Println(s)
	}
}

func ipToCIDR(ip string, n int) []string {
	start := ipToInt(ip)
	result := make([]string, 0)

	for n > 0 {
		mask := max(0, start&-start)
		for mask > n {
			mask >>= 1
		}

		result = append(result, intToIP(start)+"/"+strconv.Itoa(32-trailingZeros(mask)))
		start += mask
		n -= mask
	}

	return result
}

func ipToInt(ip string) int {
	parts := strings.Split(ip, ".")
	result := 0
	for _, p := range parts {
		val, _ := strconv.Atoi(p)
		result = result*256 + val
	}
	return result
}

func intToIP(n int) string {
	return fmt.Sprintf("%d.%d.%d.%d", n>>24, (n>>16)&255, (n>>8)&255, n&255)
}

func trailingZeros(x int) int {
	if x == 0 {
		return 32
	}
	count := 0
	for x&1 == 0 {
		x >>= 1
		count++
	}
	return count
}
```

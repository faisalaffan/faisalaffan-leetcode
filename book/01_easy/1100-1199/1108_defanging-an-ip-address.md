# 1108 — Defanging An Ip Address

## Deskripsi

**Soal:** [1108. Defanging An Ip Address](https://leetcode.com/problems/defanging-an-ip-address/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1108: Defanging an IP Address
// https://leetcode.com/problems/defanging-an-ip-address/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))       // "1[.]1[.]1[.]1"
	fmt.Println(defangIPaddr("255.100.50.0"))  // "255[.]100[.]50[.]0"
}

// LeetCode submission: defangIPaddr
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
```

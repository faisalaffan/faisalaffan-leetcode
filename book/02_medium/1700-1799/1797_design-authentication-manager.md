# 1797 — Design Authentication Manager

## Deskripsi

**Soal:** [1797. Design Authentication Manager](https://leetcode.com/problems/design-authentication-manager/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) per operation, Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(timeToLive int) AuthenticationManager`

## Solusi Go

```go
package main

// LeetCode #1797: Design Authentication Manager
// https://leetcode.com/problems/design-authentication-manager/
// Difficulty: Medium
// Time: O(n) per operation, Space: O(n)

import (
	"fmt"
)

type AuthenticationManager struct {
	ttl     int
	tokens  map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		ttl:    timeToLive,
		tokens: make(map[string]int),
	}
}

func (am *AuthenticationManager) Generate(tokenId string, currentTime int) {
	am.tokens[tokenId] = currentTime + am.ttl
}

func (am *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if expiry, ok := am.tokens[tokenId]; ok && expiry > currentTime {
		am.tokens[tokenId] = currentTime + am.ttl
	}
}

func (am *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for _, expiry := range am.tokens {
		if expiry > currentTime {
			count++
		}
	}
	return count
}

func main() {
	am := Constructor(5)
	am.Generate("aaa", 1)
	fmt.Println(am.CountUnexpiredTokens(2)) // Expected: 1
	am.Renew("aaa", 3)
	fmt.Println(am.CountUnexpiredTokens(6)) // Expected: 0
	am.Generate("bbb", 7)
	am.Generate("ccc", 8)
	fmt.Println(am.CountUnexpiredTokens(10)) // Expected: 2
	am.Renew("bbb", 11)
	fmt.Println(am.CountUnexpiredTokens(12)) // Expected: 1
}
```

# 0415 — Add Strings

## Deskripsi

**Soal:** [0415. Add Strings](https://leetcode.com/problems/add-strings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m), Space: O(max(n,m))  
**Kompleksitas Ruang:** O(max(n,m))

**Algoritma:** —

**Fungsi Solusi:** `func AddStrings(num1, num2 string) string`

## Solusi Go

```go
package main

// LeetCode #415: Add Strings
// https://leetcode.com/problems/add-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(max(n,m))
func AddStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var result []byte
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		carry = sum / 10
		result = append([]byte{byte(sum%10 + '0')}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println(AddStrings("11", "123"))
	fmt.Println(AddStrings("456", "77"))
	fmt.Println(AddStrings("0", "0"))
}
```

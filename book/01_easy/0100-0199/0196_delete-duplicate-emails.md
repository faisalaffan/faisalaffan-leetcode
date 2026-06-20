# 0196 — Delete Duplicate Emails

## Deskripsi

**Soal:** [0196. Delete Duplicate Emails](https://leetcode.com/problems/delete-duplicate-emails/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func DeleteDuplicateEmails() string`

## Solusi Go

```go
package main

// LeetCode #196: Delete Duplicate Emails
// https://leetcode.com/problems/delete-duplicate-emails/
// Difficulty: Easy

import "fmt"

func DeleteDuplicateEmails() string {
	return "DELETE p1 FROM Person p1, Person p2 WHERE p1.email = p2.email AND p1.id > p2.id"
}

func main() {
	fmt.Println(DeleteDuplicateEmails())
}
```

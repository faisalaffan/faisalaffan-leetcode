# 2704 — To Be Or Not To Be

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ToBeOrNotToBe(val interface{}) *Expect`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2704: To Be Or Not To Be
// https://leetcode.com/problems/to-be-or-not-to-be/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Expect-type assertion.

import (
	"fmt"
	"reflect"
)

func main() {
	{
		expect := ToBeOrNotToBe(5)
		result, err := expect.toBe(5)
		fmt.Println(result, err)
	}
	{
		expect := ToBeOrNotToBe(5)
		result, err := expect.notToBe(5)
		fmt.Println(result, err)
	}
}

type Expect struct {
	val interface{}
}

func (e *Expect) toBe(expected interface{}) (bool, string) {
	if reflect.DeepEqual(e.val, expected) {
		return true, ""
	}
	return false, "Not Equal"
}

func (e *Expect) notToBe(expected interface{}) (bool, string) {
	if !reflect.DeepEqual(e.val, expected) {
		return true, ""
	}
	return false, "Equal"
}

func ToBeOrNotToBe(val interface{}) *Expect {
	return &Expect{val: val}
}
```

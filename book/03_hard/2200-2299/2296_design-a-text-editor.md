# 2296 — Design A Text Editor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** —

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor() TextEditor`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// 2296. Design a Text Editor
// ----------------------------------------------------------------
// Two-stack approach.  left holds characters before the cursor (top = rightmost),
// right holds characters after the cursor (reversed so right.top is the first
// character after the cursor).

type TextEditor struct {
	left  []byte
	right []byte
}

func Constructor() TextEditor {
	return TextEditor{
		left:  make([]byte, 0, 512),
		right: make([]byte, 0, 512),
	}
}

// addText inserts text at the cursor position.
func (t *TextEditor) addText(text string) string {
	t.left = append(t.left, []byte(text)...)
	return t.peek()
}

// deleteText deletes the k characters immediately before the cursor.
// Returns the string of the last min(10, len(left)) characters after deletion.
func (t *TextEditor) deleteText(k int) string {
	if k > len(t.left) {
		k = len(t.left)
	}
	t.left = t.left[:len(t.left)-k]
	return t.peek()
}

// cursorLeft moves the cursor left by k characters (or to the start).
func (t *TextEditor) cursorLeft(k int) string {
	if k > len(t.left) {
		k = len(t.left)
	}
	// Move k chars from left to right.
	// left's top k chars go to right in reversed order so the cursor stays
	// correctly between left and right.
	moved := t.left[len(t.left)-k:]
	t.left = t.left[:len(t.left)-k]
	for i := len(moved) - 1; i >= 0; i-- {
		t.right = append(t.right, moved[i])
	}
	return t.peek()
}

// cursorRight moves the cursor right by k characters (or to the end).
func (t *TextEditor) cursorRight(k int) string {
	if k > len(t.right) {
		k = len(t.right)
	}
	moved := t.right[len(t.right)-k:]
	t.right = t.right[:len(t.right)-k]
	for i := len(moved) - 1; i >= 0; i-- {
		t.left = append(t.left, moved[i])
	}
	return t.peek()
}

// peek returns the last min(10, len(left)) characters of left.
func (t *TextEditor) peek() string {
	n := len(t.left)
	start := n - 10
	if start < 0 {
		start = 0
	}
	return string(t.left[start:])
}

// ---------------------------------------------------------------------------
//  Wrapper (returns last peek value)

func DesignATextEditor() interface{} {
	editor := Constructor()
	editor.addText("leetcode")
	editor.deleteText(4)
	editor.addText("practice")
	editor.cursorRight(3)
	editor.cursorLeft(8)
	editor.deleteText(10)
	editor.cursorLeft(2)
	return editor.cursorRight(6)
}

func main() {
	fmt.Println(DesignATextEditor())

	editor := Constructor()
	if s := editor.addText("hello"); s != "hello" {
		fmt.Printf("FAIL addText hello: got %q\n", s)
	}
	if s := editor.cursorLeft(2); s != "hel" {
		fmt.Printf("FAIL cursorLeft 2: got %q\n", s)
	}
	if s := editor.cursorRight(2); s != "hello" {
		fmt.Printf("FAIL cursorRight 2: got %q\n", s)
	}
	if s := editor.deleteText(2); s != "hel" {
		fmt.Printf("FAIL deleteText 2: got %q\n", s)
	}
	if s := editor.addText("p"); s != "help" {
		fmt.Printf("FAIL addText p: got %q\n", s)
	}
	fmt.Println("Done testing 2296.")
}
```

package main

// LeetCode #2296: Design a Text Editor
// https://leetcode.com/problems/design-a-text-editor/
// Difficulty: Hard
//
// Approach: Two stacks (slices). Left stack holds characters before cursor,
// right stack holds characters after cursor (reversed). O(1) per operation
// amortized. cursorLeft/cursorRight move at most k characters between stacks.

import "fmt"

func main() {
	// Example from problem
	editor := Constructor()
	fmt.Println(editor.addText("leetcode"))
	fmt.Println(editor.deleteText(4))
	fmt.Println(editor.addText("practice"))
	fmt.Println(editor.cursorRight(3))
	fmt.Println(editor.cursorLeft(8))
	fmt.Println(editor.deleteText(10))
	fmt.Println(editor.cursorLeft(2))
	fmt.Println(editor.cursorRight(6))
}

type TextEditor struct {
	left  []byte
	right []byte
}

func Constructor() TextEditor {
	return TextEditor{
		left:  make([]byte, 0, 1024),
		right: make([]byte, 0, 1024),
	}
}

func (t *TextEditor) addText(text string) string {
	t.left = append(t.left, []byte(text)...)
	return t.peek()
}

func (t *TextEditor) deleteText(k int) string {
	if k > len(t.left) {
		k = len(t.left)
	}
	t.left = t.left[:len(t.left)-k]
	return t.peek()
}

func (t *TextEditor) cursorLeft(k int) string {
	if k > len(t.left) {
		k = len(t.left)
	}
	// Move k chars from left to right
	moved := t.left[len(t.left)-k:]
	t.left = t.left[:len(t.left)-k]
	// Push to right in reverse order (so rightmost char is at end)
	for i := len(moved) - 1; i >= 0; i-- {
		t.right = append(t.right, moved[i])
	}
	return t.peek()
}

func (t *TextEditor) cursorRight(k int) string {
	if k > len(t.right) {
		k = len(t.right)
	}
	// Move k chars from right back to left
	moved := t.right[len(t.right)-k:]
	t.right = t.right[:len(t.right)-k]
	for i := len(moved) - 1; i >= 0; i-- {
		t.left = append(t.left, moved[i])
	}
	return t.peek()
}

// peek returns the last min(10, len(left)) chars of left
func (t *TextEditor) peek() string {
	n := len(t.left)
	start := n - 10
	if start < 0 {
		start = 0
	}
	return string(t.left[start:])
}

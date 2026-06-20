# 0428 — Serialize And Deserialize N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// LeetCode #428: Serialize and Deserialize N-ary Tree
// https://leetcode.com/problems/serialize-and-deserialize-n-ary-tree/
// Difficulty: Hard
//
// BFS with child count. Serialize format: "val childCount val childCount ..."
// Deserialize uses a queue to rebuild the tree.

func main() {
	// Build tree: root = [1,null,3,2,4,null,5,6]
	// 1 has children 3,2,4; 3 has children 5,6
	root := &Node{Val: 1}
	n3 := &Node{Val: 3}
	n2 := &Node{Val: 2}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	root.Children = []*Node{n3, n2, n4}
	n3.Children = []*Node{n5, n6}

	codec := &Codec{}
	serialized := codec.serialize(root)
	fmt.Println("Serialized:", serialized)
	deserialized := codec.deserialize(serialized)
	fmt.Println("Root val:", deserialized.Val)
	fmt.Println("Children count:", len(deserialized.Children))

	// Empty tree
	empty := codec.serialize(nil)
	fmt.Println("Empty:", empty)
	fmt.Println("Empty deserialized:", codec.deserialize(empty))
}

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Codec handles serialization/deserialization.
type Codec struct{}

func (c *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}
	var parts []string
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		parts = append(parts, strconv.Itoa(node.Val))
		parts = append(parts, strconv.Itoa(len(node.Children)))
		queue = append(queue, node.Children...)
	}
	return strings.Join(parts, " ")
}

func (c *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	tokens := strings.Split(data, " ")
	if len(tokens) == 0 {
		return nil
	}
	val, _ := strconv.Atoi(tokens[0])
	childCount, _ := strconv.Atoi(tokens[1])
	root := &Node{Val: val, Children: make([]*Node, 0, childCount)}

	type frame struct {
		node *Node
		need int
	}
	queue := []*frame{{node: root, need: childCount}}
	idx := 2
	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]
		for i := 0; i < f.need; i++ {
			val, _ = strconv.Atoi(tokens[idx])
			childCount, _ = strconv.Atoi(tokens[idx+1])
			idx += 2
			child := &Node{Val: val, Children: make([]*Node, 0, childCount)}
			f.node.Children = append(f.node.Children, child)
			if childCount > 0 {
				queue = append(queue, &frame{node: child, need: childCount})
			}
		}
	}
	return root
}
```

package main

// LeetCode #588: Design In-Memory File System
// https://leetcode.com/problems/design-in-memory-file-system/
// Difficulty: Hard

import (
	"fmt"
	"sort"
	"strings"
)

type FileSystem struct {
	root *TrieNode
}

type TrieNode struct {
	name     string
	isFile   bool
	content  string
	children map[string]*TrieNode
}

func NewFileSystem() FileSystem {
	return FileSystem{
		root: &TrieNode{
			children: make(map[string]*TrieNode),
		},
	}
}

func (fs *FileSystem) ls(path string) []string {
	node := fs.traverse(path)
	if node.isFile {
		return []string{node.name}
	}
	names := []string{}
	for name := range node.children {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (fs *FileSystem) mkdir(path string) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 1 && parts[0] == "" {
		return
	}
	node := fs.root
	for _, part := range parts {
		if part == "" {
			continue
		}
		if _, ok := node.children[part]; !ok {
			node.children[part] = &TrieNode{
				name:     part,
				children: make(map[string]*TrieNode),
			}
		}
		node = node.children[part]
	}
}

func (fs *FileSystem) addContentToFile(filePath, content string) {
	node := fs.traverse(filePath)
	if node.isFile {
		node.content += content
	} else {
		// Need to create the file
		parts := strings.Split(strings.Trim(filePath, "/"), "/")
		dirPath := "/" + strings.Join(parts[:len(parts)-1], "/")

		fs.mkdir(dirPath)
		node = fs.traverse(filePath)
		node.isFile = true
		node.content = content
	}
}

func (fs *FileSystem) readContentFromFile(filePath string) string {
	node := fs.traverse(filePath)
	return node.content
}

func (fs *FileSystem) traverse(path string) *TrieNode {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	node := fs.root
	for _, part := range parts {
		if part == "" {
			break
		}
		if child, ok := node.children[part]; ok {
			node = child
		} else {
			child := &TrieNode{
				name:     part,
				children: make(map[string]*TrieNode),
			}
			node.children[part] = child
			node = child
		}
	}
	return node
}

func main() {
	fs := NewFileSystem()

	fmt.Println(fs.ls("/")) // []

	fs.mkdir("/a/b/c")
	fs.addContentToFile("/a/b/c/d", "hello")
	fmt.Println(fs.ls("/"))          // ["a"]
	fmt.Println(fs.readContentFromFile("/a/b/c/d")) // "hello"

	fs.addContentToFile("/a/b/c/d", " world")
	fmt.Println(fs.readContentFromFile("/a/b/c/d")) // "hello world"
}

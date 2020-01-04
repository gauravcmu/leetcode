package main

import (
"fmt"
)

func main() {
	fs := Constructor()

	b := fs.CreatePath("/leet", 1)
	if b != true {
		fmt.Printf("failed to create path\n")
	}
	n := fs.Get("/leet")
	fmt.Printf("got: %d\n", n)

	b = fs.CreatePath("/leet/code", 11)
	if b != true {
		fmt.Printf("failed to create path\n")
	}
	n = fs.Get("/leet/code")
	fmt.Printf("got: %d\n", n)

	//negative test.
	n = fs.Get("/leets/code")
	fmt.Printf("got: %d\n", n)

	//negative test.
	b = fs.CreatePath("/leets/code", 11)
	if b != true {
		fmt.Printf("failed to create path\n")
	}
}

type FileSystem struct {
	children map[byte]*FileSystem
	isValue  bool
	value    int
}

func Constructor() FileSystem {
	return FileSystem{
		children: make(map[byte]*FileSystem, 0),
		isValue:  false,
	}
}

func createNode() *FileSystem {
	return &FileSystem{
		children: make(map[byte]*FileSystem, 0),
		isValue:  false,
	}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	current := this
	var old *FileSystem

	fmt.Printf("CreatePath(): %s %d\n", path, value)
	for i, b := range path {
		if b == '/' && i == 0 {
			//first slash, ignore.
			continue
		}
		if b == '/' && old.isValue != true {
			fmt.Printf("CreatePath(): not found path, return false %s %d\n", path, value)
			//the old should point to a value..
			return false
		}
		if _, ok := current.children[byte(b)]; ok {
			current = current.children[byte(b)]
		} else {
			fmt.Printf("CreatePath(): adding node %s %d\n", path, value)
			current.children[byte(b)] = createNode()
			current = current.children[byte(b)]
		}

		old = current
	}
	current.isValue = true
	current.value = value
	return true
}

func (this *FileSystem) Get(path string) int {
	current := this

	var old *FileSystem

	for i, b := range path {
		if b == '/' && i == 0 {
			continue
		}
		if _, ok := current.children[byte(b)]; ok {
			current = current.children[byte(b)]
		} else if b != '/' && !ok {
			fmt.Printf("GetPath(): not found path in children, return -1 %s\n", path)
			return -1
		}
		if b == '/' && old.isValue != true {
			//the old should point to a value..
			fmt.Printf("GetPath(): not found path, return -1 %s\n", path)
			return -1
		}
		old = current
	}
	return current.value
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */


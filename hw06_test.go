package main

import (
	"testing"
)

func TestMainHW06(t *testing.T) {
	hw06()
}

func TestMainDoubleLinkedList(t *testing.T) {

	//check nil constructor
	strList := DoubleLinkedList{} // all nil

	if strList.Len() != 0 {
		t.Fatalf("Size is not empty")
	}

	if strList.Head() != nil {
		t.Fatalf("Head is not nil")
	}

	if strList.Last() != nil {
		t.Fatalf("Last is not nil")
	}

	// prepend
	strList.PushFront("Varvara")

	if strList.Len() != 1 {
		t.Fatalf("Size is not equal")
	}

	if strList.Head() == nil {
		t.Fatalf("Head is nil")
	}

	// append
	strList.PushBack("Knopa")

	if strList.Len() != 2 {
		t.Fatalf("Size is not equal")
	}

	if strList.Last() == nil {
		t.Fatalf("Last is nil")
	}

	// get elements by indexes
	if strList.Item(0).Value() != "Varvara" {
		t.Fatalf("Item returns wrong element")
	}

	if strList.Item(1).Value() != "Knopa" {
		t.Fatalf("Item returns wrong element")
	}

	// Insert element
	strList.Insert(0, "Simka")

	if strList.Len() != 3 {
		t.Fatalf("Size is not equal")
	}

	if strList.Item(0).Value() != "Varvara" {
		t.Fatalf("Item returns wrong element")
	}

	if strList.Item(1).Value() != "Simka" {
		t.Fatalf("Item returns wrong element")
	}

	if strList.Item(2).Value() != "Knopa" {
		t.Fatalf("Item returns wrong element")
	}

	// Remove element
	strList.Remove(1)

	// get elements by indexes
	if strList.Item(0).Value() != "Varvara" {
		t.Fatalf("Item returns wrong element")
	}

	if strList.Item(1).Value() != "Knopa" {
		t.Fatalf("Item returns wrong element")
	}
}
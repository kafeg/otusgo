package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

// --- samples

type HW06User struct { // structure with named fields
	Id int64
	Name string
	Age int
	friends []int64 // private element
}

func (u HW06User) IsOk() bool {
	for _, fid := range u.friends {
		if u.Id == fid {
			return true
		}
	}
	return false
}

func (u HW06User) HappyBirthday1() {
	u.Age++ // value will be lost, because it's a copy of structure in this method
}

func (u *HW06User) HappyBirthday2() {
	u.Age++ // OK, because we have pointer to the structure
}

// --- end of samples

func hw06() {

	//struct{} // empty structure, does not occupy memory

	{
		u1 := HW06User{}               // Zero Value for type HW06User
		u2 := &HW06User{}              // The same but it's a pointer
		u3 := HW06User{1, "Vasya", 23, []int64{1, 2, 3}} // init by fields number
		u4 := HW06User{ // init by fields names
			Id:      1,
			Name:    "Vasya",
			Age:     23,
			friends: []int64{1, 2, 3},
		}

		fmt.Printf("%v %v %v %v", u1, u2, u3, u4)
		fmt.Println(u3.IsOk())
	}

	// anon struncts
	{
		var wordCounts []struct{w string; n int}
		fmt.Println(wordCounts)

		//structure with special tags which can be used by the external library
		var resp struct {
			Ok bool `json:"ok"`
			Total int `json:"total"`
				Documents []struct{
				Id int `json:"id"`
				Title string `json:"title"`
			} `json:"documents"`
		}

		respTp := reflect.TypeOf(resp)
		okField, _ := respTp.FieldByName("Ok")

		data := `{
					"ok": true,
					"total": 1,
					"documents": [{
						"id": 2,
						"title": "Maw"
					}]
				}`
		json.Unmarshal([]byte(data), &resp)
		fmt.Println(resp.Documents[0].Title)
		fmt.Println(okField)
	}

	// get address
	{
		var x struct {
			a int
			b string
			c [10]rune
		}
		bPtr := &x.b
		c3Ptr := &x.c[2]

		fmt.Printf("%v %v", bPtr, c3Ptr)
	}

	// pointers
	{
		//a := "qwe" // string
		//aPtr := &a // *string
		//b := *aPtr // string, value "qwe"
		//var n *int // nil
		//nv := *n // panic
	}

	//copy structure
	{
		a := struct{x, y int}{0, 0}
		b := a
		a.x = 1
		fmt.Println(b.x) // 0, because all data copied
	}

	// copy pointer to structure
	{
		a := new(struct{x, y int})
		b := a
		a.x = 1
		fmt.Println(b.x) // 1, because we copied only structure pointer
	}

	//inherited methods form the inner structures
	{
		type LinkStorage struct {
			sync.Mutex // только тип!
			storage map[string]string // тип и имя
		}

		storage := LinkStorage{}
		storage.Mutex.Lock() // имя типа используется
		storage.Mutex.Unlock() // как имя элемента структуры

		// instead of
		storage.Mutex.Lock()
		storage.Unlock()
		// we can call
		storage.Lock()
		storage.Unlock()
	}
}

// --- homework

type DoubleLinkedItem struct {
	value interface{}
	next *DoubleLinkedItem
	prev *DoubleLinkedItem
}

func (l *DoubleLinkedItem) Value() interface{} {
	return l.value
}

func (l *DoubleLinkedItem) Next() *DoubleLinkedItem {
	return l.next
}

func (l *DoubleLinkedItem) Prev() *DoubleLinkedItem {
	return l.prev
}

type DoubleLinkedList struct {
	size int // private, accessible by Len()
	head *DoubleLinkedItem
	last *DoubleLinkedItem
}

func (l *DoubleLinkedList) Len() (int) {
	return l.size
}

func (l *DoubleLinkedList) Head() (*DoubleLinkedItem) {
	return l.head
}

func (l *DoubleLinkedList) Last() (*DoubleLinkedItem) {
	return l.last
}

func (l *DoubleLinkedList) PushFront(v interface{}) {
	newItem := new(DoubleLinkedItem)
	newItem.value = v
	newItem.next = l.head
	if l.head != nil {
		l.head.prev = newItem
	}
	l.head = newItem

	if l.size == 0 {
		l.last = newItem
	}

	l.size++
}

func (l *DoubleLinkedList) PushBack(v interface{}) {
	newItem := new(DoubleLinkedItem)
	newItem.value = v
	newItem.prev = l.last
	if l.last != nil {
		l.last.next = newItem
	}
	l.last = newItem

	if l.size == 0 {
		l.head = newItem
	}

	l.size++
}

func (l *DoubleLinkedList) Item(index int) (*DoubleLinkedItem) {

	//fmt.Printf("%v %v\n", index, l.size)

	if index >= l.size || index < 0 {
		panic("Index outside of bounds")
	}

	item := l.head
	for i := 0; i < index; i++ {
		//fmt.Printf("%v %v %v %v\n", i, index, l.size, item)
		item = item.next
	}

	return item
}

func (l *DoubleLinkedList) Insert(i int, v interface{}) {

	if i == -1 {
		l.PushFront(v)
		return
	}

	if i == l.size {
		l.PushBack(v)
		return
	}

	// insert value after especial element
	newItem := new(DoubleLinkedItem)
	newItem.value = v

	currItem := l.Item(i)

	newItem.prev = currItem

	if currItem.next != nil {
		currItem.next.prev = newItem
		newItem.next = currItem.next
	}
	currItem.next = newItem

	l.size++
}

func (l *DoubleLinkedList) Remove(i int) {
	// insert value after especial element

	currItem := l.Item(i)

	currItem.prev.next = currItem.next
	currItem.next.prev = currItem.prev
	currItem = nil

	l.size--
}

// --- end of homework

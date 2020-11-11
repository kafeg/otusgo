package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"unsafe"
)

func hw04() {

	//arrays, variants of creation
	{
		var arr1 [256]int       // fix length, can't be enlarged without copy
		var arr2 [10][10]string // can be multidimensional
		var arr3 = [...]int{ 1, 2, 3} // [1 2 3]
		arr4 := [10]int{1, 2, 3, 4, 5} // [1 2 3 4 5 0 0 0 0 0]
		arr4[5] = 6 // [1 2 3 4 5 6 0 0 0 0]

		slice1 := arr4[3:5] // [4 5]
		slice2 := arr4[:5] // [1 2 3 4 5]
		slice3 := arr4[3:] // [4 5 6 0 0 0 0]

		fmt.Printf("%v %v %v %v %v %v %v %v\n", arr1, arr2, arr3, arr4, len(arr4), slice1, slice2, slice3)
	}

	//slices - like arrays, but with non-fixed length and have capacity
	{
		var s1 []int   // not initialized slice, nil
		s2 := []int{}  // empty slice
		s3 := make([]int, 3) // with func make, s == {0,0,0} // length 3, capacity 10
		s4 := make([]int, 3, 10) // length 3, capacity 10

		fmt.Printf("%v %v %v %v %v %v %v %v %v %v %v %v\n", s1, s2, s3, s4, len(s1), len(s2), len(s3), len(s4), cap(s1), cap(s2), cap(s3), cap(s4))
	}

	//slices operations
	{
		s := make([]int, 3)
		i := 2
		s[i] = 1                 // if i < len(s)
		//s[len(s) + 10] = 1       // will panic, because overflow
		s = append(s, 1)       // add 1 to the end of slice
		s = append(s, 1, 2, 3) // add 1, 2, 3 to the end of the slice
		var s3 []int               // s == nil
		s3 = append(s3, 1)  // s == {1}  append could work with nil-slices
		s = append(s, s3...)       // add slice s2 to the s
	}

	//slices operations
	{
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s2 := s[:]    // copy of s (shallow)
		s3 := s[3:5]  // []int{3,4}
		s4 := s[3:]   // []int{3, 4, 5, 6, 7, 8, 9}
		s5 := s[:5]   // []int{0, 1, 2, 3, 4}

		//to solve 'unused'
		s = s2
		s = s3
		s = s4
		s = s5
	}

	{
		// runtime/slice.go
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}

		s := make([]int, 3, 10)
		l := len(s)  // length of slice
		c := cap(s)  // capacity of slice

		//to solve 'unused'
		l = c
		c = l
	}

	//capacity increased when we add something to the slice
	{
		arr := []int{1}
		for i := 0; i < 100; i++ {
			fmt.Printf("len: %d \tcap %d  \tptr %0x\n",
				len(arr), cap(arr), &arr[0])
			arr = append(arr, i)
		}
	}

	//Copy on Write and Copy on Re-allocation
	{
		//sub-slices and slices copies targets on the same 'array unsafe.Pointer' until one of them not re-allocated

		arr := []int{1, 2}
		arr2 := arr
		arr2[0] = 42
		fmt.Println(arr[0]) // 42 - because both slices have shared pointer to the 'array unsafe.Pointer'


		arr2 = append(arr2, 3, 4, 5, 6, 7, 8, 9, 0)
		arr2[0] = 1
		fmt.Println(arr[0]) // 42 - because not related to arr2 now (arr2 was reallocated and poited to the new memory)
	}

	//real copy of slice
	{
		//functions like this should return new slice
		//func AppendUniq(slice []int, slice2 []int) []int {
		//...
		//}
		//s = AppendUniq(s, s2)

		s := []int{1,2,3}
		s2 := make([]int, len(s))
		copy(s2, s) // copy slice with re-allocation
	}

	//sort slices
	{
		// no generic-functios, so need to implement sort function to each type
		s1 := []int{3, 2, 1}
		sort.Ints(s1)

		s2 := []string{"hello", "cruel", "world"}
		sort.Strings(s2)

		// sort own types

		type User struct {
			Name     string  // public member, be available for e.g. in json.Marshal
			Age int  // private member
		}

		s3 := []User{
			{"vasya", 19},
			{"petya", 18},
		}
		sort.Slice(s3, func(i, j int) bool {
			return s3[i].Age < s3[j].Age
		})
	}

	//maps, dicts
	{
		var cache1 map[string]string     // nil
		cache2 := map[string]string{}    // empty, len(cache) == 0
		cache3 := map[string]string{     // map with init values
			"one":   "один",
			"two":   "два",
			"three": "три",
		}
		cache4 := make(map[string]string)  // the same to map[int]string{}
		cache5 := make(map[string]string, 100)  // map with pre-allocation for 100 key/values

		key := "two"
		value1 := cache3[key]     // get value, Zero Value if key not found
		value2, ok1 := cache3[key] // get value and result (ok if key found)
		_, ok2 := cache3[key]     // key exists?
		cache3[key] = value1      // set value for initialized map
		delete(cache3, key)      // delete key/val from the map

		fmt.Printf("%v %v %v %v %v, %v, %v, %v\n", cache1, cache2, cache3, cache4, cache5, ok1, ok2, value2)
	}

	//maps iterations
	{
		cache := map[string]int{     // map with init values
			"one":   1,
			"two":   2,
			"three": 3,
		}

		cache["three"]++ //update inplace

		for key, val := range cache {
			fmt.Printf("%v %v\n", key, val)
		}
		for key, _ := range cache {  // keys only
			fmt.Printf("%v\n", key)
		}
		for _, val := range cache {  // vals only
			fmt.Printf("%v\n", val)
		}

		//keys list
		var keys []string
		for key, _ := range cache {
			keys = append(keys, key)
		}

		//vals list
		values := make([]int, 0, len(cache))
		for _, val := range cache {
			values = append(values, val)
		}
	}

	//zero-values in maps/slices
	{
		var seq []string             // nil

		l1 := len(seq)       // 0
		c := cap(seq)       // 0
		seq = append(seq, "hello")   // []string{"hello"}

		var cache map[string]string  // nil
		l2 := len(cache)     // 0
		v, ok := cache["not_exists_key"] // "", false

		fmt.Printf("%v %v %v %v %v\n", l1, c, l2, v, ok)
	}

	// non thread save, so need to be safe by the mutex
	{
		sharedCache := map[string]string{}
		var cacheMutex sync.RWMutex
		get := func(key string) string {
			var s string
			cacheMutex.RLock()
			s = sharedCache[key]
			cacheMutex.RUnlock()
			return s
		}
		set := func(key string, val string) {
			cacheMutex.Lock()
			sharedCache[key] = val
			cacheMutex.Unlock()
		}

		set("olo", "lo")
		get("olo")

		fmt.Printf("%v\n", get("olo"))
	}
}

func rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}

func wordsCount(text string) map[string]int {
	wordCounters := map[string]int{}

	stringList := strings.Split(text, " ")

	for _, word := range stringList {

		//prepare word
		word = strings.ReplaceAll(word, ".", "")
		word = strings.ReplaceAll(word, string('\n'), "")
		word = strings.ToLower(word)

		//skip articles
		if len(word) <= 3 {
			continue
		}

		_, ok := wordCounters[word]
		if ok {
			wordCounters[word]++
		} else {
			wordCounters[word] = 1
		}
	}

	ranked := rankMapStringInt(wordCounters)
	ranked = ranked[0:10]
	result :=  map[string]int{}
	for _, index := range ranked {
		result[index] =  wordCounters[index]
		//fmt.Printf("%3d: %s -> %d\n", i, index, wordCounters[index])
	}

	return result
}

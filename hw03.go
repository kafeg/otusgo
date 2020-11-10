package hw03

import "fmt"

var Storage map[string]string         // zero value
var storage = make(map[string]string) // auto-type ('var' required outside functions, unable to use ':=' here)

type User struct {
	Name     string  // public member, be available for e.g. in json.Marshal
	password string  // private member
}

func Answer() int { // should return int
	return 42 // and returning int
}

func itoa(i int) (s string) {
	negative := i < 0
	if negative {
		i = 0 - i
	}

	if i == 0 {
		return "0"
	}

	for i > 0 {
		tmp := i % 10
		i = i / 10
		//fmt.Printf("tmp %v\n", tmp)
		s = fmt.Sprint(tmp) + s
	}
	if negative {
		s = "-" + s
	}

	return s
}

func hw03() {
	{
		var i int = 10
		j := i // short definition, only in functions body
		i = j
	}

	//types
	{
		var a int; // zero-value = 0
		var b uint = 1
		var c int8 = 1
		var d uint8 = 1
		var e int16 = int16(d) // cast type, 'e = d' - is error!
		var f uint16 = 1
		var g byte = 1 // = uint8
		var h rune = 1 // = int32
		var i float32 = 1
		var j float64 = 1
		var k string = "test" // zero value = ""
		// Pointers:  uintptr , *int , *string, ...
		// complex, ...

		//print types
		fmt.Printf("%T %T %T %T %T %T %T %T %T %T %T\n", a, b, c, d, e, f, g, h, i, j, k)
	}

	// strings
	{
		// string - is a const byte array (bytes, not runes - important!!!)
		s := "hello world"       // create string
		var c byte = s[0]       // get byte(!) in string
		var s2 string = s[5:10]  // get substring (in bytes!)
		s2 = s + " again"       // concat strings, here copy
		l := len(s)              // string length in bytes

		fmt.Printf("%v %v %d %v %T\n", c, s2, l, l, l)
	}

	// unicode
	{
		var r1 rune = 'Я' // rune - unicode bytes array
		var r2 rune = '\n'
		var r3 rune = '本'
		var r4 rune = '\xff'   // some bytes
		var r5 rune = '\u12e4' // unicode code-point
		fmt.Printf("%v %v %v %v %v\n", r1, r2, r3, r4, r5)

		// get first rune from the string and its lingth in bytes
		// DecodeRuneInString(s string) (r rune, size int)

		// get string length in bytes
		// RuneCountInString(s string) (n int)

		// check is string valid
		// ValidString(s string) bool

		//convert to slice
		s := "привет"
		ba := []byte(s)
		ra := []rune(s)
		fmt.Printf("% v\b\n", ba)
		fmt.Printf("% v\n\n", ra)
	}

	// iters
	{
		s := "привет"

		// iteration per bytes
		for i := 0; i < len(s); i++ {
			//b := s[i]
			// i - index, 1,2,3...
			// b - byte, uint8
		}

		//iteration by runes
		for i, r := range s {
			fmt.Printf("%v %v\n", i, r)
			// i - 1,2,4,6,9... (by rune len)
			// r - rune, int32
		}
	}

	// string functs
	{
		//Contains(s, substr string) bool

		// startsWith ?
		//HasPrefix(s, prefix string) bool

		//Join(a []string, sep string) string
		
		//Split(s, sep string) []string
	}
}

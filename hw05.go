package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// --- samples of functions

//simple declarations with/without args and return values
func Hello() {
	fmt.Println("Hello World!")
}

func greet(user string) {
	fmt.Println("Hello " + user)
}

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func Hello2() string {
	//func Hello3() string { // !!! unresolved reference!
	//	return "hello"
	//}
	//return Hello2()

	return ""
}

// return multi-values
func addMult(a, b int) (int, int) {
	return a + b, a * b
}

// return by '='
func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a + b, a - b
	s = x * x
	d = y * y
	return // <=> return s, d
}

// variadic fuctnion
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++ // this is a reference to the I from the parent function
		return i
	}
}

// recurse
// n! = n×(n-1)! where n >0
func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1)
	} else {
		return 1 // 1! == 1
	}
}

// --- end samples of functions

// regular struct
type Employee struct {
	name, surname string
}

func FullName(e Employee) string {
	return e.name + " " + e.surname
}

func (e Employee) FullName() string {
	return e.name + " " + e.surname
}

// basic type
type Age int
func (age Age) LargerThan(a Age) bool {
	return age > a
}

func (age *Age) Increase() {
	*age++
}
// --- Structs methods

// --- end of Structs methods

// idiomatic error handling
//func (router HttpRouter) parse(reader *bufio.Reader) (Request, error) {
//	requestText, err := readCRLFLine(reader) //string, err Response
//	if err != nil {
//		//No input, or it doesn't end in CRLF
//		return nil, err
//	}
//	requestLine, err := parseRequestLine(requestText) //RequestLine, err Response
//	if err != nil {
//		return nil, err
//	}
//	if request := router.routeRequest(requestLine); request != nil {
//		return request, nil
//	}
//	//Valid request, but no route to handle it
//	return nil, requestLine.NotImplemented()
//}

func hw05() {
	// scopes

	// scope redefinition in inner scope
	{
		var a = 4
		println(a) // 4
		{
			println(a) // 4
			var a = 22
			println(a) // 22
		}
	}

	{
		v := "outer"
		fmt.Println(v)
		{
			v := "inner"
			fmt.Println(v)
			{
				fmt.Println(v)
			}
		}
		{
			fmt.Println(v)
		}
		fmt.Println(v)
	}

	// scope sample for redefine 'err'
	parseInt := func(s string) (int, error) {
		n, err := strconv.Atoi(s)
		if err != nil {
			b, err := strconv.ParseBool(s)
			if err != nil {
				return 0, err
			}
			if b {
				n = 1
			}
		}
		return n, err
	}

	fmt.Println(parseInt("TRUE"))

	// ---

	// functions
	// check functions sample above

	// variadic function call
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// anonymous functions
	func() {
		fmt.Println("Hello!")
	}() // "Hello!"

	var foo1 func() = func() {
		fmt.Println("Hello!")
	}
	foo1() // Hello!

	foo2 := func() {
		fmt.Println("Hello!")
	}
	foo2()

	// real usage of anonymous function
	people := []string{"Alice", "Bob", "Dave"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println(people)

	// closures
	// inner function in closure has access to the parent vars, so we have something like 'local static var'
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	newInts := intSeq()
	fmt.Println(newInts()) // 1

	// closure real usage sample
	//import (
	//	"fmt"
	//"net/http"
	//)
	//func main() {
	//	http.HandleFunc("/hello", hello)
	//	http.ListenAndServe(":3000", nil)
	//}
	//func hello(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "<h1>Hello!</h1>")
	//}

	// regular structs methods
	print(Employee{"alexander", "davydov"}.FullName())
	print(FullName(Employee{"alexander", "davydov"}))

	// basic type methods
	var a Age // 0
	a = a + 1 // 1
	a.Increase() //2

	// error handling
	// wrap errors
	//err := errors.New("error")
	//err = errors.Wrap(err, "open failed")
	//err = errors.Wrap(err, "read config failed")
	//fmt.Println(err) // read config failed: open failed: error
	//fmt.Printf("%+v\n", err) // stacktrace

	// defer
	{
		a := func() {
			i := 0
			defer fmt.Println(i)
			i++
			return
		}
		a()
	}

	{
		b := func() {
			for i := 0; i < 4; i++ {
				defer fmt.Print(i)
			}
		}
		b()
	}

	{
		c := func() (i int) {
			defer func() { i++ }()
			return 1
		}
		c()
	}

	// panic
	{
		var user = os.Getenv("USER")

		if user == "" {
			panic("no value for $USER")
		}
	}

	// recover
	{
		//type Work struct {
		//	name string
		//}
		//
		//do := func(work *Work) {
		//
		//}
		//
		//safelyDo := func(work *Work) {
		//	defer func() {
		//		if err := recover(); err != nil {
		//			log.Println("work failed:", err)
		//		}
		//	}()
		//	do(work)
		//}

		//server := func(workChan <-chan *Work) {
		//	for work := range workChan {
		//		go safelyDo(work)
		//	}
		//}

		//workChan := chan<- Work{"1223"}
		//server()
	}

}

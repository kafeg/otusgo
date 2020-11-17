package main

import "fmt"

// interface - is a set of methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Diuck interface
type Duck interface {
	Talk() string
	Walk()
	Swim()
}

// And Dog - struct which implicitly realizes the Duck interface, by realizing required methods
// So Dog is a dog, but it realizes Duck interface
type Dog struct {
	name string
}
func (d Dog) Talk() string {
	return "AGGGRRRR"
}
func (d Dog) Walk() { }
func (d Dog) Swim() { }

func quack(d Duck) {
	print(d.Talk()) // here we can call Talk for any animal who realizes 'Duck' interface
}

// we can realize internal interfaces, for e.g. Stringer:
//type Stringer interface {
//	String() string
//}

type MyVeryOwnStringer struct { s string}
func (s MyVeryOwnStringer) String() string {
	return "my string representation of MyVeryOwnStringer"
}


// any type can realize many interfaces
type Hound interface {
	Hunt()
}
type Poodle interface {
	Bark()
}

type GoldenRetriever struct{name string}

func (GoldenRetriever) Hunt() { fmt.Println("hunt") }
func (GoldenRetriever) Bark() { fmt.Println("bark") }

func f1(i Hound) { i.Hunt() }
func f2(i Poodle) { i.Bark() }

// many types could realize the same interface
type ScandinavianClip struct{name string}
func (d ScandinavianClip) Bark() { fmt.Printf("%v barked\n", d.name) }
type ToyPoodle struct{name string}
func (d ToyPoodle) Bark() { fmt.Printf("%v barked\n", d.name) }

// one interface can include another. For e.g. from the 'io' package: Reader, Writer, ReadWriter { Reader, Writer }
type Greeter interface {
	hello()
}
type Stranger interface {
	Bye() string
	Greeter
	fmt.Stringer
	//but the method names should not be repeated/conflicted
}

// type switch
type I1 interface { M1() }
type T1 struct{}
func (T1) M1() {}
type I2 interface { I1; M2() }
type T2 struct{}
func (T2) M1() {}
func (T2) M2() {}

func hw07() {
	quack(Dog{}) // we created Dog, but pass it as Duck interface

	fmt.Println(MyVeryOwnStringer{"hello"}) // my string representation of MyVeryOwnStringer{}, fmt.Println knows we have realized 'Stringer.String()' and call that method

	// any type can realize many interfaces
	t := GoldenRetriever{"jack"}
	f1(t) // "hunt" by Hound interface
	f2(t) // "bark" by Poodle interface

	// many types could realize the same interface
	var tp, sc Poodle
	tp = ToyPoodle{"jack"}
	sc = ScandinavianClip{"jones"}
	tp.Bark() // "bark"
	sc.Bark() // "bark"

	// type switch
	var v I1
	switch v.(type) {
	case T1:
		fmt.Println("T1")
	case T2:
		fmt.Println("T2")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("default")
	}
}

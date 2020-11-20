package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup
type DogWait struct { name string; walkDuration time.Duration }
func (d DogWait) Walk(wg *sync.WaitGroup) {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
	wg.Done() //goroutine finished
}

func (d *DogWait) Bark() { fmt.Printf("%s", d.name) }

// --- sync.Cond
func (d *DogWait) Eat(food *DogFood) {
	food.Lock()
	food.cond.Wait()
	food.food--
	food.Unlock()
}

type DogFood struct {
	sync.Mutex
	food int
	cond *sync.Cond
}

func NewDogFood(food int) *DogFood {
	r := DogFood{food: food}
	r.cond = sync.NewCond(&r)
	return &r
}
// --- sync.Cond

// sync.Mutex, sample from stdlib
//type Mutex struct {
//	state int32
//	sema uint32
//}
//// A Locker represents an object that can be locked and unlocked.
//type Locker interface {
//	Lock()
//	Unlock()
//}

var i int // i == 0
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}


func hw09() {

	// sync.WaitGroup
	dogs := []DogWait{{"vasya", time.Second}, {"john", time.Second}}
	var wg sync.WaitGroup
	for _, d := range dogs {
		wg.Add(1) //add new goroutine and wait for it's finish
		go d.Walk(&wg)
	}
	wg.Wait() //wait for all goroutines to finish
	fmt.Println("everybody's home")

	// sync.Mutex
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}
	wg.Wait()
	fmt.Println("value of i after 1000 operations is", i)

	// sync.Pool
	var dogPack = sync.Pool{
		New: func() interface{} { return &DogWait{} },
	}

	dog := dogPack.Get().(*DogWait) //take Dog from the pool
	dog.name = "ivan"
	dog.Bark()
	dogPack.Put(dog) // return Dog to the pool

	// sync.Once
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	// sync.Cond
	food := NewDogFood(4)
	for _, d := range []DogWait{{name: "Vasya"}, {name: "Bob"}} {
		wg.Add(1)
		go func(d DogWait) {
			defer wg.Done()
			d.Eat(food)
		}(d)
	}
	println("Waiting for food to arrive...\n")
	time.Sleep(1 * time.Second)
	food.cond.Broadcast()
	wg.Wait()
	fmt.Printf("Food left: %d\n", food.food)

	// Race detector
	c1 := make(chan bool)
	m1 := make(map[string]string)
	go func() {
		m1["1"] = "a" // First conflicting access.
		c1 <- true
	}()
	m1["2"] = "b" // Second conflicting access.
	<-c1
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
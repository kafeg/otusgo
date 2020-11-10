package hw01 // package name

// includes/imports
import (
	"github.com/beevik/ntp" // for e.g. external module directly from the github https://github.com/beevik/ntp
	"fmt" // internal Golang modules
	"os"
	"time"
)

// function without args but returning two unnamed values
func getCurrentTime() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time, err
}

// function with args, but not returning any values
func printCurrentTime(time time.Time) {
	fmt.Printf("Time: %v\n", time) // fmt formats: https://golang.org/pkg/fmt/
}

func getCurrentTimeWithMetaData() (response *ntp.Response, err error) {
	response, err = ntp.Query("0.beevik-ntp.pool.ntp.org") // here return values by names, not pisitional as in 'getCurrentTime'
	return // return statement is required
}

// sample func for test case, count runes in the given string
func Count(s string, r rune) int {
	var cnt int
	for _, l := range s {
		if l == r {
			cnt += 1
		}
	}
	return cnt
}

// samples to learn
func hw01() {
	time1, err := getCurrentTime() // call a function, vars created automatically with required types by the oper :=, similar to 'auto' in C++
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in NTP library 1: %v\n", err) //print error to stderr channel, '\n' required
		os.Exit(1) // exit from app with return code 1
	}
	printCurrentTime(time1) // Local from the NTP

	response, err := getCurrentTimeWithMetaData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in NTP library 2: %v\n", err) //print error to stderr channel, '\n' required
		os.Exit(1) // exit from app with return code 1
	}
	time2 := time1.Local().Add(response.ClockOffset) //correct offset between server and client
	printCurrentTime(time2)

	fmt.Printf("Offset: %v\n", response.ClockOffset)

	// --- just random basic samples from the lection

	// vars
	{
		fmt.Printf("Vars\n")
		var a int = 5 // var a with type 'int'
		var b = 6 // var b with auto-type from value
		c := 7 // var c with auto-type from value and with operator := instead of 'var'
		fmt.Printf("%d, %d, %d\n", a, b, c)

		c, d := 9, 132 // new value for c and new var d
		b = 15
		fmt.Printf("%d, %d, %d, %d\n", a, b, c, d)
	}

	// constants
	{
		fmt.Printf("Constants\n")
		const (
			a = 1 // a == 1
			b = 2 // b == 2
			c    // c == 2
			d    // d == 2
		)
		fmt.Printf("%d %d %d %d\n", a, b, c, d)

	}

	// constants and iota, iota is special incrementer
	{
		fmt.Printf("Constants iota\n")
		const (
			a = iota // a == 0
			b = iota // b == 1
			c = iota // c == 2
			d = iota * 2 // here iota should be 3 and we multiplies it on 2
			e // e == 8 (implicitely e = iota * 2)
		)
		fmt.Printf("%d %d %d %d %d\n", a, b, c, d, e)
	}

	// 'if' block
	{
		fmt.Printf("Conditions\n")
		if response.ClockOffset < 1 {
			fmt.Printf("Offset less than 1 second!\n")
		} else {
			fmt.Printf("Offset more than 1 second!\n")
		}

		//define global a, b and some stupid usage to prevent 'unused' err
		a, b := 1, 2
		a = b
		b = a

		//here a, b is in local scope for the 'if' block, vars above ignored
		if a, b := 5, 7; a > b {
			println("yes")
		} else {
			println("no")
		}
	}

	// cycles
	{
		fmt.Printf("Cycles\n")
		//simple for
		for i := 0; i < 10; i++ {
			println(i)
		}

		//while
		i := 0
		for i < 10 {
			println(i)
			i++
		}

		//infinity cycle
		i = 0 //use i from above
		for {
			println(i)
			i++
			if i > 10 {
				break
			}
		}
	}

	// iterate over array with operator 'range', we have i - index and v - value
	{
		fmt.Printf("Range array\n")
		a := [...]int{81, 54, 43, 66} // array can't be increased, it's read-only type, but can be copied
		for i, v := range a {
			println(i, v)
		}
	}

	//switches
	{
		fmt.Printf("Switches\n")

		//simple, like with 'break' on each  case
		a := 15
		switch a {
		case 10:
			println("One")
		case 15:
			println("Two")
		case 30:
			println("Three")
		default:
			println("DEFAULT!")
		}

		//like without 'break' on each step, 'fallthrough' used for it (inversed break logic)
		switch a {
		case 10:
			println("One")
			fallthrough
		case 15:
			println("Two")
			fallthrough
		case 30:
			println("Three")
			fallthrough
		default
			:println("DEFAULT!")
		}

		//conditional switch
		a = 10
		switch {
		case a < 10:
			println("One")
		case a < 30 && a >= 10:
			println("Two")
		case a > 3:
			println("Three")
		default:
			println("DEFAULT!")
		}
	}


}

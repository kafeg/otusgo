package main // package name

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

func main() {
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
}

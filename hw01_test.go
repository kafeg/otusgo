package hw01 //the same name with 'hw01.go' package file

import (
	"testing" //internal testing module (TestCountNative)
	"github.com/stretchr/testify/require" // external testing module (TestCountTestify)
)

func TestMainHW01(t *testing.T) {
	hw01()
}

//simple test case
func TestGetTime(t *testing.T) {

	time, err := getCurrentTime()

	if err != nil {
		t.Fatalf("Error in NTP library: %v", err)
	}

	if len(time.String()) == 0 {
		t.Fatalf("Time is empty")
	}
}

func TestCountNative(t *testing.T) {
	s := "qwerasdfe"
	e := 2
	if c := Count(s, 'e'); c != e {
		t.Fatalf("bad count for %s: got %d expected %d", s, c, e)
		//t.Fail()   // mark test as broken but continue execution
		//t.FailNow()  // mark test as broken and stop current test-case
		//t.Logf(formar string, ...interface{})  // debug message
		//t.Errorf(formar string, ...interface{})  // t.Logf + t.Fail
		//t.Fatalf(formar string, ...interface{})  // t.Logf + t.FailNow
		//t.SkipNow()  // skip current test-case
	}
}

func TestCountTestify(t *testing.T) {
	s := "qwerasdfe"
	require.Equal(t, Count(s, 'e'), 2, "counting 'e' in "+s)
	require.Equal(t, Count(s, 'x'), 0, "counting 'x' in "+s)
	require.Equal(t, Count(s, 'f'), 1, "counting 'f' in "+s)
}

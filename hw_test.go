package main //the same name with 'hw01.go' package file

import (
	"fmt"
	"math/rand"
	"testing"                             //internal testing module (TestCountNative)
	"github.com/stretchr/testify/require" // external testing module (TestCountTestify)
	"time"
)


// --- hw01
func TestMainHW01(t *testing.T) {
	fmt.Printf("--- Tests for hw01 ---\n")
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


// --- hw03
func TestMainHW03(t *testing.T) {
	fmt.Printf("--- Tests for hw03 ---\n")
	hw03()
}

// test for itoa func from hw03.go
func TestITOA(t *testing.T) {
	type pair struct {
		i int
		s string
	}

	tests := []pair{
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-3, "-3"},
	}

	for _, tst := range tests {
		if tst.s == itoa(tst.i) {
			fmt.Printf("%d - %s\n", tst.i, "OK")
		} else {
			fmt.Printf("%d - %s\n", tst.i, "FAIL")
		}

		require.Equal(t, tst.s, itoa(tst.i), "itoa '"+tst.s+"'")
	}
}

func TestUnpackString(t *testing.T) {
	type pair struct {
		in string
		out string
	}

	tests := []pair{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
	}

	for _, tst := range tests {
		require.Equal(t, tst.out, unpackString(tst.in), "unpackString '"+tst.out+"'")
	}
}


// --- hw04
func TestMainHW04(t *testing.T) {
	fmt.Printf("--- Tests for hw04 ---\n")
	hw04()
}

func TestWordCount(t *testing.T) {
	testString := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean dictum dignissim pharetra. Nullam nec eros laoreet, tempus sapien at, luctus justo. Fusce orci erat, ac lorem rutrum, lobortis dignissim metus. Maecenas sagittis tellus commodo orci interdum feugiat. Vestibulum aliquam dolor vel gravida viverra. Quisque porta fermentum aliquam. Nam molestie in metus ut viverra.
Ut vulputate elementum. Curabitur facilisis consequat elit, vitae eros elementum sit amet. Proin faucibus nulla posuere, interdum augue quis, consectetur eros eros eros eros eros. Cras egestas diam vel vestibulum luctus. Sed ut augue finibus, efficitur augue at, dapibus turpis. In congue feugiat arcu sit amet rhoncus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Vestibulum venenatis varius sem, id pulvinar mauris ultrices sed. Mauris quis rutrum massa.
Vestibulum in tellus non nulla volutpat rutrum et ac dolor. Etiam tempor volutpat orci vitae euismod. Praesent ut massa dictum, enim id, euismod nunc. Integer nunc lorem, fermentum id laoreet a, tincidunt a leo. Sed pulvinar est ac odio feugiat, id suscipit metus cursus. Maecenas eget condimentum est. Ut lobortis lobortis nisl sit amet consequat.
Donec ullamcorper congue varius. Duis sagittis tellus nisl, sed eleifend nulla lacinia non. Nunc sed libero ac sem sodales maximus. Pellentesque pellentesque pellentesque enim. Vestibulum sem mauris, pharetra et eros nec, dictum dignissim. Nulla facilisi. Nullam iaculis lobortis nibh, nec vestibulum sem sollicitudin at. Sed commodo arcu eu elit tempus, eget pulvinar est posuere. Praesent ac leo rhoncus, rhoncus felis eu, posuere magna. Curabitur aliquet, nibh sit amet vestibulum pellentesque, lectus quam ullamcorper urna, sed facilisis tellus nunc eget libero. Nunc et tempor dolor. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.
Fusce maximus eros in lacus vestibulum, eu accumsan ligula aliquet. Nullam semper vestibulum nibh, in mattis risus laoreet ac. Aliquam erat volutpat. Vestibulum auctor lectus vel lobortis egestas. Nunc ullamcorper tincidunt massa, eu sagittis nisi consectetur at. Vestibulum leo ipsum, pharetra a porttitor ac, venenatis vitae lectus. Integer elementum risus a bibendum placerat. Vestibulum a blandit metus. Nunc et varius mi, vel dictum sem. Mauris non ante ornare, convallis leo id, semper enim. Sed posuere scelerisque lorem quis finibus. Vestibulum nulla ligula, imperdiet ut felis eget, auctor maximus dolor.
Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vivamus molestie lorem et convallis vestibulum. Nunc cursus iaculis neque, sodales cursus nulla blandit ut. Integer sit amet aliquam erat, non tempor lacus. Quisque ligula quam, finibus consectetur quam eget, congue blandit ligula. Etiam mauris velit, cursus eu nunc ac, tincidunt cursus neque. Pellentesque tincidunt leo. Aenean luctus libero nec magna gravida, pharetra congue nunc elementum. Quisque lobortis mattis sapien vel rhoncus. Vestibulum quis sapien vel enim tempor scelerisque quis ut tellus. Curabitur leo metus, aliquet lobortis sagittis sit amet, blandit eu risus. Aenean eu tellus risus. Suspendisse at pharetra odio, in porttitor enim.
Nulla a aliquet leo. Proin sollicitudin sit amet nibh eget pellentesque. Nullam vitae faucibus felis. Duis non tortor efficitur, venenatis sapien feugiat, finibus arcu. Duis a nulla pharetra pharetra pharetra.`

	testResult := map[string]int {
		"amet":6,
		"cursus":5,
		"eros":9,
		"lobortis":7,
		"nulla":7,
		"nunc":10,
		"pellentesque":6,
		"pharetra":8,
		"tellus":6,
		"vestibulum":13,
	}

	wordsCnt := wordsCount(testString)
	fmt.Printf("%v\n", wordsCnt)

	require.Equal(t, testResult, wordsCnt, "wordsCount")
}


// --- hw05
func TestMainHW05(t *testing.T) {
	fmt.Printf("--- Tests for hw05 ---\n")
	hw05()
}


// --- hw06
func TestMainHW06(t *testing.T) {
	fmt.Printf("--- Tests for hw06 ---\n")
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

	//fmt.Printf("--- TestMainDoubleLinkedList ---\n")
}


// --- hw07
func TestMainHW07(t *testing.T) {
	fmt.Printf("--- Tests for hw07 ---\n")
	hw07()
}

// --- hw08
func TestMainHW08(t *testing.T) {
	fmt.Printf("--- Tests for hw08 ---\n")
	hw08()
}

func TestDoParallelTasks(t *testing.T) {
	var tasks []func() error

	for i := 0; i < 10; i++ {
		tasks = append(tasks, func() error {
			secs := (rand.Intn(9 - 1) + 1)
			time.Sleep(time.Duration(secs) * time.Second)

			if secs % 2 == 0 {
				return nil
			} else {
				return &ErrorString{"TaskError"}
			}
		})
	}

	err := doParallelTasks(tasks, 3, 7)
	if err != nil {
		t.Fatalf("Err in parallel jobs: %v", err)
	}
}


// --- hw09
func TestMainHW09(t *testing.T) {
	fmt.Printf("--- Tests for hw09 ---\n")
	hw09()
}

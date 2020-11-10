package hw03

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMainHW03(t *testing.T) {
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
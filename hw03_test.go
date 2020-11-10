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

	test := []pair{
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-3, "-3"},
	}

	for _, td := range test {
		if td.s == itoa(td.i) {
			fmt.Printf("%d - %s\n", td.i, "OK")
		} else {
			fmt.Printf("%d - %s\n", td.i, "FAIL")
		}

		require.Equal(t, td.s, itoa(td.i), "itoa '"+td.s+"'")
	}
}

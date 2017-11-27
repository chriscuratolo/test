package boo_test

import (
	"testing"

	"github.com/test1/foo"
)

// data and error string output
var cases = []struct {
	number int
	log    string
}{
	{99, "got 100"},
	{100, "got 101"},
}

//unit test
func TestBoo(t *testing.T) {
	// t.Fatal("not implemented")
	b := foo.Boo(cases[0].number)
	if *b != 101 {
		t.Error(cases[0].log)
	}
}

package main

import (
	"testing"
)

func TestQ1(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "123456789012",
			out: 1,
		},
	}
	for i, test := range tests {
		if o := Q1(test.in, 3, 2); o != test.out {
			t.Errorf("TestQ1 %v: expected %v got %v\n", i, test.out, o)
		}
	}
}

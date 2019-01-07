package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		desc   string
		input  uint
		result uint
		ok     bool
	}{
		{
			desc:   "zero",
			input:  0,
			result: 1,
			ok:     true,
		},
		{
			desc:   "one",
			input:  1,
			result: 1,
			ok:     true,
		},
		{
			desc:   "two",
			input:  2,
			result: 2,
			ok:     true,
		},
		{
			desc:   "three",
			input:  3,
			result: 6,
			ok:     true,
		},
		{
			desc:   "eleven",
			input:  11,
			result: 39916800,
			ok:     true,
		},
		{
			desc:   "overflow 22",
			input:  22,
			result: 0,
			ok:     false,
		},
		{
			desc:   "overflow 33",
			input:  33,
			result: 0,
			ok:     false,
		},
		{
			desc:   "overflow 44",
			input:  44,
			result: 0,
			ok:     false,
		},
	}

	for _, tt := range tests {
		want, ok := factorial(tt.input)
		if ok != tt.ok {
			t.Fatalf("test \"%s\" expected 'ok' = %t, got %t 'result' = %d", tt.desc, tt.ok, ok, want)
		}
		if want != tt.result {
			t.Fatalf("test: \"%s\" expected 'result' = %d, got %d", tt.desc, tt.result, want)
		}
	}
}

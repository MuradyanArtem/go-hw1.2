package rpn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  []string
	}{
		{
			name: "valid expression",
			in:   "2*3+10",
			out: []string{
				"2",
				"3",
				"*",
				"10",
				"+",
			},
		},
		{
			name: "valid expression 2",
			in:   "10+2*3",
			out: []string{
				"10",
				"2",
				"3",
				"*",
				"+",
			},
		},
		{
			name: "valid expression braces 1",
			in:   "(10+2)*3",
			out: []string{
				"10",
				"2",
				"+",
				"3",
				"*",
			},
		},
		{
			name: "invalid token",
			in:   "(10+2&*3",
			out:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := New(tc.in)
			assert.Equal(t, tc.out, res)
		})
	}
}

func TestCalculate(t *testing.T) {
	testCases := []struct {
		name string
		in   []string
		out  int
	}{
		{
			name: "valid expression",
			in: []string{
				"2",
				"3",
				"*",
				"10",
				"+",
			},
			out: 16,
		},
		{
			name: "valid expression 2",
			in: []string{
				"10",
				"2",
				"3",
				"*",
				"+",
			},
			out: 16,
		},
		{
			name: "valid expression 3",
			in: []string{
				"10",
				"2",
				"+",
				"3",
				"*",
			},
			out: 36,
		},
		{
			name: "valid expression 3",
			in: []string{
				"10",
				"2",
				"+",
				"3",
				"*",
			},
			out: 36,
		},
		{
			name: "valid expression 4",
			in: []string{
				"10",
				"2",
				"+",
				"3",
				"/",
			},
			out: 4,
		},
		{
			name: "valid expression 5",
			in: []string{
				"10",
				"2",
				"+",
				"3",
				"-",
			},
			out: 9,
		},
		{
			name: "invalid expression 1",
			in: []string{
				"10",
				"#",
				"+",
				"3",
				"*",
			},
			out: 0,
		},
		{
			name: "invalid expression 2",
			in: []string{
				"+",
			},
			out: 0,
		},
		{
			name: "invalid expression 2",
			in: []string{
				"-",
			},
			out: 0,
		},
		{
			name: "invalid expression 3",
			in: []string{
				"*",
			},
			out: 0,
		},
		{
			name: "invalid expression 4",
			in: []string{
				"/",
			},
			out: 0,
		},
		{
			name: "invalid expression 5",
			in: []string{
				"4",
				"0",
				"/",
			},
			out: 0,
		},
		{
			name: "invalid expression 6",
			in: []string{
				"0",
			},
			out: 0,
		},
		{
			name: "invalid expression 7",
			in: []string{
				"0",
				"0",
			},
			out: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := Calculate(tc.in)
			assert.Equal(t, tc.out, res)
		})
	}
}

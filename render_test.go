package vrender

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func Test_Render(t *testing.T) {
	type testCase struct {
		name     string
		input    interface{}
		expected string
	}
	cases := []*testCase{
		{
			name:     "Prints zero time as time.Time{}",
			input:    time.Time{},
			expected: "time.Time{}",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Render(c.input)
			assert.Equal(t, out, c.expected)
		})
	}
}

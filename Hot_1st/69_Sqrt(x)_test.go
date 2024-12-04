package hot

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 4, expected: 2},
		{input: 8, expected: 2},
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		input := rand.Intn(1 << 31)
		expected := int(math.Sqrt(float64(input)))

		tests = append(tests, struct {
			input    int
			expected int
		}{input: input, expected: expected})
	}

	for _, test := range tests {
		res := MySqrt(test.input)
		if !equal69(res, test.expected) {
			t.Errorf("Input:%v, expected:%v, got:%v", test.input, test.expected, res)
		}
	}

	for _, test := range tests {
		res := MySqrt2(test.input)
		if !equal69(res, test.expected) {
			t.Errorf("Input:%v, expected:%v, got:%v", test.input, test.expected, res)
		}
	}
}

func equal69(a, b int) bool {
	return a == b
}

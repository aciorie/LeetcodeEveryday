package hot

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "/home/", expected: "/home"},
		{input: "/home//foo/", expected: "/home/foo"},
		{input: "/home/user/Documents/../Pictures", expected: "/home/user/Pictures"},
		{input: "/../", expected: "/"},
		{input: "/.../a/../b/c/../d/./", expected: "/.../b/d"},
		{input: "/.././GVzvE/./xBjU///../..///././//////T/../../.././zu/q/e",
			expected: "/zu/q/e"},
	}

	for _, test := range tests {
		res := SimplifyPath(test.input)
		if !equal71(res, test.expected) {
			t.Errorf("Input:%v,\nexpect:%v,\ngot:%v", test.input, test.expected, res)
		}
	}
}

func equal71(a, b string) bool {
	return a == b
}

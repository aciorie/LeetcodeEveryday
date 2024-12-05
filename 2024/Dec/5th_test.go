package Dec

import "testing"

func TestCanChange(t *testing.T) {
	tests := []struct {
		start    string
		target   string
		expected bool
	}{
		{start: "_L__R__R_", target: "L______RR", expected: true},
		{start: "R_L_", target: "__LR", expected: false},
		{start: "_R", target: "R_", expected: false},
	}

	for _, test := range tests {
		res := CanChange(test.start, test.target)
		if res != test.expected {
			t.Errorf("start:%v,target:%v,expect:%v,got:%v", test.start, test.target, test.expected, res)
		}
	}
}

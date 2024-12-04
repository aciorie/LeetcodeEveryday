package Dec

import "testing"

func TestCanMakeSubsequence(t *testing.T) {
	tests := []struct {
		str1   string
		str2   string
		expect bool
	}{
		{"abc", "ad", true},
		{"zc", "ad", true},
		{"ab", "d", false},
	}

	for _, test := range tests {
		result := CanMakeSubsequence(test.str1, test.str2)
		if result != test.expect {
			t.Errorf("CanMakeSubsequence(%q, %q) = %v; want %v", test.str1, test.str2, result, test.expect)
		}
	}
}

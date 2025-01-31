package hot

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	inputs := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}

	for _, input := range inputs {
		res := FindAnagrams2(input.s, input.p)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("s:%v,p:%v,want:%v,got:%v", input.s, input.p, input.want, res)
		}
	}
}

package jan

import (
	"reflect"
	"testing"
)

func TestVowelStrings(t *testing.T) {
	inputs := []struct {
		words   []string
		queries [][]int
		want    []int
	}{
		{[]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}, []int{2, 3, 0}},
		{[]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}, []int{3, 2, 1}},
	}

	for _, input := range inputs {
		res := VowelStrings(input.words, input.queries)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("words:%v,queries:%v,want:%v,got:%v", input.words, input.queries, input.want, res)
		}
	}
}

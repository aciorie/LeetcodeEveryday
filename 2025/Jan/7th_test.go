package jan

import (
	"reflect"
	"testing"
)

func TestStringMatching(t *testing.T) {
	inputs := []struct {
		words []string
		want  []string
	}{
		{[]string{"mass", "as", "hero", "superhero"}, []string{"as", "hero"}},
		{[]string{"leetcode", "et", "code"}, []string{"et", "code"}},
		{[]string{"blue", "green", "bu"}, []string{}},
	}

	for _, input := range inputs {
		res := StringMatching(input.words)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("words:%v,want:%v,got:%v", input.words, input.want, res)
		}
	}
}

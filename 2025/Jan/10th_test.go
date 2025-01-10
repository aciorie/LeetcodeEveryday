package jan

import (
	"reflect"
	"testing"
)

func TestWordSubsets(t *testing.T) {
	inputs := []struct {
		words1 []string
		words2 []string
		want   []string
	}{
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}, []string{"facebook", "google", "leetcode"}},
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "l"}, []string{"apple", "google", "leetcode"}},
	}

	for _, input := range inputs {
		res := WordSubsets(input.words1, input.words2)
		if !reflect.DeepEqual(input.want, res) {
			t.Errorf("words1:%v,words2:%v,want:%v,got:%v", input.words1, input.words2, input.want, res)
		}
	}
}

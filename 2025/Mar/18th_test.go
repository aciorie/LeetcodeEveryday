package mar

import (
	"reflect"
	"testing"
)

func TestLongestNiceSubarray(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 8, 48, 10}, 3},
		{[]int{3, 1, 5, 11, 13}, 1},
	}

	for _, input := range inputs {
		res := LongestNiceSubarray(input.nums)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}

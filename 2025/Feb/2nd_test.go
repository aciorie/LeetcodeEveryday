package feb

import "testing"

func TestCheck(t *testing.T) {
	inputs := []struct {
		nums []int
		want bool
	}{
		// {[]int{3, 4, 5, 1, 2}, true},
		{[]int{2, 1, 3, 4}, false},
		{[]int{1, 2, 3}, true},
	}

	for _, input := range inputs {
		res := Check(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}

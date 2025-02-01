package feb

import "testing"

func TestIsArraySpecial(t *testing.T) {
	inputs := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 5}, false},
		{[]int{1}, true},
		{[]int{2, 1, 4}, true},
		{[]int{4, 3, 1, 6}, false},
	}

	for _, input := range inputs {
		res := IsArraySpecial2(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}

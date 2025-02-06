package feb

import "testing"

func TestTupleSameProduct(t *testing.T) {
	inputs := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 4, 6}, 8},
		{[]int{1, 2, 4, 5, 10}, 16},
	}

	for _, input := range inputs {
		res := TupleSameProduct(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}

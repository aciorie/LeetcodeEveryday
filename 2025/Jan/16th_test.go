package jan

import "testing"

func TestXorAllNums(t *testing.T) {
	inputs := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{2, 1, 3}, []int{10, 2, 5, 0}, 13},
		{[]int{1, 2}, []int{3, 4}, 0},
	}
	for _, input := range inputs {
		res := XorAllNums(input.nums1, input.nums2)
		if res != input.want {
			t.Errorf("nums1:%v,nums2:%v,want:%v,got:%v", input.nums1, input.nums2, input.want, res)
		}
	}
}

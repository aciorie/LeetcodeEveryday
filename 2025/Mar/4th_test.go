package mar

import "testing"

func TestCheckPowersOfThree(t *testing.T) {
	inputs := []struct {
		nums int
		want bool
	}{
		// {[]int{3, 4, 5, 1, 2}, true},
		{12, true},
		{91, true},
		{21, false},
	}

	for _, input := range inputs {
		res := CheckPowersOfThree(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}

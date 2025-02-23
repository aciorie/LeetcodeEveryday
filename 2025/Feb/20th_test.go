package feb

import "testing"

func TestFindDifferentBinaryString(t *testing.T) {
	inputs := []struct {
		nums []string
		want string
	}{
		{[]string{"01", "10"}, "11"},
		{[]string{"00", "01"}, "11"},
		{[]string{"111", "011", "001"}, "101"},
	}

	for _, input := range inputs {
		res := FindDifferentBinaryString(input.nums)
		if res != input.want {
			t.Errorf("nums:%v,want:%v,got:%v", input.nums, input.want, res)
		}
	}
}

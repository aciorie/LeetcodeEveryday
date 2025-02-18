package feb

import "testing"

func TestSmallestNumber(t *testing.T) {
	inputs := []struct {
		pattern string
		want    string
	}{
		{"IIIDIDDD", "123549876"},
		{"DDD", "4321"},
	}
	for _, input := range inputs {
		res := SmallestNumber(input.pattern)
		if res != input.want {
			t.Errorf("pattern:%v,want:%v,got:%v", input.pattern, input.want, res)
		}
	}
}

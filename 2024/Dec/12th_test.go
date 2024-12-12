package Dec

import "testing"

func TestPickGifts(t *testing.T) {
	inputs := []struct {
		gifts    []int
		k        int
		expected int64
	}{
		{gifts: []int{25, 64, 9, 4, 100}, k: 4, expected: 29},
		{gifts: []int{1, 1, 1, 1}, k: 4, expected: 4},
	}

	for _, input := range inputs {
		res := PickGifts(input.gifts, input.k)
		if res != input.expected {
			t.Errorf("Gifts:%v,k:%v,expected:%v,got:%v", input.gifts, input.k, input.expected, res)
		}
	}
}

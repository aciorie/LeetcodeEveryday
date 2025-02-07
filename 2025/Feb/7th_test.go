package feb

import (
	"reflect"
	"testing"
)

func TestQueryResults(t *testing.T) {
	inputs := []struct {
		limit   int
		queries [][]int
		want    []int
	}{
		{
			limit:   4,
			queries: [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}},
			want:    []int{1, 2, 2, 3},
		},
		{
			limit:   4,
			queries: [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}},
			want:    []int{1, 2, 2, 3, 4},
		},
	}
	for _, input := range inputs {
		res := QueryResults(input.limit, input.queries)
		if !reflect.DeepEqual(res, input.want) {
			t.Errorf("limit:%v,queries:%v,want:%v,got:%v", input.limit, input.queries, input.want, res)
		}
	}
}

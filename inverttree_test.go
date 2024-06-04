package invertbinarytree

import (
	"testing"
)

func TestInverTreeArray(t *testing.T) {

	examples := [][]int{
		{5, 3, 8, 1, 7, 2, 6},
		{6, 8, 9},
		{5, 3, 8, 1, 7, 2, 6, 100, 3, -1},
		{},
	}

	answer := [][]*int{
		OptIntArray([]int{5, 8, 3, 6, 2, 7, 1}),
		OptIntArray([]int{6, 9, 8}),
		append(OptIntArray([]int{5, 8, 3, 6, 2, 7, 1}), []*int{nil, nil, nil, nil, nil, OptInt(-1), OptInt(3), OptInt(100)}...),
		{},
	}

	for idx, example := range examples {
		myAnswer := PrintOptIntArray(InverTreeArray(example))
		answer := PrintOptIntArray(answer[idx])
		if myAnswer != answer {
			t.Errorf("answer: %v my answer: %v", answer, myAnswer)
		}
	}

}

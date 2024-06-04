package invertbinarytree

import (
	"fmt"
	"strings"
)

// Q1：Given the root of a binary tree, invert the tree, and return its root.(Don’t use recursion.)

// example1:
// input: [5, 3, 8, 1, 7, 2, 6]
// ouput: [5, 8, 3, 6, 2, 7, 1]

// example2:
// input: [6, 8, 9]
// output: [6, 9, 8]

// example3:
// input: [5, 3, 8, 1, 7, 2, 6, 100, 3, -1]
// output: [5, 8, 3, 6, 2, 7, 1, null, null, null, null, null, -1, 3, 100]

// example3:
// input: []
// output: []

// Constraints:
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

func OptInt(v int) *int {
	return &v
}

func OptIntArray(source []int) []*int {
	out := make([]*int, len(source))
	for i, v := range source {
		out[i] = OptInt(v)
	}
	return out
}

func PrintOptIntArray(source []*int) string {
	ary := make([]string, len(source))
	for i, v := range source {
		if v == nil {
			ary[i] = fmt.Sprintf("%v", v)
		} else {
			ary[i] = fmt.Sprintf("%v", *v)
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(ary, ", "))
}

func InverTreeArray(source []int) []*int {
	if source == nil {
		return nil
	}
	l := len(source)
	remain := l
	layer := 1
	layerNeed := 1
	used := 0
	out := make([]*int, l)
	for remain > 0 {
		start := used
		to := used + layerNeed
		for i := 0; i < layerNeed; i++ {
			if sourceIdx := to - i - 1; sourceIdx < l {
				out[start+i] = OptInt(source[sourceIdx])
				used = used + 1
			} else {
				if start+i >= len(out) {
					out = append(out, make([]*int, to-len(out))...)
				}
				out[start+i] = nil
			}
		}
		remain = l - used
		fmt.Printf("layer: %v, layer need: %v, used: %v, remain: %v, out: %v\n", layer, layerNeed, used, remain, PrintOptIntArray(out))

		if remain > 0 {
			layer += 1
			layerNeed = layerNeed * 2
		}
	}
	return out
}

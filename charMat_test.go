package main

import (
	"fmt"
	"testing"
)

var continueListTest = []struct {
	ins  []int
	res  bool
	cont map[int][]int
}{
	{
		ins: []int{1, 2, 3, 4, 5},
		res: true,
		cont: map[int][]int{
			1: {0, 1, 2, 3, 4},
			2: {0, 1, 2, 3},
			3: {0, 1, 2},
			4: {0, 1},
		},
	},
	{
		ins: []int{2, 3, 4, 5},
		res: true,
		cont: map[int][]int{
			1: {0, 1, 2, 3},
			2: {0, 1, 2},
			3: {0, 1},
			4: {0},
		},
	},
	{
		ins: []int{2, 2, 3, 4},
		res: false,
		cont: map[int][]int{
			1: {0, 1, 2, 3},
			2: {1, 2},
			3: {1},
			4: {},
		},
	},
}

func isSliceSame(sl1, sl2 []int) bool {
	z := len(sl1) == len(sl2)
	for i := range sl1 {
		z = z && (sl1[i] == sl2[i])
	}
	return z
}

func TestIsContinue(t *testing.T) {
	for i, test := range continueListTest {
		if test.res != IsContinue(test.ins) {
			t.Errorf("Testing Instance %d resulted in error", i)
		}
	}
}

func TestGetAllContinuousSubList(t *testing.T) {
	for i, test := range continueListTest {
		for n := range test.cont {
			if !isSliceSame(test.cont[n], GetAllContinuousSubList(test.ins, n)) {
				fmt.Println(GetAllContinuousSubList(test.ins, n))
				t.Errorf("Testing Instance %d sub %d resulted in error", i, n)
			}
		}
	}
}

var swapTest = []struct {
	send []int
	recv []int
	il   int
	ir   int
	ns   []int
	nr   []int
}{
	{
		send: []int{1, 2, 3, 4, 5},
		recv: []int{2, 3},
		il:   0,
		ir:   1,
		ns:   []int{3, 4, 5},
		nr:   []int{1, 2, 2, 3},
	},
}

func TestSwapSubList(t *testing.T) {
	for i, test := range swapTest {
		ns, nr := SwapSubList(test.send, test.recv, test.il, test.ir)
		fmt.Println(ns, nr)
		if !(isSliceSame(ns, test.ns) && isSliceSame(nr, test.nr)) {
			t.Errorf("Testing Instance %d resulted in error", i)
		}
	}
}

func TestRandSample(t *testing.T) {
	z := RandSample([]int{1, 2, 3, 4, 5}, 4)
	if len(z) != 4 {
		t.Errorf("Error")
	}
}

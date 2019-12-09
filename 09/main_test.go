package main

import (
	"reflect"
	"testing"
)

func TestQ1(t *testing.T) {
	tests := []struct {
		arr []int64
		out []int64
	}{
		{

			[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		}, {
			[]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			[]int64{1219070632396864},
		}, {
			[]int64{104, 1125899906842624, 99},
			[]int64{1125899906842624},
		},
	}
	for i, tc := range tests {
		out := Q1(tc.arr, nil)
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("TestQ1 case %v expected: \n%v got \n%v\n", i, tc.out, out)
		}
	}
}

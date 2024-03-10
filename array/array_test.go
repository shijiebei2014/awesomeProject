package array

import (
	"reflect"
	"testing"
)

func Test_mergeSorted(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ok", args{arr1: []int{1, 3, 5}, arr2: []int{2, 4, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"ok", args{arr1: []int{10, 30, 50}, arr2: []int{2, 4, 6}}, []int{2, 4, 6, 10, 30, 50}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSorted(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

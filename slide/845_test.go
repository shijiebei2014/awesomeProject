package slide

import "testing"

func Test_longestMountain(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{A: []int{2, 1, 4, 7, 3, 2, 5}}, 5},
		{"2", args{A: []int{2, 2, 2}}, 0},
		{"3", args{A: []int{2, 3}}, 0},
		{"4", args{A: []int{3, 2}}, 0},
		{"5", args{A: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}, 0},
		{"6", args{A: []int{0, 1, 0}}, 3},
		{"7", args{A: []int{0, 1, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.args.A); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}

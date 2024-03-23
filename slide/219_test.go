package slide

import "testing"

func Test_containNearByDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 3, 1}, 3}, true},
		{"2", args{[]int{1, 0, 1, 1}, 1}, true},
		{"3", args{[]int{1, 2, 3, 1, 2, 3}, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containNearByDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containNearByDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

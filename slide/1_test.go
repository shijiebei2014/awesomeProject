package slide

import "testing"

func Test_noRepeat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"baca"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noRepeat(tt.args.s); got != tt.want {
				t.Errorf("noRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

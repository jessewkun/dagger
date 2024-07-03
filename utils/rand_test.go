package utils

import "testing"

func Test_RandomNum(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{1, 10}, 5},
		{"test1", args{0, 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomNum(tt.args.min, tt.args.max)
			t.Logf("got: %d", got)
		})
	}
}

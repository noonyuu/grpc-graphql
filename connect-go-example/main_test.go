package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "add(0, 1)", args: args{x: 0, y: 1}, want: 1},
		{name: "add(1, 1)", args: args{x: 1, y: 1}, want: 2},
		{name: "add(1, 2)", args: args{x: 1, y: 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

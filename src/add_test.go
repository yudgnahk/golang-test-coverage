package golang_test_coverage

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 + 1 = 2",
			args: args{a: 1, b: 1},
			want: 2,
		},
		{
			name: "1 + 3 = 4",
			args: args{a: 1, b: 3},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddThree(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 + 1 + 1 = 3",
			args: args{a: 1, b: 1, c: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddThree(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("AddThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

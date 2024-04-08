package golang_test_coverage

import "testing"

func TestMinus(t *testing.T) {
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
			name: "1 - 1 = 0",
			args: args{a: 1, b: 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Minus(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Minus() = %v, want %v", got, tt.want)
			}
		})
	}
}

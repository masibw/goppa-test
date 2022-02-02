package math

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
			name: "Can add up two numbers..",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "Can add up two numbers(includes negative value).",
			args: args{
				a: -1,
				b: 5,
			},
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
			name: "Can subtract b from a.",
			args: args{
				a: 2,
				b: 1,
			},
			want: 1,
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

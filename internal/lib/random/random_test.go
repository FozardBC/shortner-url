package random

import (
	"testing"
)

func TestNewRandomString(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				size: 3,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				size: 6,
			},
			want: 6,
		},

		{
			name: "3",
			args: args{
				size: 12,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewRandomString(tt.args.size); len(got) != tt.want {
				t.Errorf("NewRandomString() = %v, want %v", got, tt.want)
			}

		})
	}
}

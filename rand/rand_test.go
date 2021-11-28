package rand

import (
	"testing"
)

func TestGenerateN(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				len: 100,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateN(tt.args.len)
			if len(got) != tt.args.len {
				t.Errorf("GenerateN() = %v, want %v", len(got), tt.args.len)
			}
		})
	}
}

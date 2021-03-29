package reconstruct_original_digits_from_english

import "testing"

func Test_originalDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test01",
			args: args{s: "owoztneoer"},
			want: "012",
		},
		{
			name: "test02",
			args: args{s: "fviefuro"},
			want: "45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := originalDigits(tt.args.s); got != tt.want {
				t.Errorf("originalDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

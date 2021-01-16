package make_the_string_great

import "testing"

func Test_makeGood(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "only one matches",
			args: args{s: "leEeetcode"},
			want: "leetcode",
		},
		{
			name: "all matches",
			args: args{s: "abBAcC"},
			want: "",
		},
		{
			name: "single char",
			args: args{s: "s"},
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}

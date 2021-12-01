package internal

import "testing"

func Test_envOrDefault(t *testing.T) {
	type args struct {
		varName   string
		orDefault string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing is it will return default value for empty env var",
			args: args{
				varName:   "XXX",
				orDefault: "iwantthis",
			},
			want: "iwantthis",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := envOrDefault(tt.args.varName, tt.args.orDefault); got != tt.want {
				t.Errorf("envOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

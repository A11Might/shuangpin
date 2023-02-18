package util

import "testing"

func Test_handlingText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{text: " 你 好 ，。你， 不。好, "}, "你好,.你,不.好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandlingText(tt.args.text); got != tt.want {
				t.Errorf("HandlingText() = %v, want %v", got, tt.want)
			}
		})
	}
}

package shuangpin

import (
	"reflect"
	"testing"
)

func TestText(t *testing.T) {
	type args struct {
		title  string
		length int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []string
		want2 string
	}{
		{"test1", args{title: "你好", length: 100}, "", []string{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := TextWithSymbol(tt.args.title, tt.args.length)
			if got != tt.want {
				t.Errorf("Text() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Text() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Text() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_chineseToPinyinWithSymbol(t *testing.T) {
	type args struct {
		chinese string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"test1", args{chinese: "你好,"}, [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chineseToPinyinWithSymbol(tt.args.chinese); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chineseToPinyinWithSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func TestAdd(t *testing.T) {
	//if Add(1, 2) != 3 {
	//	t.Error("1+2 != 3")
	//}
	// 子测试
	t.Run("test1", func(t *testing.T) {
		if Add(1, 2) != 3 {
			t.Error("1+2 != 3")
		}
	})

	t.Run("test2", func(t *testing.T) {
		if Add(2, 2) != 4 {
			t.Error("2+2 != 4")
		}
	})
	//如果测试用例很多，还可以用一个类似表格去表示
	t.Run("test3", func(t *testing.T) {
		type args struct {
			a int
			b int
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			{"test1", args{1, 2}, 3},
			{"test2", args{2, 2}, 4},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Add(tt.args.a, tt.args.b); got != tt.want {
					t.Errorf("Add() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

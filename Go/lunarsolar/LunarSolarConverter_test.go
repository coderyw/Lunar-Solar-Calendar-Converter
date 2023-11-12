// Package lunarsolar
// @Author: yinwei
// @File: LunarSolarConverter_test.go
// @Version: 1.0.0
// @Date: 2023/11/12 22:13

package lunarsolar

import (
	"testing"
)

func TestLunarToSolar(t *testing.T) {
	type args struct {
		lunar Lunar
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "一月初一", args: args{
			lunar: Lunar{
				IsLeap:     false,
				LunarDay:   1,
				LunarMonth: 1,
				LunarYear:  2023,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LunarToSolar(tt.args.lunar)

			t.Log(got)
		})
	}
}

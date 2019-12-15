package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_convertString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{
			"standard",
			args{""},
			[][]bool{{}},
		},
		{
			"simple",
			args{`.#..#
.....
#####
....#
...##`},
			[][]bool{{false, true, false, false, true}, {false, false, false, false, false}, {true, true, true, true, true}, {false, false, false, false, true}, {false, false, false, true, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLineOfSights(t *testing.T) {
	type args struct {
		aMap [][]bool
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"null",
			args{[][]bool{{}}, 0, 0},
			0,
		},
		{
			"example",
			args{[][]bool{{false, true, false, false, true}, {false, false, false, false, false}, {true, true, true, true, true}, {false, false, false, false, true}, {false, false, false, true, true}}, 3, 4},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			if got := calculateLineOfSights(tt.args.aMap, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("calculateLineOfSights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHighestLosCount(t *testing.T) {
	type args struct {
		losCountMap *map[string]int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{
			"one",
			args{&map[string]int{"0,0": 1}},
			"0,0",
			1,
		},
		{
			"8",
			args{&map[string]int{"0,2": 6, "1,0": 7, "1,2": 6, "2,2": 6, "3,2": 6, "3,4": 8, "4,0": 7, "4,2": 5, "4,3": 7, "4,4": 7}},
			"3,4",
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getHighestLosCount(tt.args.losCountMap)
			if got != tt.want {
				t.Errorf("getHighestLosCount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getHighestLosCount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

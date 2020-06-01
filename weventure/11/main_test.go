package main

import "testing"

func TestArea(t *testing.T){
	areaTests := []struct{
		name string
		shape Shape
		hasArea float64
	}{
		{"Rectangle",Rectangle{ 12,6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{ 10,2},10},
	}

	for _, tt := range areaTests{
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g",tt.shape, got, tt.hasArea)
		}
	}
}
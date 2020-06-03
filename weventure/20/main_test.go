package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T){

	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{{
		"struct with one string field",
		struct {
			Name string
		}{"Toan"},
		[]string{"Toan"},
	},
	{
		"struct with two string fields",
		struct{
			Name string
			City string
		}{"Toan", "HCM"},
		[]string{"Toan", "HCM"},
	},
	{
		"struct with non string field",
		struct{
			Name string
			Age int
		}{"Toan", 33},
		[]string{"Toan"},
	},
	{
		"Nested fields",
		struct {
			Name string
			Profile struct{
				Age int
				City string
			}
		}{"Toan", struct {
			Age int
			City string
		}{33, "HCM"}},
		[]string{"Toan", "HCM"},
	},
	{
		"Pointers to things",
		&Person{
			"Chris",
			Profile{33, "London"},
		},
		[]string{"Chris", "London"},
	},
	{
		"Slices",
		[]Profile{
			{33,"London"},
			{34, "Reykjavik"},
		},
		[]string{"London", "Reykjavik"},
	},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls){
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

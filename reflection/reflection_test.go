package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age int
	City string
}

type Person struct {
	Name string
	Profile Profile
}


func TestWalk(t *testing.T) {

	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"Struct with one string field",
			struct {
				Name string
			}{ "Akash"},
			[]string{"Akash"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{ "Akash", "Singapore"},
			[]string{"Akash", "Singapore"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age int
			}{ "Akash", 23},
			[]string{"Akash"},
		},
		{
			"Nested fields",
			Person{
				"Akash",
				Profile{23, "Singapore"},
			},
			[]string{"Akash", "Singapore"},
		},
		{
			"Pointers to things",
			&Person{
				"Akash",
				Profile{23, "Singapore"},
			},
			[]string{"Akash", "Singapore"},
		},
		{
			"Slices",
			[]Profile {
				{23, "Singapore"},
				{22, "Jamshedpur"},
			},
			[]string{"Singapore", "Jamshedpur"},
		},
		{
			"Arrays",
			[2]Profile {
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
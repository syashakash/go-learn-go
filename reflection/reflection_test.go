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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{23, "Singapore"}
			aChannel <- Profile{24, "Berlin"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Singapore", "Berlin"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string)  {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
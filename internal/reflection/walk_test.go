package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"One field struct",
			struct{ Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"Two field struct",
			struct{ Name, City string }{"Chris", "Frankfurt"},
			[]string{"Chris", "Frankfurt"},
		},
		{
			"Two field struct, one with non string type",
			struct {
				Name string
				Age  int8
			}{"Chris", 20},
			[]string{"Chris"},
		},
		{
			"Nested structs",
			Person{
				"Chris",
				Profile{20, "Frankfurt"},
			},
			[]string{"Chris", "Frankfurt"},
		},
		{
			"Nested pointers",
			&Person{
				"Chris",
				Profile{20, "Frankfurt"},
			},
			[]string{"Chris", "Frankfurt"},
		},
		{
			"Array",
			[2]Profile{{21, "Berlin"}, {25, "Hamburg"}},
			[]string{"Berlin", "Hamburg"},
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

	for _, tt := range cases {
		var got []string

		walk(tt.Input, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, tt.ExpectedCalls) {
			t.Errorf("expected %v, got %v", tt.ExpectedCalls, got)
		}
	}
}

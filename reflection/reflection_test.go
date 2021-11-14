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
			"Struct with one string field",
			struct {
				Name string
			}{"Chirs"},
			[]string{"Chirs"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chirs", "London"},
			[]string{"Chirs", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chirs", 33},
			[]string{"Chirs"},
		},
		{
			"Nested fields",
			Person{
				"Chirs",
				Profile{33, "London"},
			},
			[]string{"Chirs", "London"},
		},
		{
			"Pinters to things",
			&Person{
				"Chirs",
				Profile{33, "London"},
			},
			[]string{"Chirs", "London"},
		},
		{
			"Slice",
			[]Profile{
				Profile{33, "London"},
				Profile{34, "Reykjavik"},
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

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}

		})
	}
}

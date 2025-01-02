package myref

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name   string
		Input  interface{}
		Output []string
	}{
		{
			Name: "struct with one field",
			Input: struct {
				Name string
			}{"Peppe"},
			Output: []string{"Peppe"},
		},
		{
			Name: "struct with two field",
			Input: struct {
				Name string
				City string
			}{"Peppe", "London"},
			Output: []string{"Peppe", "London"},
		},
		{
			Name: "struct with one string and one int",
			Input: struct {
				Name string
				age  int
			}{"Peppe", 18},
			Output: []string{"Peppe"},
		},
		{
			Name:   "nested fields",
			Input:  Person{"Chris", Profile{33, "London"}},
			Output: []string{"Chris", "London"},
		},
	}

	for _, el := range cases {
		t.Run(el.Name, func(t *testing.T) {
			var got []string

			walk(el.Input, func(name string) {
				got = append(got, name)
			})

			if !reflect.DeepEqual(got, el.Output) {
				t.Errorf("got %v, want %v", got, el.Output)
			}
		})
	}

	// walk(x interface{}, fn func(string))
}

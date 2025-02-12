package dataprocessing_test

import (
	"reflect"
	"testing"

	data_processing "github.com/gppmad/gocloudplayground/data_processing"
)

func TestDataProcessingRobust(t *testing.T) {

	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{name: "positive", input: []int{1, 5, 3}, expectedOutput: []int{2, 10, 6}},
		{name: "negative", input: []int{-1, -5, -3}, expectedOutput: []int{-2, -10, -6}},
	}

	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			output := data_processing.DataProcessing(el.input)
			if !reflect.DeepEqual(el.expectedOutput, output) {
				t.Errorf("got %v, want %v", output, el.expectedOutput)
			}
		})
	}

}

package common

import (
	"reflect"
	"testing"
)

func TestToInt(t *testing.T) {

	input := []string{"1", "2", "3"}
	want := []int{1, 2, 3}

	got, err := ToInt(input)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(got, want) {
		t.Fail()
	}

}

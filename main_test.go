package main

import (
	"reflect"
	"testing"
)

func TestUnique[T comparable](t *testing.T) {

	testCases := []struct {
		  arg, want []comparable
	}{
		{
			arg:  []int{1, 1, 2, 3, 3, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			arg:  []string{"1", "1", "2", "3", "3", "3", "4"},
			want: []string{"1", "2", "3", "4"},
		},
		{
			arg: []float64{1.0, 1, 2, 3, 3, 3, 4},
			want: []float64{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		testUniqueHelper(t, tc.arg, tc.want)
	}

}

func testUniqueHelper[T comparable](t *testing.T, arg, want []T) {
	t.Helper()

	// got := unique(arg)
	got := uniqueGeneric(arg)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unique(%v) = %v; want %v", arg, got, want)
	}


}

func TestUniqueInts(t *testing.T) {

	arg := []int{1, 1, 2, 3, 3, 3, 4}
	want := []int{1, 2, 3, 4}
	testUniqueHelper(t, arg, want)	
}

func TestUniqueStrings(t *testing.T) {

	arg := []string{"1", "1", "2", "3", "3", "3", "4"}
	want := []string{"1", "2", "3", "4"}
	testUniqueHelper(t, arg, want)	

}

func TestUniqueFloats(t *testing.T) {

	arg := []float64{1.0, 1, 2, 3, 3, 3, 4}
	want := []float64{1, 2, 3, 4}
	testUniqueHelper(t, arg, want)	

}


package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("sum of all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v", want, got)
		}
	})

	t.Run("sum of all tails for empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v", want, got)
		}
	})
}

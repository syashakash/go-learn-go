package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectAnswer := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	}
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertCorrectAnswer(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertCorrectAnswer(t, got, want, numbers)
	})
}

func TestSumALl(t *testing.T) {
	assertCorrectAnswer := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	}

	t.Run("collection of arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0,9})
		expected := []int{3,9}
		assertCorrectAnswer(t, got, expected)
	})

	t.Run("collection of arrays and sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0,9})
		expected := []int{2,9}
		assertCorrectAnswer(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0,9}
		assertCorrectAnswer(t, got, expected)
	})
}
package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 5} //array literal
		//Create a slice out of the array when calling the function
		got := Sum(numbers[:]) //numbers[:] : slice literal
		expected := 15

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {

		numbers := [...]int{1, 2, 3}
		got := Sum(numbers[:])
		expected := 6

		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	})
}

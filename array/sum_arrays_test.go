package sum

import "testing"
import "reflect"

func TestSumAll(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {

		numbers1 := [5]int{1, 2, 3, 4, 5} //array literal
		numbers2 := [5]int{5, 4, 3, 2, 1} //array literal
		//Create a slice out of the array when calling the function
		got := SumAll(numbers1[:], numbers2[:]) //numbers[:] : slice literal
		expected := [5]int{6, 6, 6, 6, 6}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %d want %d given, %v, %v", got, expected, numbers1, numbers2)
		}
	})
}

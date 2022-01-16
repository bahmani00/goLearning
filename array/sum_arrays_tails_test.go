package sum

import "testing"
import "reflect"

func TestSumAllTails(t *testing.T) {
	t.Run("Sum of all Tails", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	})

}

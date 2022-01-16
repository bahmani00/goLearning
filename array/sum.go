package sum

import "fmt"

func Sum(nums []int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}

func PlayWithArrays() {
	//https://stackoverflow.com/a/21722697/336511
	a1 := [3]int{1, 2, 3} // Array literal(value)
	b1 := a1              // Copy the contents of a into b
	a1[0] = 0
	fmt.Println(a1) // Prints "[0 2 3]"
	fmt.Println(b1) // Prints "[1 2 3]"

	a2 := []int{1, 2, 3} // Slice literal(reference)
	b2 := a2             // a and b now point to the same memory
	a2[0] = 0
	fmt.Println(a2) // Prints "[0 2 3]"
	fmt.Println(b2) // Prints "[0 2 3]"
}

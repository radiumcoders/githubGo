package variadic


// functions that can take multiple inputs
func Multiply(numbers ...*int) int {
	product := 1
	for _ , num := range numbers {
		product *= *num
	}
	return product
}

func Add(numbers ...*int) int {
	sum := 0
	for _ , num := range numbers {
		sum += *num
	}
	return sum
}


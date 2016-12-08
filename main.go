// package main is main package
package main

func main() {
	SumInt(1, 2, 3, 4)
}

// SumInt sum integer that puts to function
func SumInt(nums ...int) (total int) {
	for _, i := range nums {
		total += i
	}
	return total
}

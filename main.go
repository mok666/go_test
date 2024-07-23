package gocipractice

func evenOrOdd(n int) (res string) {
	var result string
	if n%2 == 0 {
		result = "even"
	} else {
		result = "odd"
	}

	return result
}

func main() {
	// var n int
	// fmt.Println("Enter a number: ")
	// fmt.Scanln(&n)
	// fmt.Println(evenOrOdd(n))
	evenOrOdd(2)
	evenOrOdd(3)
	evenOrOdd(4)
	evenOrOdd(5)
	evenOrOdd(6)
}

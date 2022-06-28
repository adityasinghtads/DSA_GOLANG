package main

import "fmt"

func main() {
	// Using Hello and print functions
	fmt.Println("Hello World!")

	// Calling other functions in the main function( Basically all the different basics of GO)
	fmt.Println("calling in ifelse function")
	ifelse()
	fmt.Println("Pointer function incoming")
	pointer()
}

func ifelse() {
	a := 0
	var b int
	b = 0
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if a > b {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("B is the greater one")
	}
}

func declareVariable() {
	fmt.Println("Declaring variables in GO \n VAR A INT \n A = 0 \n ********* \n A := 0")
}

func forLoop() {
	fmt.Println("Entering into a For loop of displayng numbers from one to five")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func pointer() {
	fmt.Println("pointing towards pointer")
	a := 10
	b := *&a
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(b)
}

func switchCase() {
	fmt.Println("Entering Switch Case function")
	a := 0
	fmt.Scanln(a)
	switch a {
	case 1:
		fmt.Println("It is in First case")
	case 2:
		fmt.Println("It is in First case")

	case 3:
		fmt.Println("It is in First case")

	case 4:
		fmt.Println("It is in First case")
		fallthrough
	default:
		a = 100
		fmt.Println(a)
	}

}

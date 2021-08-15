// WHY GO?
	Because Go language is an effort to combine the ease of programming of an interpreted, dynamically typed language 
	with the efficiency and safety of a statically typed, compiled language. 
	It also aims to be modern, with support for networked and multicore computing. 


// FIRST GO PROGRAM:

	package main

	//  It contains import “fmt”, 
	// it is a preprocessor command which tells the compiler to include the files lying in the package.
	import "fmt"

	// main function, it is beginning of execution of program.
	func main() {

	// fmt.Println() is a standard library function to print something as a output on screen.
	// In this, fmt package has transmitted Println method which is used to display the output.
		fmt.Println("Hello")
	}


	MISTAKE:

		1) this does not work
		/*
		func main() 
		{
			fmt.Println("Hello")
		}
		*/

		2) all preprocessors should be enclosed in quotes

		"fmt" not fmt
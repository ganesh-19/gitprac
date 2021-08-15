a) writing a function

1) DEFER:
	In Go language, defer statements delay the execution of the function or method or an anonymous method 
	until the nearby functions returns. Or in other words, defer function or method call arguments evaluate 
	instantly, but they execute until the nearby functions returns. 
	You can create a deferred method, or function, or anonymous function by using the defer keyword.

2) SYNTAX:

	Syntax:

	// Function
	defer func func_name(parameter_list Type)return_type{
	// Code
	}

	// Method
	defer func (receiver Type) method_name(parameter_list){
	// Code
	}

	defer func (parameter_list)(return_type){
	// code
	}()


	Important Points:

	    In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out) order as shown in Example 2.
	    In the defer statements, the arguments are evaluated when the defer statement executed, not when they called.
	    Defer statements are generally used to ensure that the files are closed when your work is finished with them, or to close the channel, or to catch the panics in the program.


3) CODE1:

	func maths(a,b int) (int, int){
		mul := a*b
		add := a+b
		return mul, add
	}

	func hero(name string){
		
		fmt.Println("Protagonist is:", name)
	}

	func main(){
		a:= 2
		b:= 4
		
		mul, _ := maths(a,b)
		fmt.Println("mul is :", mul)
		
		//error -- cannot declare variable and anonymous func as defer?
		//defer_mul := defer maths(a,b)
		//fmt.Println("mul using defer is: ", defer_mul)
		
		hero("Nick")
		
		defer hero("Jack")
		
		hero("Camm")	
	}


4) CODE2:

	import (
		"fmt"
	)

	func maths(a,b int) (int, int){
		mul := a*b
		add := a+b
		return mul, add
	}

	func main(){
		a:= 2
		b:= 4
		
		fmt.Println("Start")
		
		//defer is used for builtin function
		defer fmt.Println("End")
		
		mul, add := maths(a,b)
		fmt.Println("2+4 is:", add)
		fmt.Println("2*4 is:", mul)
	}



a) how to define type and struct
b) 

1) METHODS:

	 Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it. 
	 With the help of the receiver argument, the method can access the properties of the receiver. 
	 Here, the receiver can be of struct type or non-struct type. 
	 When you create a method in your code the receiver and receiver type must present in the same package. 
	 And you are not allowed to create a method in which the receiver type is already defined in another package 
	 including inbuilt type like int, string, etc. If you try to do so, then the compiler will give an error.

	Difference Between Method and Function
		Method 	
			It contain receiver. 	
			It can accept both pointer and value. 	
			Methods of the same name but different types can be defined in the program. 	

		Function
			It does not contain receiver.
			It cannot accept both pointer and value.
			Functions of the same name but different type are not allowed to define in the program.
			
2) SYNTAX:
	
	func(reciver_name Type) method_name(parameter_list)(return_type){
	// Code
	}

3) METHOD WITH STRUCT TYPE RECEIVER:

	In Go language, you are allowed to define a method whose receiver is of a struct type. 
	This receiver is accessible inside the method as shown in the below example:


	CODE:

		package main

		import (
			"fmt"
		)

		// struct -- named rangers
		type rangers struct{
			name string
			color string
		}

		// method named profile
		// receives from the variable mystic
		func (ran rangers) profile(){
			fmt.Println(ran.name, ran.color)
		}

		func main(){
			
			// creating a dictionary from the struct rangers
			mystic := rangers{
				name: "Xander",
				color: "green",
			}
			
			// calling the function
			mystic.profile()

		}


4) METHOD WITH NON STRUCT TYPE RECEIVER:

	you are allowed to create a method with non-struct type receiver as long as the type and the method 
	definitions are present in the same package. 
	If they present in different packages like int, string, etc, then the compiler will give an error 
	because they are defined in different packages.

	CODE:

		package main

		import (
			"fmt"
		)

		// type definition
		//here int is referred by num
		type num int


		// method named addn
		// here we need two receivers 
		// (variable type_name)

		/* error
		func (num1 num) addn (num2 num) int {
			return num1 + num2 
		}
		*/

		// need to mention the type at the end as return statement 
		// however instead of int num is used
		func (num1 num) addn (num2 num) num {
			return num1 + num2 
		}


		func main(){
			
			// creating a var for receiver
			value1 := num(20)
			value2 := num(30)
			
			// calling the function
			summ := value1.addn(value2)
			
			fmt.Println("sum is:", summ)

		}

5) METHODS WITH POINTER RECEIVER:

	In Go language, you are allowed to create a method with a pointer receiver. 
	With the help of a pointer receiver if a change made in the method will reflect in the caller 
	which is not possible with the value receiver.


	Syntax:

		func (p *Type) method_name(...Type) Type {
		// Code
		}

	CODE:

		// here author represents struct -- data structure
		type author struct {
			name string
			branch string
		}

		// here *a.branch = res.branch 
		// so *a = res
		// if *a = res then a = &res
		// so *author refers to &res ???


		func (a *author) show_1 (abranch string) {
			(*a).branch = abranch
		}



		func main() {
			
			// initialising values for data structure author
			
			res := author{
				name : "ABC",
				branch : "MECH",
			}
			fmt.Println("Branch name before:", res.branch)

			
			//using &res
			p:= &res
			p.show_1("CIVIL")
			fmt.Println("Branch in res after res.show_1() and using &res is:", res.branch)
			
			
		}

6) METHOD CAN ACCEPT BOTH POINTER AND VALUE:

	
	As we know that in Go, when a function has a value argument, then it will only accept the values of the parameter, 
	and if you try to pass a pointer to a value function, then it will not accept or vice versa. 
	But a Go method can accept both value and pointer whether it is defined with pointer or value receiver. 

	CODE:

		// here author represents struct -- data structure
		type author struct {
			name string
			branch string
		}

		// here *a.branch = res.branch 
		// so *a = res
		// if *a = res then a = &res
		// so *author refers to &res ???


		func (a *author) show_1 (abranch string) {
			(*a).branch = abranch
		}



		func main() {
			
			// initialising values for data structure author
			
			res := author{
				name : "ABC",
				branch : "MECH",
			}
			fmt.Println("Branch name before:", res.branch)
			
			// calling show_1 -- string "ECE" from is passed to the parameter and *a points to the original res variable 
			res.show_1("ECE")
			
			// changes are made permanently
			fmt.Println("Branch in res after res.show_1() is:", res.branch)
			
			//using &res
			p:= &res
			p.show_1("CIVIL")
			fmt.Println("Branch in res after res.show_1() and using &res is:", res.branch)
			
			
		}


1) UNDERSCORE: 

	_(underscore) in Golang is known as the Blank Identifier. 
	Identifiers are the user-defined name of the program components used for the identification purpose. 
	Golang has a special feature to define and use the unused variable using Blank Identifier. 
	Unused variables are those variables which are defined by the user throughout the program but he/she never 
	make the use of these variables. These variables make the program almost unreadable. 
	As you know, Golang is more concise and readable programming language so it doesn’t allow the programmer 
	to define an unused variable if you do such, then the compiler will throw an error.
	
	The real use of Blank Identifier comes when a function returns multiple values, 
	but we need only a few values and want to discard some values. 
	Basically, it tells the compiler that this variable is not needed and ignored it without any error. 
	It hides the variable’s values and makes the program readable. 
	So whenever you will assign a value to Bank Identifier it becomes unusable.


2) CODE:

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
		
		mul, _ := maths(a,b)
		fmt.Println("mul is :", mul)
	}


3) BLANK IDENTIFIER:

	Important Points:

    You can use multiple Blank Identifier in the same program. 
    So you can say a Golang program can have multiple variables using the same identifier name 
    which is the blank identifier.
    There are many cases that arise the requirement of assignment of values just to complete the syntax 
    even knowing that the values will not be going to be used in the program anywhere. 
    Like a function returning multiple values. Mostly blank identifier is used in such cases.
    You can use any value of any type with the Blank Identifier.

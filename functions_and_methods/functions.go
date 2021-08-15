1) FUNCTIONS:

	A function is a collection of statements that perform some specific task and return the result to the caller. 
	A function can also perform some specific task without returning anything.
	A function can be declared even after calling 

2) Function Declaration

	Function declaration means a way to construct a function.
	
	SYNTAX: 

		func function_name(Parameter-list)(Return_type){
		    // function body.....
		    return ...
		}

		> no need to mention return type if return statement not used -- else error


	CODE:

		func add (a, b int) int {
			sum := a+b
			return sum
		}


	Here (a,b int) - means a and b are both int 

	
3) The declaration of the function contains:
 

    func: It is a keyword in Go language, which is used to create a function.
    function_name: It is the name of the function.
    Parameter-list: It contains the name and the type of the function parameters.
    Return_type: It is optional and it contain the types of the values that function returns. 
    If you are using return_type in your function, then it is necessary to use a return statement in your function


4) FUNCITON CALLING:

	Here function called after declaration

	func main() {
		fmt.Println("Hello, playground")
		fmt.Println(add(1,2))
	}

	func add (a, b int) int {
		sum := a+b
		return sum
	}

5) FUNCTION ARGUMENTS:

	In Go language, the parameters passed to a function are called actual parameters, 
	whereas the parameters received by a function are called formal parameters.

	Note: 
		By default Go language use call by value method to pass arguments in function.

	
6) CALL BY VALUE:

	In this parameter passing method, values of actual parameters are copied to function’s formal parameters 
	and the two types of parameters are stored in different memory locations. 
	So any changes made inside functions are not reflected in actual parameters of the caller.

	func main() {
		a:=2
		b:=3 
		fmt.Println(swap(a,b))
		fmt.Println(a,b)
	}

	func swap (a, b int) int {
		c:=a
		a = b
		b = c
		return a
	}

	Here only one value is returned -- How to return multiple values 


7) CALL BY REFERENCE:

	
	Both the actual and formal parameters refer to the same locations, 
	so any changes made inside the function are actually reflected in actual parameters of the caller.

	func main() {
		p:=2
		q:=3 
		fmt.Println(swap(&p,&q))
		fmt.Println(p,q)
	}

	func swap (a, b *int) int {
		c:= *a
		*a = *b
		*b = c
		return c
	}




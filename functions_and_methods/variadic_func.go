1) VARIADIC FUNCTIONS:

	The function that called with the varying number of arguments is known as variadic function. 
	Or in other words, a user is allowed to pass zero or more arguments in the variadic function. 
	fmt.Printf is the example of the variadic function, it required one fixed argument at the starting after that it can accept any number of arguments. 

2) Important Points:

    a) In the declaration of the variadic function, the type of the last parameter is preceded by an ellipsis, i.e, (…).
    b) (...) - three dots used 
    c) It indicates that the function can be called at any number of parameters of this type.

    SYNTAX:

	    function function_name(para1, para2...type)type{
	    // code...
	    }

    Inside the function …type behaves like a slice. For example, suppose we have a function signature, i.e, add( b…int)int, now the a parameter of type[]int.
    You can also pass an existing slice in a variadic function. To do this, we pass a slice of the complete array to the function as shown in Example 2 below.
    When you do not pass any argument in the variadic function, then the silce inside the function is nil.
    The variadic functions are generally used for string formatting.
    You can also pass multiple slice in the variadic function.
    You can not use variadic parameter as a return value, but you can return it as a slice.

3) CODE:

	package main

	import (
		"fmt"
		"strings"
	)

	func joinstr(elements...string) string{
		return strings.Join(elements, "-")
	}

	func main(){
		fmt.Printf("Zero argument --- %s", joinstr())
		
		fmt.Println(joinstr("Goku", "Gohan", "Vegeta", "Xander"))
		fmt.Println(joinstr("mystic", "spd"))
	}


4) USING SLICE:

	import (
	"fmt"
	"strings"
	)

	func joinstr(elements...string) string{
		return strings.Join(elements, "-")
	}

	func main(){
		fmt.Printf("Zero argument --- %s", joinstr())
		
		//pass a slice to a function
		
		element := []string{"Apollo", "Hercules", "Anaxagoras"}
		// not that element is used not elements 
		// element is argument
		// elements are paramaters
		fmt.Println(joinstr(element...))
	}



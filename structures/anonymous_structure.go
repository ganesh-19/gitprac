1) STRUCTURE:

	A structure or struct in Golang is a user-defined type, which allows us to create a group of elements of 
	different types into a single unit. Any real-world entity which has some set of properties or fields can be 
	represented as a struct.

2) ANONYMOUS STRUCTURE:

	In Go language, you are allowed to create an anonymous structure. 
	An anonymous structure is a structure which does not contain a name. 
	It useful when you want to create a one-time usable structure.

	You can create an anonymous structure using the following syntax:

		variable_name := struct{
		// fields
		}{// Field_values}


3) CODE: 
	
	package main

	import (
		"fmt"
	)

	func main() {
		Pokemon := struct{
	    		name   string
	    		pokemon_type string
		} {
			name : "Pikachu",
			pokemon_type : "Electric",
		}
		fmt.Println(Pokemon)
	}


4) ANONYMOUS FIELDS:

	Variables inside the structure are called fields.

	In a Go structure, you are allowed to create anonymous fields. 
	Anonymous fields are those fields which do not contain any name you just simply mention the type of the 
	fields and Go will automatically use the type as the name of the field. 

	You can create anonymous fields of the structure using the following syntax:

		type struct_name struct{
		    int
		    bool
		    float64
		}


	a) IMPORTANT POINTS:

		In a structure, you are not allowed to create two or more fields of the same type as shown below:

		type student struct{
		int
		int
		}

		If you try to do so, then the compiler will give an error.
		You are allowed to combine the anonymous fields with the named fields as shown below:

		type student struct{
		 name int
		 price int
		 string
		}


5) CODE:

	package main

	import (
		"fmt"
	)

	type fighters struct{
	    		  string
	    		  int
		} 
		
	func main() {
		
		goku := fighters{"goku", 10000}
		vegeta := fighters {"vegeta", 9500}
		
		fmt.Println("goku's power level:", goku.int)
		fmt.Println("vegeta's power level:", vegeta.int)
	}


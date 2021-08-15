1) STRUCTURES:

	A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different 
	types into a single type. Any real-world entity which has some set of properties/fields can be represented as a 
	struct. This concept is generally compared with the classes in object-oriented programming. 
	It can be termed as a lightweight class that does not support inheritance but supports composition. 

	For Example, an address has a name, street, city, state, Pincode. 
	It makes sense to group these three properties into a single structure address as shown below.

2) Declaring a structure:

	 type Address struct {
	      name string 
	      street string
	      city string
	      state string
	      Pincode int
	}
	 
	In the above, the type keyword introduces a new type. 
	It is followed by the name of the type (Address) and the keyword struct to illustrate that we’re defining a struct. 
	The struct contains a list of various fields inside the curly braces. Each field has a name and a type.


	Note: We can also make them compact by combining the various fields of the same type as shown in 
	the below example:


	type Address struct {
	    name, street, city, state string
	    Pincode int
	}


3) To Define a structure: 

	The syntax for declaring a structure:

		var a Address

	The above code creates a variable of a type Address which is by default set to zero. 
	For a struct, zero means all the fields are set to their corresponding zero value. 
	So the fields name, street, city, state are set to “”, and Pincode is set to 0

	You can also initialize a variable of a struct type using a struct literal as shown below:

		var a = Address{"Akshay", "PremNagar", "Dehradun", "Uttarakhand", 252636}

	Note:

	    Always pass the field values in the same order in which they are declared in the struct. 
	    Also, you can’t initialize only a subset of fields with the above syntax.
	    Go also supports the name: value syntax for initializing a struct 
	    (the order of fields is irrelevant when using this syntax). 
	    And this allows you to initialize only a subset of fields. 
	    All the uninitialized fields are set to their corresponding zero value

	    Example:

        var a = Address{Name:”Akshay”, street:”PremNagar”, state:”Uttarakhand”, Pincode:252636} //city:””

4) CODE:

	type address struct{
		name string
		city string
		pincode string
	}

	func main() {
		//declaring a variable of 'struct' type
		// all struct fields are initialised to zero
		
		var a address
		fmt.Println("variable a with address structure is:", a)
		
		a1  := address{"Sam", "London", "2354546"}
		fmt.Println("variable a1 with address structure is:", a1)
		
		a2 := address{name: "Sam", city: "Paris", pincode: "464757"}
		fmt.Println("variable a2 with address structure is:", a2)
		
		a3 := address{name: "Sam"}
		fmt.Println(a3)
	}


5) HOW TO ACCESS FIELDS OF A STRUCT:


	To access individual fields of a struct you have to use dot (.) operator.

	CODE: 


		type address struct{
		name string
		city string
		pincode string
	}

	func main() {
		a := address{name: "Sam", city:"London", pincode:"42366868"}
		
		fmt.Println(a.city, a.name)
		
		a.city = "Istanbul"
		
		fmt.Println(a)	
	}

	
6) POINTERS TO STRUCT: 

	Pointers in Go programming language or Golang is a variable which is used to store the 
	memory address of another variable

	CODE 1:

		package main

		import (
			"fmt"
		)

		type address struct{
			name string
			city string
			pincode string
		}

		func main() {
			a := &address{name: "Sam", city:"London", pincode:"42366868"}
			
			fmt.Println((*a).city, (*a).name)
			
			(*a).city = "Istanbul"
			
			fmt.Println(*a)	
		}


	CODE 2:

		package main

		import (
			"fmt"
		)

		type address struct{
			name string
			city string
			pincode string
		}

		func main() {
			a := &address{name: "Sam", city:"London", pincode:"42366868"}
			
			fmt.Println(a.city, a.name)
			
			a.city = "Istanbul"
			
			fmt.Println(a)	
		}


1) NESTED STRUCTURE:

	A structure or struct in Golang is a user-defined type, which allows us to create a group of elements of 
	different types into a single unit. Any real-world entity which has some set of properties or fields 
	can be represented as a struct. Go language allows nested structure. 
	A structure which is the field of another structure is known as Nested Structure. 
	Or in other words, a structure within another structure is known as a Nested Structure.


2) Syntax:

	type struct_name_1 struct{
	  // Fields
	} 
	type struct_name_2 struct{
	  variable_name  struct_name_1

	}

3) CODE:

	package main

	import (
		"fmt"
	)

	type address struct{
		name string
		city string
		pincode string
	}

	type home struct{
		//  variable referring to structure address
		nested_var address
	}

	func main() {
		a := home{
			// variable referring to structure address
			nested_var: address{name: "Sam", city:"London", pincode:"42366868"},
		}
		fmt.Println("\nhome address:")
		
		fmt.Println(a)
	}


4) CODE 2:

	package main

	import (
		"fmt"
	)

	type Student struct {
	    name   string
	    branch string
	    year   int
	}
	  
	// Creating nested structure
	type Teacher struct {
	    name    string
	    subject string
	    exp     int
	    details Student
	} 

	func main() {
		result := Teacher { name :"Suman", subject: "golang", exp:5, 
			// variable referring to structure address
			// structure within a structure compressed information
			details: Student {name: "Sam", branch:"ECE", year: 2018},
		}
		fmt.Println("\nDetails:")
		
		fmt.Println(result)
		
		// Display the values
	    	fmt.Println("Details of the Teacher")
	    	fmt.Println("Teacher's name: ", result.name)
	   	fmt.Println("Subject: ", result.subject)
	    	fmt.Println("Experience: ", result.exp)
		
	      	fmt.Println("\nDetails of Student")
	    	fmt.Println("Student's name: ", result.details.name)
	        fmt.Println("Student's branch name: ", result.details.branch)
	    	fmt.Println("Year: ", result.details.year) 
	}
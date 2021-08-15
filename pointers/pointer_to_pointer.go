1) "DOUBLE POINTER":
	Pointers in Go programming language or Golang is a variable which is used to store the 
	memory address of another variable. A pointer is a special variable so it can point to a 
	variable of any type even to a pointer. Basically, this looks like a chain of pointers. 
	When we define a pointer to pointer then the first pointer is used to store the address 
	of the second pointer. This concept is sometimes termed as Double Pointers.

2) "How to Declare a pointer to pointer in Golang?":
	Declaring Pointer to Pointer is similar to declaring pointer in Go. The difference is 
	we have to place an additional ‘*’ before the name of pointer name. This is generally done 
	when we are declaring the pointer variable using the var keyword along with the type. 
	Below example and image will explain the concept in much better way.

3) "Example":
		type Employee struct {
		
			// taking variables
			name  string
			empid int
		} 

		func main(){
			emp1 := Employee{name : "Sam", empid : 100}
			emp2 := Employee{"Jerry",102}
			
			emp_mem1 := &emp1
			
			//cannot declare pointer like this for struct as the data type is not int
			//var emp_mem4 *int = &emp2
			emp_mem4 := &emp2
			
			emp_mem2 := &emp_mem1
			
			//./prog.go:22:6: cannot use &emp_mem4 (type **Employee) as type **int in assignment
			//var emp_mem3 **int = &emp_mem4
			emp_mem3 := &emp_mem4
			
			fmt.Println(emp1)
			fmt.Println(emp2)
			fmt.Println(emp_mem1)
			fmt.Println(emp_mem2)
			fmt.Println(emp_mem3)
			fmt.Println(emp_mem4)
			fmt.Println(*emp_mem1)
			fmt.Println(*emp_mem2)
			fmt.Println(*emp_mem3)
			fmt.Println(*emp_mem4)	
		}
	
	'OUTPUT':
		{Sam 100}
		{Jerry 102}
		&{Sam 100}
		0xc000102018
		0xc000102020
		&{Jerry 102}
		{Sam 100}
		&{Sam 100}
		&{Jerry 102}
		{Jerry 102}
	

4) "EXAMPLE2":

	import "fmt"
	
	// Main Function
	func main() {
	
		// taking a variable
		// of integer type
		var v int = 100
	
		// taking a pointer
		// of integer type
		var pt1 *int = &v
	
		// taking pointer to
		// pointer to pt1
		// storing the address
		// of pt1 into pt2
		var pt2 **int = &pt1
	
		fmt.Println("The Value of Variable v is = ", v)
	
		// changing the value of v by assigning
		// the new value to the pointer pt1
		*pt1 = 200
	
		fmt.Println("Value stored in v after changing pt1 = ", v)
	
		// changing the value of v by assigning
		// the new value to the pointer pt2
		**pt2 = 300
	
		fmt.Println("Value stored in v after changing pt2 = ", v)
	} 

	'OUTPUT':
		The Value of Variable v is =  100
		Value stored in v after changing pt1 =  200
		Value stored in v after changing pt2 =  300
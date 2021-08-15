1) "POINTERS TO STRUCT":
	Pointers in Golang is a variable which is used to store the memory address of another variable. 
	Pointers in Golang is also termed as the special variables. The variables are used to store some 
	data at a particular memory address in the system.

	You can also use a pointer to a struct. A struct in Golang is a user-defined type which allows 
	to group/combine items of possibly different types into a single type. To use pointer to a struct, 
	you can use & operator i.e. address operator. Golang allows the programmers to access the fields of 
	a structure using the pointers without any dereferencing explicitly.

2) "Example": 
	Here, we are creating a structure named Employee which has two variables. In the main function, 
	create the instance of the struct i.e. emp. After that, you can pass the address of the struct 
	to the pointer which represents the pointer to struct concept. There is no need to use 
	dereferencing explicitly as it will give the same result as you can see in the below program 
	(two times ABC).

		type Employee struct {
	
			// taking variables
			name  string
			empid int
		} 
		
		func main(){
			emp1 := Employee{name : "Sam", empid : 100}
			
			emp_mem := &emp1 
			
			fmt.Println(emp_mem, emp_mem.name)
			// *emp_mem.name gives error
			fmt.Println(*emp_mem, (*emp_mem).name)
		}
	
	"output":
		&{Sam 100} Sam
		{Sam 100} Sam

3) You can also modify the values of the structure members or structure literals 
	by using the pointer

	"CODE":
		type Employee struct {
	
			// taking variables
			name  string
			empid int
		} 
		
		func main(){
			emp1 := Employee{name : "Sam", empid : 100}
			
			emp_mem := &emp1 
			fmt.Println(emp_mem, emp_mem.name)
			fmt.Println(*emp_mem, (*emp_mem).name)
			
			emp_mem.name = "Mary"
			fmt.Println(emp_mem, emp_mem.name)
			fmt.Println(*emp_mem, (*emp_mem).name)
			
			(&emp1).name = "Joe"
			fmt.Println(emp_mem, emp_mem.name)
			fmt.Println(*emp_mem, (*emp_mem).name)
			
			(*emp_mem).name = "Tom"
			fmt.Println(emp_mem, emp_mem.name)
			fmt.Println(*emp_mem, (*emp_mem).name)
		}
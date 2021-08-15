1) "COMPARING POINTERS IN GOLANG":
	Pointers in Go programming language or Golang is a variable which is used to store the 
	memory address of another variable. Pointers in Golang are also termed as the special variables.
	The variables are used to store some data at a particular memory address in the system. 
	The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).
	In Go language, you are allowed to compare two pointers with each other. Two pointers values 
	are only equal when they point to the same value in the memory or if they are nil. 
	You can perform a comparison on pointers with the help of == and != operators provided by 
	the Go language

2) "== OPERATOR":
	This operator return true if both the pointer points to the same variable. 
	Or return false if both the pointer points to different variables. 

	'Syntax':
		pointer_1 == pointer_2

3) "!= OPERATOR":
	This operator return false if both the pointer points to the same variable. 
	Or return true if both the pointer points to different variables. 

	'Syntax':
		pointer_1 != pointer_2

	'CODE':
		type Employee struct {
	
			// taking variables
			name  string
			empid int
		} 
		
		func main(){
		
			emp1 := Employee{"Jerry",102}
			emp_mem1 := &emp1
			emp_mem2 := &emp1
			
			fmt.Println(emp_mem1 == emp_mem2)
			fmt.Println(emp_mem1 != emp_mem2)
			fmt.Println(&emp_mem1 == &emp_mem2)
			fmt.Println(&emp_mem1 != &emp_mem2)
			fmt.Println(*emp_mem1 == *emp_mem2)
			fmt.Println(*emp_mem1 != *emp_mem2)		
		}
		


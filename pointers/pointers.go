1) "POINTERS":
	Pointers in Go programming language or Golang is a variable which is used to store the memory 
	address of another variable. Pointers in Golang is also termed as the special variables. 
	The variables are used to store some data at a particular memory address in the system. 
	The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).

2) "What is the need of the pointers?":
	To understand this need, first, we have to understand the concept of variables. 
	Variables are the names given to a memory location where the actual data is stored. 
	To access the stored data we need the address of that particular memory location. 
	To remember all the memory addresses(Hexadecimal Format) manually is an overhead 
	that’s why we use variables to store data and variables can be accessed just by using their name.
	
	Golang also allows saving a hexadecimal number into a variable using the literal expression 
	i.e. number starting from 0x is a hexadecimal number.

	'Example': In the below program, we are storing the hexadecimal number into a variable. 
	But you can see that the type of values is int and saved as the decimal or you can say the 
	decimal value of type int is storing. But the main point to explain this example is that 
	we are storing a hexadecimal value(consider it a memory address) but it is not a pointer 
	as it is not pointing to any other memory location of another variable. 
	It is just a user-defined variable. So this arises the need for pointers.

	'CODE':
		// storing the hexadecimal
		// values in variables
		x := 0xFF
		y := 0x9C
		
		// Displaying the values
		fmt.Printf("Type of variable x is %T\n", x)
		fmt.Printf("Value of x in hexdecimal is %X\n", x)
		fmt.Printf("Value of x in decimal is %v\n", x)
		
		fmt.Printf("Type of variable y is %T\n", y)
		fmt.Printf("Value of y in hexdecimal is %X\n", y)
		fmt.Printf("Value of y in decimal is %v\n", y)

	"OUTPUT":
		Type of variable x is int
		Value of x in hexdecimal is FF
		Value of x in decimal is 255
		Type of variable y is int
		Value of y in hexdecimal is 9C
		Value of y in decimal is 156

3) "WHAT IS A POINTER?":
	A pointer is a special kind of variable which is not only used to store the memory addresses of other variables 
	but it also points where the memory is located and provides the ways to find out the value stored at that memory 
	location. It is generally termed as Special kind of Variable because it is almost declared as a variable but with 
	*(dereferencing operator).

4) "DECLARATION OF A POINTER":
	Before we start there are two important operators which we will use in pointers i.e.
	* Operator also termed as the dereferencing operator used to declare pointer variable and access the value 
	stored in the address.
	& operator termed as address operator used to returns the address of a variable or to access the 
	address of a variable to a pointer.

	'Declaring a pointer':

		var pointer_name *Data_Type

	'Example': 
		Below is a pointer of type string which can store only the memory addresses of string variables.

		var s *string

5) "Initialization of Pointer":
	To do this you need to initialize a pointer with the memory address of another variable using the 
	address operator as shown in the below example:

		// normal variable declaration
		var a = 45

		// Initialization of pointer s with 
		// memory address of variable a
		var s *int = &a

		or  var s *int 
			s = &a


6) "EXAMPLE":
		a := 3 
		ranger := "blue"

		var a_mem *int = &a 
		var ran_mem *string 
		ran_mem = &ranger

		fmt.Println(a_mem, ran_mem)
		fmt.Println(*a_mem, *ran_mem)

7) "IMPORTANT POINTS":


7A) The default value or zero-value of a pointer is always nil. Or you can say that an uninitialized pointer will 
	always have a nil value. 

	"example":
	var a_mem *int
	var ran_mem *string 

	fmt.Println(a_mem, ran_mem)
	fmt.Println(*a_mem, *ran_mem)

	'output':
	 <nil> <nil>

7B) Declaration and initialization of the pointers can be done into a single line.

	Example:
		var s *int = &a

7C) If you are specifying the data type along with the pointer declaration then the pointer will be able to 
	handle the memory address of that specified data type variable. For example, if you taking a pointer of 
	string type then the address of the variable that you will give to a pointer will be only of string data 
	type variable, not any other type.

7D) To overcome the above mention problem you can use the Type Inference concept of var keyword. 
	There is no need to specify the data type of during the declaration. The type of a pointer variable can 
	also be determined by the compiler like a normal variable. Here we will not use the * operator. 
	It will internally determine by the compiler as we are initializing the variable with the address of 
	another variable.

7E) You can also use the shorthand (:=) syntax to declare and initialize the pointer variables. 
	The compiler will internally determine the variable is a pointer variable if we are passing the address of 
	variable using &(address) operator to it. 

	"Example":
		a, name := 3, "nick"
		
		a_mem := &a
		var ran_mem = &name

		fmt.Println(a_mem, ran_mem)
		fmt.Println(*a_mem, *ran_mem)

8) "DEREFERENCING THE POINTER":

	As we know that * operator is also termed as the dereferencing operator. 
	It is not only used to declare the pointer variable but also used to access the value stored in the 
	variable which the pointer points to which is generally termed as indirecting or dereferencing. 
	* operator is also termed as the value at the address of. 

	"Example":
		a, name := 3, "nick"
		
		a_mem := &a
		var ran_mem = &name

		fmt.Println(a_mem, ran_mem)
		fmt.Println(*a_mem, *ran_mem)

9) You can also change the value of the pointer or at the memory location 
	instead of assigning a new value to the variable.

	'Example':

		a, name := 3, "nick"
		
		a_mem := &a
		var ran_mem = &name

		fmt.Println(a_mem, ran_mem)
		fmt.Println(*a_mem, *ran_mem)
		
		*ran_mem = "Xander"
		fmt.Println(*a_mem, name)
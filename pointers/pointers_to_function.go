1) "POINTERS TO FUNCTION":
	Pointers in Go programming language or Golang is a variable which is used to store the memory 
	address of another variable. You can also pass the pointers to the function like the variables. 
	There are two ways to do this as follows:

		> Create a pointer and simply pass it to the function
		> Passing an address of the variable

2) "Create a pointer and simply pass it to the Function":
		func ranger(name *string) string{
			return *name + "ranger"
		}
			
		func main() {
			name := "nick"
			var ran_mem = &name
			
			x := ranger(ran_mem)
			fmt.Println(x)	
			fmt.Println(name)
		}

		'output':
			nickranger
			nick

3) "Passing an address of the variable to the function call":
		func ranger(name *string) {
			*name = "Xander"
		}
			
		func main() {
			name := "nick"
			var ran_mem = &name
			fmt.Println(name)
			ranger(ran_mem)
			fmt.Println(name)
		}

		"output":
			nick
			Xander



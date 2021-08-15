1) The Go language reserve two functions for special purpose and the functions are main() and init() function.

2) main() function

	In Go language, the main package is a special package which is used with the programs that are executable 
	and this package contains main() function. 
	The main() function is a special type of function and it is the entry point of the executable programs. 
	It does not take any argument nor return anything. Go automatically call main() function, 
	so there is no need to call main() function explicitly and every executable program must contain 
	single main package and main() function.

	CODE:

		import (
			"fmt"
			"strings"
			"sort"
			"time"
		)

		func main(){

			s := []int{10,20,30,40,50,60,70,80,90}
			
			sort.Ints(s)
			
			fmt.Println("sorted slice is", s)
			
			//finding index
			fmt.Println("Index value", strings.Index("Dragonballz", "l"))
			
			//finding time
			fmt.Println("Time", time.Now().Unix())
		}


3) init() Function

	init() function is just like the main function, does not take any argument nor return anything. 
	This function is present in every package and this function is called when the package is initialized. 

	This function is declared implicitly, so you cannot reference it from anywhere and you are allowed 
	to create multiple init() function in the same program and they execute in the order they are created. 

	You are allowed to create init() function anywhere in the program and they are called in lexical 
	file name order (Alphabetical Order). And allowed to put statements if the init() function, 
	but always remember to init() function is executed before the main() function call, 
	so it does not depend to main() function. 

	The main purpose of the init() function is to initialize the global variables that cannot be initialized in the global context. 


	CODE:

		package main

		import (
			"fmt"
			"time"
		)

		var saiyan = "Vegeta"


		func main() {
			fmt.Println("main function")
			fmt.Println("time is:", time.Now().Unix())
			fmt.Println("saiyan in init1 is :", saiyan)
		}

		func init() {
			fmt.Println("init1")
			var saiyan = "Goku"
			fmt.Println("saiyan in init1 is :", saiyan)
		}

		func init() {
			fmt.Println("init2")
			fmt.Println("saiyan in init2 is:", saiyan)
		}

	
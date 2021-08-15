1) "GOROUTINES":
	Go language provides a special feature known as a Goroutines. A Goroutine is a function or method which executes 
	independently and simultaneously in connection with any other Goroutines present in your program. 
	Or in other words, every concurrently executing activity in Go language is known as a Goroutines. 
	You can consider a Goroutine like a light weighted thread. The cost of creating Goroutines is very small as 
	compared to the thread. Every program contains at least a single Goroutine and that Goroutine is known as the 
	main Goroutine. All the Goroutines are working under the main Goroutines if the main Goroutine terminated, 
	then all the goroutine present in the program also terminated. Goroutine always works in the background.

2) "How to create a Goroutine?":
	You can create your own Goroutine simply by using go keyword as a prefixing to the function or method call.

	'SYNTAX':
		func name(){
		// statements
		}
		// using go keyword as the 
		// prefix of your function call
		go name()

3) "CODE":

	'Code1':
	
	// Go program to illustrate
	// the concept of Goroutine
	package main
	
	import "fmt"
	
	func display(str string) {
		for w := 0; w < 6; w++ {
			fmt.Println(str)
		}
	}
	
	func main() {
	
		// Calling Goroutine
		go display("Welcome")
	
		// Calling normal function
		display("GeeksforGeeks")
	}

	"OUTPUT":
	
		GeeksforGeeks
		GeeksforGeeks
		GeeksforGeeks
		GeeksforGeeks
		GeeksforGeeks
		GeeksforGeeks
	
	In the above example, we simply create a display() function and then call this function in 
	two different ways first one is a Goroutine, i.e. go display(“Welcome”) and another one is 
	a normal function, i.e. display(“GeeksforGeeks”). But there is a problem, it only displays 
	the result of the normal function that does not display the result of Goroutine because when 
	a new Goroutine executed, the Goroutine call return immediately. The control does not wait 
	for Goroutine to complete their execution just like normal function they always move forward 
	to the next line after the Goroutine call and ignores the value returned by the Goroutine. 
	So, to executes a Goroutine properly, we made some changes in our program as shown in the below 
	code: 

4) "CODE2":
	import (
		"fmt"
		"time"
	)

	func display(str string){
		for i:=0; i<5; i++{
			time.Sleep(1 * time.Second)
			fmt.Println(str)
		}
	}
	func main(){
		greet := "Welcome"
		name := "Sam"
		
		go display(greet)
		display(name)
	}

"OUTPUT":
	Sam
	Welcome
	Welcome
	Sam
	Sam
	Welcome
	Welcome
	Sam
	Sam
	
We added the Sleep() method in our program which makes the main Goroutine sleeps for 1 second in 
between 1-second the new Goroutine executes, displays “welcome” on the screen, and then terminate 
after 1-second main Goroutine re-schedule and perform its operation. This process continues until 
the value of the z<6 after that the main Goroutine terminates. Here, both Goroutine and the normal 
function work concurrently.

5) "Advantages of Goroutines":

	>Goroutines are cheaper than threads.
	>Goroutine are stored in the stack and the size of the stack can grow and shrink according 
	to the requirement of the program. But in threads, the size of the stack is fixed.
	>Goroutines can communicate using the channel and these channels are specially designed to 
	prevent race conditions when accessing shared memory using Goroutines.
	>Suppose a program has one thread, and that thread has many Goroutines associated with it. 
	If any of Goroutine blocks the thread due to resource requirement then all the remaining 
	Goroutines will assign to a newly created OS thread. All these details are hidden from the 
	programmers.

6) "Anonymous Goroutine":

	In Go language, you can also start Goroutine for an anonymous function or in other words, 
	you can create an anonymous Goroutine simply by using go keyword as a prefix of that function 
	as shown in the below Syntax:
	
	'Syntax':
	
		// Anonymous function call
		go func (parameter_list){
		// statement
		}(arguments)
	
	'Example':
		import (
			"fmt"
			"time"
		)
		
		// Main function
		func main() {
		
			fmt.Println("Welcome!! to Main function")
		
			// Creating Anonymous Goroutine
			go func() {
		
				fmt.Println("Welcome!! to GeeksforGeeks")
		//proper way of calling anonymous function
			}()
		
			time.Sleep(1 * time.Second)
			fmt.Println("GoodBye!! to Main function")
		// wrong and cannot be done
		//  ()
			fmt.Println("GoodBye!! to Main function")
		} 

1) "MULTIPLE GOROUTINE":
	A Goroutine is a function or method which executes independently and simultaneously in connection with any 
	other Goroutines present in your program. Or in other words, every concurrently executing activity in Go 
	language is known as a Goroutines. In Go language, you are allowed to create multiple goroutines in a single 
	program. You can create a goroutine simply by using go keyword as a prefixing to the function or method call.

2) "SYNTAX":
		func name(){
		// statements
		}
		// using go keyword as the 
		// prefix of your function call
		go name()

3) "CODE":
	
	// try with and without time
		package main

		import (
			"fmt"
			"time"
		)

		//go routine 1
		func name (){
			arr1 := [4]string{"Zeus", "Athena", "Euclid", "Pythagoras"}
			time.Sleep(100 * time.Millisecond)
			
			for i:= range arr1 {
				fmt.Println("routine name called:", i)
			}
		}

		// go routine 2
		func state(){
			arr2 := [4] string {"Sparta", "Troy", "Athens", "Macedonia"}
			time.Sleep(200* time.Millisecond)
			
			for i:= range arr2 {
				fmt.Println("routine state called:", i)
			}
		}

		func main(){
			fmt.Println("!...Main Go-routine Start...!")
			fmt.Println("routine 1 called")
			go name()
			fmt.Println("routine 2 called")
			go state()
			time.Sleep(3500 * time.Millisecond)
				fmt.Println("Main end") 

		}

4) "Working": 
	> Here, we have two goroutines, i.e, name and state and one main goroutine. 
	> When we run this program first the main goroutine strat and print “!...Main Go-routine Start...!“
	> here the main goroutine is like a parent and other goroutines are its children 
	> so first main goroutine run after that those other goroutines start 
	> if the main goroutine terminates, then other goroutines are also terminated.
	> after the main goroutine, name and state goroutines start their working concurrently. 
	> name goroutine starts it working from 100 Millisecond and state start its working from 200 Millisecond 
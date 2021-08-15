1) "SELECT":
	In Go language, the select statement is just like switch statement, but in the select statement,
	case statement refers to communication, i.e. sent or receive operation on the channel. 

2) "Syntax":

	select{

	case SendOrReceive1: // Statement
	case SendOrReceive2: // Statement
	case SendOrReceive3: // Statement
	.......
	default: // Statement

3) "Important points":

3A) Select statement waits until the communication(send or receive operation) is prepared for 
	some cases to begin. 

	package main
	// Main function
	func main() {
		//select statement without any case
		select {}
	} 

	//error:
	// fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [select (no cases)]:
	// main.main()
	// 	/tmp/sandbox518147089/prog.go:8 +0x25

3B) The default statement in the select statement is used to protect select statement from blocking. 
	This statement executes when there is no case statement is ready to proceed.

	
// Go program to illustrate the
// concept of select statement
package main
  
import "fmt"
  
// main function
func main() {
      
    // creating channel
    mychannel:= make(chan int)
   select{
     case <- mychannel: 
       
     default:fmt.Println("Not found")
}   
}

// Output:

// Not found

The blocking of select statement means when there is no case statement is ready and the select 
statement does not contain any default statement, then the select statement block until at least 
one case statement or communication can proceed.

Example:

// Go program to illustrate the
// concept of select statement
package main
  
// main function
func main() {
      
    // creating channel
    mychannel:= make(chan int)
      
    // channel is not ready 
   // and no default case
   select{
       case <- mychannel:
         
   }
}

// Output:

// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan receive]:
// main.main()
//     /home/runner/main.go:14 +0x52
// exit status 2

In select statement, if multiple cases are ready to proceed, then one of them can be selected 
randomly.

Example:

// Go program to illustrate the
// concept of select statement
package main
  
import "fmt"
      
      
    // function 1
    func portal1(channel1 chan string){
        for i := 0; i <= 3; i++{
            channel1 <- "Welcome to channel 1"
        }
          
    }
      
    // function 2
     func portal2(channel2 chan string){
        channel2 <- "Welcome to channel 2"
    }
  
// main function
func main() {
      
    // Creating channels
   R1:= make(chan string)
   R2:= make(chan string)
     
   // calling function 1 and 
   // function 2 in goroutine
   go portal1(R1)
   go portal2(R2)
     
   // the choice of selection 
   // of case is random
   select{
       case op1:= <- R1:
       fmt.Println(op1)
       case op2:= <- R2:
       fmt.Println(op2)
   }
}

// Output:

// Welcome to channel 2
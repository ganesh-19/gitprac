1) "CHANNEL":
	In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this 
	communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to 
	send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive 
	data through the same channel

2) "Creating a Channel":

	In Go language, a channel is created using chan keyword and it can only transfer data of the same type, 
	different types of data are not allowed to transport from the same channel.

	'Syntax':
		var Channel_name chan Type
	You can also create a channel using make() function using a shorthand declaration.
	'Syntax':
		channel_name:= make(chan Type)

	"Example":
		func main(){
		
			// using var keyword
			var mychannel chan int
			
			fmt.Println("value of channel is:", mychannel)
			fmt.Printf("\ntype of channel is: %T", mychannel)
			
			//using make()
			
			mychannel1 := make(chan int)
			fmt.Println("\nvalue of channel using make is:", mychannel1)
			fmt.Printf("\nTYpe of channel using make is %T", mychannel1)
		}

		'output':
			// value of channel is: <nil>

			// type of channel is: chan int
			// value of channel using make is: 0xc0000a4000
			
			// TYpe of channel using make is chan int

3) "Send and Receive Data From a Channel":

	In Go language, channel work with two principal operations one is sending and another one is receiving, 
	both the operations collectively known as communication. And the direction of <- operator indicates whether 
	the data is received or send. In the channel, the send and receive operation block until another side is not 
	ready by default. It allows goroutine to synchronize with each other without explicit locks or condition 
	variables.

		'Send operation': 
		The send operation is used to send data from one goroutine to another goroutine with the help of a channel. 
		Values like int, float64, and bool can safe and easy to send through a channel because they are copied so 
		there is no risk of accidental concurrent access of the same value. Similarly, strings are also safe to 
		transfer because they are immutable. But for sending pointers or reference like a slice, map, etc. through 
		a channel are not safe because the value of pointers or reference may change by sending goroutine or by the 
		receiving goroutine at the same time and the result is unpredicted. So, when you use pointers or references 
		in the channel you must make sure that they can only access by the one goroutine at a time.

			Mychannel <- element

		The above statement indicates that the data(element) send to the channel(Mychannel) with the help of a <- 
		operator.

		'Receive operation': 
			The receive operation is used to receive the data sent by the send operator.

			element := <-Mychannel

		The above statement indicates that the element receives data from the channel(Mychannel). 
		If the result of the received statement is not going to use is also a valid statement. 
		You can also write a receive statement as:

			<-Mychannel

	"CODE":
		func summ(ch chan int){
			fmt.Println(100+ <- ch)
			//return 100+ <- ch
		}
		
		func main(){
			
			// using var keyword
			// when using this --- get error --- deadlock
			//var ch chan int
			ch := make(chan int)
			
			// go routine cannot return
			// x := go summ (ch)
			go summ (ch)
			ch <- 13
			fmt.Println("added to 100")
			
		}

4) "CLOSING A CHANNEL":
	You can also close a channel with the help of close() function. 
	This is an in-built function and sets a flag which indicates that no more value will send to this channel.

	'Syntax':

		close()

	You can also close the channel using for range loop. Here, the receiver goroutine can check the channel is 
	open or close with the help of the given syntax:

		ele, ok:= <- Mychannel

	Here, if the value of ok is true which means the channel is open so, read operations can be performed. 
	And if the value of is false which means the channel is closed so, read operations are not going to perform.

		"CODE":
			package main

			import (
				"fmt"
			)
			
			func myfun (ch chan string) {
				for i :=0; i<5; i++ {
				// here channel connects with the channel in main function -- sort of like return 
					ch <- "dragonballz"
				}
				close(ch)	
			}
			
			func main() {
			// creating a channel
			ch := make(chan string)
			
			// go routine
			go myfun(ch)
			
			// When the value of ok is
				// set to true means the
				// channel is open and it
				// can send or receive data
				// When the value of ok is set to
				// false means the channel is closed 
			
			for {
				res, ok := <- ch
				if ok == false {
					fmt.Println("channel close", ok)
					break
				}
				
				fmt.Println("Channel open", res, ok)
				}
				
			}



		
1) "UNIDIRECTIONAL CHANNEL":
	As we know that a channel is a medium of communication between concurrently running goroutines 
	so that they can send and receive data to each other. By default a channel is bidirectional 
	but you can create a unidirectional channel also. A channel that can only receive data or a 
	channel that can only send data is the unidirectional channel. The unidirectional channel can 
	also create with the help of make() function

2) "SYNTAX":
	// Only to receive data
	c1:= make(<- chan bool)
	// Only to send data
	c2:= make(chan<-bool)

	'EXAMPLE':
		func main(){
		
			ch := make(<- chan string)
			ch1 := make(chan <- int)
			
			fmt.Printf("\n type of ch: %T", ch)
			fmt.Printf("\n type of ch1: %T", ch1)
		}

3) "Converting Bidirectional Channel into the Unidirectional Channel":
	In Go language, you are allowed to convert a bidirectional channel into the unidirectional channel, or in 
	other words, you can convert a bidirectional channel into a receive-only or send-only channel, but vice 
	versa is not possible.

	'CODE':
		  
		func sending(s chan<- string) {
			s <- "GeeksforGeeks"
		}
		
		func main() {
		
			// Creating a bidirectional channel
			mychanl := make(chan string)
		
			// Here, sending() function convert
			// the bidirectional channel
			// into send only channel
			go sending(mychanl)
		
			// Here, the channel is sent 
			// only inside the goroutine
			// outside the goroutine the 
			// channel is bidirectional
			// So, it print GeeksforGeeks
			fmt.Println(<-mychanl)
		} 

4) "Use of Unidirectional Channel": 
	The unidirectional channel is used to provide the type-safety of the program so, that the program gives 
	less error. Or you can also use a unidirectional channel when you want to create a channel that can 
	only send or receive data.

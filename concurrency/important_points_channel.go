1) "Important Points":

    >>Blocking Send and Receive: 
		In the channel when the data sent to a channel the control is blocked in that send statement 
		until other goroutine reads from that channel. Similarly, when a channel receives data from the goroutine 
		the read statement block until another goroutine statement.
    >>Zero Value Channel: 
		The zero value of the channel is nil.
2) "For loop in Channel":
	A for loop can iterate over the sequential values sent on the channel until it closed. 

		'Syntax':

		for item := range Chnl { 
			// statements..
		}

	"CODE":
		func main() {
			//create a channel using make function
			ch := make(chan string)
			
			//anonymous go routine
			go func() {
				ch <- "The edge"
				ch <- "Taxi driver"
				ch <- "Lord of the rings"
				ch <- "Children of the heaven"
				ch <- "The way home"
				
				// if close not used -- error
				// fatal error: all goroutines are asleep - deadlock!
				
				close(ch)
			} ()
			
			// using loop to loop through channel
			for i:= range ch {
				fmt.Println(i)
			}
		}

3) "Length of the Channel": 
	In channel, you can find the length of the channel using len() function. Here, the length indicates the 
	number of value queued in the channel buffer.
		
	"CODE":
		func main() {
			//create a channel using make function
			ch := make(chan string)
			
			//anonymous go routine
			go func() {
				ch <- "The edge"
				ch <- "Taxi driver"
				ch <- "Lord of the rings"
				ch <- "Children of the heaven"
				ch <- "The way home"
				
				//printing lenght of the channel here 
				fmt.Println("length inside routine is:", len(ch))
				// if close not used -- error
				// fatal error: all goroutines are asleep - deadlock!
				
				close(ch)
			} ()
			
			// using loop to loop through channel
			for i:= range ch {
				
				// when trying to print lenght of the channel inside loop
				fmt.Println(len(ch), i)
			}
			
			// without mentioning explicitly deadlock happens
			// even though close() used deadlock happens
			// but when used inside routine no problem
			// chn := make(chan string)
			
			// chn <- "Shinjaki no Kyogen"
			// chn <- "Mirai Nikki"
			// chn <- "Dragon ball z"
			// chn <- "Pokemon"
			// close(chn)
			// fmt.Println("without go routine:", len(chn))

			// problem is not about routine
			// should mention length in make
			
			chn1 := make(chan string, 4)
			
			chn1 <- "Shinjaki no Kyogen"
			chn1 <- "Mirai Nikki"
			chn1 <- "Dragon ball z"
			chn1 <- "Pokemon"
			close(chn1)
			fmt.Println("without go routine:", len(chn1))
		}

4) 'Capacity of the Channel': 
		In channel, you can find the capacity of the channel using cap() function. Here, the capacity indicates the 
		size of the buffer. 

	"CODE":
		func main(){

			chn1 := make(chan string, 8)
			
			chn1 <- "Shinjaki no Kyogen"
			chn1 <- "Mirai Nikki"
			chn1 <- "Dragon ball z"
			chn1 <- "Pokemon"
			close(chn1)
			fmt.Println("capacity:", cap(chn1))
			fmt.Println("length:", len(chn1))
		}

5) "Select and case statement in Channel": 
		In go language, select statement is just like a switch statement without any input parameter. 
		This select statement is used in the channel to perform a single operation out of multiple operations 
		provided by the case block.
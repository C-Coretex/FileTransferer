package main

import(
	 Output "fmt"
	 Input "bufio"
	 "os"
	 "net"
	 "strconv"
	 "strings"
)

const (
	//Send file to the server (client)
	Send = iota + 1
	//Get file from the client (server)
	Get
)

func main() {
	Output.Println("\nHey all, I'm the Go app")
	Output.Println();
	/*Output.Println("\nAnd what's your name?")
	reader := Input.NewReader(os.Stdin)
	
	text, _ := reader.ReadString('\n')
	Output.Println("Oh, it's nice to meet you, ", text)*/
	
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if(err != nil){
		Output.Println("Oops, we've got an error: " + err.Error() + "\n")
		os.Exit(1)
	}
	
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	
Repeat:
	Output.Println("{1} - Send a file (Client)")
	Output.Println("{2} - Get a file (Server)")
	reader := Input.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	option = strings.TrimRight(option, "\r\n")
	
	Output.Println("Your IP is:", localAddr)
	
	if(option == strconv.Itoa(Send)){
		Output.Println("You want to send file")
	} else if(option == strconv.Itoa(Get)){
		Output.Println("You want to get file")
	} else{
		goto Repeat
	}
	 
	 
	 Output.Println();
}
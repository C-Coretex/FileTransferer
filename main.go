package main

import(
	 Output "fmt"
	 Input "bufio"
	 "os"
	 "net"
	 "strconv"
	 "strings"
)

//FileExists gets the full path to the file
func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

const (
	//Send file to the server (client)
	Send = iota + 1
	//Get file from the client (server)
	Get
)

func main() {
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
		Output.Println("\nYou want to send file")
		Output.Println("Full path to the file you want to send")
		
		FileDoesNotExist:
		path, _ := reader.ReadString('\n')
		path = strings.TrimRight(path, "\r\n")
		path = strings.TrimLeft(path, "\r")
		if(!FileExists(path)){
			goto FileDoesNotExist
		}
		
		Output.Println("\nIP of the server")
		server, _ := reader.ReadString('\n')
		server = strings.TrimRight(server, "\r\n")
		
		Output.Println(path, " ", server)
		
	} else if(option == strconv.Itoa(Get)){
		Output.Println("\nYou want to get file")
		
	} else{
		goto Repeat
		
	}
	 
	Output.Println();
}
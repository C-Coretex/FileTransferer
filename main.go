package main

import(
	 Output "fmt"
	 Input "bufio"
	 "os"
	 "net"
	 "strconv"
	 "strings"
	 "io"
)

//FileExists gets the full path to the file
func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func sendFileToClient(connection net.Conn, path string) {
	Output.Println("A client has connected!")
	defer connection.Close()
	file, err := os.Open(path)
	if err != nil {
		Output.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		Output.Println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	Output.Println("Sending filename and filesize!")
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	Output.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			Output.Println("Error:", err)
			break
		}
		connection.Write(sendBuffer)
		
		if err == io.EOF{
			break
		}
	}
	Output.Println("File has been sent, closing connection!")
	return
}

//StartServer function is the function to start the server
func StartServer(path string, IP string){
	server, err := net.Listen("tcp4", ":8001")
	if err != nil {
		Output.Println("Error listetning: ", err)
		os.Exit(1)
	}
	defer server.Close()
	Output.Println("Server started! Waiting for connections...")
	for {
		connection, err := server.Accept()
		if err != nil {
			Output.Println("Error: ", err)
			os.Exit(1)
		}
		Output.Println("Client connected")
		go sendFileToClient(connection, path)
	}
}

//StartClient is the function to start the client
func StartClient(IP string){
	connection, err := net.Dial("tcp", IP)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	Output.Println("Connected to server, start receiving the file name and file size")
	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)
	
	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	
	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")
	
	newFile, err := os.Create(fileName)
	
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	var receivedBytes int64
	
	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
	Output.Println("Received file completely!")
}

const (
	//Send file to the server (client)
	Send = iota + 1
	//Get file from the client (server)
	Get
)

//BUFFERSIZE is the size of the buffer that will be sent to the client
const BUFFERSIZE = 1024

func main() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if(err != nil){
		Output.Println("Oops, we've got an error: " + err.Error() + "\n")
		os.Exit(1)
	}
	
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	
	Repeat:
	
	Output.Println("{1} - Send a file (Server)")
	Output.Println("{2} - Get a file (Client)")
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
		
		StartServer(path, localAddr.String());
		
	} else if(option == strconv.Itoa(Get)){
		Output.Println("\nIP of the server")
		serverIP, _ := reader.ReadString('\n')
		serverIP = strings.TrimRight(serverIP, "\r\n")
		
		StartClient(serverIP);
		
	} else{
		goto Repeat
		
	}
	 
	
	
	Output.Println();
	Output.Scanf("h")
	_, _ = reader.ReadString('\n')
}
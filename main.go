package main

import(
	 Output "fmt"
	 Input "bufio"
	 "os"
)

func main() {
	Output.Println("Hey all, I'm the Go app")
	Output.Println("And what's your name?")
	reader := Input.NewReader(os.Stdin)
	
	text, _ := reader.ReadString('\n')
	Output.Println("Oh, it's nice to meet you, ", text)
}
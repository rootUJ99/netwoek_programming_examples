package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
func clientLogger(message string) {
	fmt.Printf("[client] -> %s", message)
}
func client() {
	scanner:=bufio.NewScanner(os.Stdin)
	clientLogger("tell your gospel over here: ")
	for scanner.Scan() {
		line:= scanner.Text()
		if line == "\n" {
			break
		}
		msg:=[]byte(line)
		conn, err := net.Dial("tcp", "localhost:9010")	
		defer conn.Close()
		if err != nil {
			fmt.Println("can not connect to the server!!", err)
		}
		n, err := conn.Write(msg)
		if err != nil {
			fmt.Println("error while sending data", err)
			return
		}
		clientLogger(fmt.Sprintf("sending this %d byes \n", n))
	}
	if err:= scanner.Err(); err != nil{
		fmt.Println("this is an error", err)
	} 

}

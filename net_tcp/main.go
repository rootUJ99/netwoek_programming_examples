package main 

import (
	"fmt"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()
	buff:= make([]byte, 1024)
	n, err:=conn.Read(buff)
	if err!=nil {
		fmt.Println("unable to read the data", err)
		return
	}
	serverlog(string(buff[:n]))
}
func server(){
	listner, err:=net.Listen("tcp", ":9010")
	
	if (err != nil){
		panic("network not created")
	}

	defer listner.Close()

	for {
		conn, err:=listner.Accept()
		if(err!=nil){
			fmt.Println("Error accepting connection", err)
			continue
		}
		go handler(conn)
	}
}
func serverlog(message string) {
	fmt.Printf("[server] -> %s \n", message)
}

func main() {
	go server()
	go client()
	fmt.Println("sending data to client and server")
	select{}
}

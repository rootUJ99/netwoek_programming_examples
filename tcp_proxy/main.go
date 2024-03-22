package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func handleReceive(conn net.Conn, serverName string){
	defer conn.Close()
	byteArr:= make([]byte, 1024)
	j, err:=conn.Read(byteArr)
	if err!=nil {
		fmt.Println("reading error")
	}
	fmt.Printf("reading %s-> %s\n", serverName,  string(byteArr[:j]))
	
}

func server() {
	fmt.Println("hello from the universe")
	listner, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("listner error", err)
	}
	defer listner.Close()
	for {
		conn, err:=listner.Accept()
		if err!=nil {
			fmt.Println("connection error")
		}
		go handleReceive(conn, "mainserver")
	}
}

func client(){
	scanner:=bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		msg:= scanner.Text()
		if msg == "\n" {
			break
		}
		sender, err:=net.Dial("tcp", "localhost:9000")
		defer sender.Close()
		if err != nil{
			fmt.Println("sender error")
		}
		_, err=sender.Write([]byte(msg))
		if err!=nil {
			fmt.Println("writing error")
		}
		fmt.Printf("sending -> %s \n", msg)
		

		sender2, err:=net.Dial("tcp", "localhost:8080")
		defer sender2.Close()
		if err != nil {
			fmt.Println("sender error")
		}
		go io.Copy(sender2, sender)
		io.Copy(sender, sender2)

	}

	if scanner.Err() != nil {
		fmt.Println("scanning error")
	}
	
}

func proxyServer(){
	fmt.Println("hello from the proxy universe")
	listner, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("listner error", err)
	}
	defer listner.Close()
	for {
		conn, err:=listner.Accept()
		if err!=nil {
			fmt.Println("connection error")
		}
		go handleReceive(conn, "proxyserver")
	}
}

func main(){
	go server()
	go proxyServer()
	go client()
	select {}
}

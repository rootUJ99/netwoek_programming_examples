package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)



func client(udpAdd *net.UDPAddr) {
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("[clinet] -> tell your gospel over here: ")
	for scanner.Scan() {
		line:= scanner.Text()
		if line == "\n" {
			break
		}
		msg:=[]byte(line)
		conn, err:= net.DialUDP("udp", nil,udpAdd)
		defer conn.Close()
		if err != nil {
			fmt.Println("can not connect to the server!!", err)
		}
		n, err := conn.Write(msg)
		if err != nil {
			fmt.Println("error while sending data", err)
			return
		}
		fmt.Println(fmt.Sprintf("[client] -> sending this %d byes \n", n))
	}
	if err:= scanner.Err(); err != nil{
		fmt.Println("this is an error", err)
	} 

}

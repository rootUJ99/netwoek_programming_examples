package main

import (
	"fmt"
	"net"
)


func handler(conn *net.UDPConn){
	buff:= make([]byte, 1024)
	for {
		n, addr, err :=conn.ReadFromUDP(buff)
		if(err !=nil){
			fmt.Println("not able to read any data", err)
			return
		}
		fmt.Printf("[server] -> received %d data on %s\n", n, addr)	
		fmt.Println("[server] ->", string(buff[:n]))	
	}
}

func server(udpAdd *net.UDPAddr){

	conn, err:=net.ListenUDP("udp", udpAdd)

	if (err!=nil){
		fmt.Println("connection failed!!")
	}

	go func (){
		handler(conn)
		conn.Close()
	}()
}

func main(){
	udpAdd:=&net.UDPAddr{
		IP  : net.ParseIP("0.0.0.0"), 
		Port: 9011, 
	}

	go server(udpAdd)
	go client(udpAdd)
	select{}
} 

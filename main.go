package main

import (
	"fmt"
	"net"
)

func main(){

	l,err := net.Listen("tcp",":6379")
	
	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println("Listening on port :6379")

	conn,err := l.Accept()

	if err != nil {
		fmt.Println(err)
		return 
	}

	defer conn.Close()

	


}
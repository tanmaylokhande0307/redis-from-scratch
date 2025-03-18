package main

import (
	"fmt"
	"net"
)

func main() {

	//create new server
	l, err := net.Listen("tcp", ":6370")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port :6370")

	// Listen for connections
	conn, err := l.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	//create infinite loop to receive commands from the client and execute them

	for {
		
		resp := NewResp(conn)

		value, err := resp.Read()
		
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		conn.Write([]byte("+OK\r\n")) //according to RESP + = simple string

	}

}

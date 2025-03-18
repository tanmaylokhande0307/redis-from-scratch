package main

import (
	"fmt"
	"io"
	"net"
	"os"
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
		buffer := make([]byte, 1024)

		_, err = conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n")) //according to RESP + = simple string

	}

}

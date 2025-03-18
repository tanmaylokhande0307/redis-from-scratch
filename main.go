package main

import (
	"fmt"
	"net"
	"strings"
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

		writer := NewWriter(conn)

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)

	}

}

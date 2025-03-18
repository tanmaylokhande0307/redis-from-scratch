package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"
    "strconv"
)

func main() {
    input := "$6\r\nTanmay\r\n"
    reader := bufio.NewReader(strings.NewReader(input))
	
	fmt.Println(strings.NewReader(input))

    b, _ := reader.ReadByte()


    if b != '$' {
      fmt.Println("Invalid type, expecting bulk strings only")
      os.Exit(1)
    }

    size, _ := reader.ReadByte()

    strSize, _ := strconv.ParseInt(string(size), 10, 64)	
	
    // consume /r/n
    reader.ReadByte()
    reader.ReadByte()
	
    name := make([]byte, strSize)
	
	fmt.Println(strSize)
    reader.Read(name)

    fmt.Println(string(name))
}


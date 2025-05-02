package main

import (
	"fmt"
	"os"

	netcomm "github.com/rossb34/go-netcomm"
)

func main() {
	fmt.Println("Running sample tcp client")

	address := "127.0.0.1:5555"
	tcpClient, err := netcomm.Connect(address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tcpClient.Close()

	n, err := tcpClient.Write([]byte("Hello"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %d bytes\n", n)
}

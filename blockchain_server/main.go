package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP port for blockchain server.")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))

	app.Run()
	fmt.Println("Blockchain Server")
}

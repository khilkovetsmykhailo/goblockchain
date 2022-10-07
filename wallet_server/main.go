package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Wallet Server:")
}

func main() {
	port := flag.Uint("port", 8080, "TCP port for wallet server.")
	gateway := flag.String("g", "0.0.0.0", "TCP port for wallet server.")
	flag.Parse()
	ws := NewWalletServer(uint16(*port), *gateway)
	ws.Run()
}

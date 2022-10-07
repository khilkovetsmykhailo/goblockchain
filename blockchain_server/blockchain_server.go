package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func Helloworld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world")
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", Helloworld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa((int(bcs.Port()))), nil))
}

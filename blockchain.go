package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	var b Block
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash

	return &b
}

func (b *Block) Print() {
	fmt.Printf("timestamp	%d\n", b.timestamp)
	fmt.Printf("nonce		%d\n", b.nonce)
	fmt.Printf("previousHash	%s\n", b.previousHash)
	fmt.Printf("transactions	%s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init Hash")

	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}

	fmt.Println(strings.Repeat("=", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	bc := NewBlockchain()

	bc.CreateBlock(2, "second hash")

	bc.Print()
}

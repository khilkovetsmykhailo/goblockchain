package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	var b Block
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash

	return &b
}

func (b *Block) Print() {
	fmt.Printf("timestamp	%d\n", b.timestamp)
	fmt.Printf("nonce		%d\n", b.nonce)
	fmt.Printf("previousHash	%x\n", b.previousHash)
	fmt.Printf("transactions	%s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())

	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}

	fmt.Println(strings.Repeat("=", 25))
}

func (b *Block) Hash() [32]byte {
	m, _ := b.MarshalJson()

	return sha256.Sum256(m)
}

func (b *Block) MarshalJson() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) AddBlock() *Block {
	lb := bc.LastBlock()
	b := bc.CreateBlock(0, lb.Hash())

	return b
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock()
	bc.Print()
	bc.AddBlock()
	bc.Print()
}

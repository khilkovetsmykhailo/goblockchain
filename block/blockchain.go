package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"goblockchain/utils"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINER_SENDER      = "The Blockchain"
	MINING_REWARD     = 10
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	var b Block
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions

	return &b
}

func (b *Block) Print() {
	fmt.Printf("timestamp	%d\n", b.timestamp)
	fmt.Printf("nonce		%d\n", b.nonce)
	fmt.Printf("previousHash	%x\n", b.previousHash)

	for _, tr := range b.transactions {
		tr.Print()
	}
}

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}

	return b
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())

	return bc
}

func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"chains"`
	}{
		Blocks: bc.chain,
	})
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}

	fmt.Println(strings.Repeat("=", 25))
}

func (b *Block) Hash() [32]byte {
	m, _ := b.MarshalJSON()

	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash string         `json:"previousHash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: fmt.Sprintf("%x", b.previousHash),
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

func (bc *Blockchain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *utils.Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)

	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)

	for _, tr := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(tr.senderBlockchainAddress, tr.recipientBlockchainAddress, tr.value))
	}

	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)

	guessBlock := Block{nonce, previousHash, 0, transactions}
	guessHash := fmt.Sprintf("%x", guessBlock.Hash())

	return guessHash[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	nonce := 0
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()

	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce = nonce + 1
	}

	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINER_SENDER, bc.blockchainAddress, MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)

	return true
}

func (bc *Blockchain) BalanceOf(blockchainAddress string) float32 {
	var balance float32 = 0.0

	for _, b := range bc.chain {
		for _, tr := range b.transactions {
			if tr.recipientBlockchainAddress == blockchainAddress {
				balance = balance + tr.value
			}

			if tr.senderBlockchainAddress == blockchainAddress {
				balance = balance - tr.value
			}
		}
	}

	return balance
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32, publicKey *ecdsa.PublicKey, s *utils.Signature) bool {
	t := NewTransaction(sender, recipient, value)

	if sender == MINER_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	if bc.BalanceOf(sender) < value {
		log.Println("ERROR: No balance")
		return false
	}

	if bc.VerifyTransactionSignature(publicKey, s, t) {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("ERROR: Verify Transaction")
	}

	return false
}

func (tr *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" Sender: %s\n", tr.senderBlockchainAddress)
	fmt.Printf(" Recipient: %s\n", tr.recipientBlockchainAddress)
	fmt.Printf(" Value: %.1f\n", tr.value)
}

func (tr *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    tr.senderBlockchainAddress,
		Recipient: tr.recipientBlockchainAddress,
		Value:     tr.value,
	})
}

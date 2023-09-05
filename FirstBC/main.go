package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previoushash string
	timestamp    int64
	transations  []string
}

func NewBlock(nonce int, previoushash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previoushash = previoushash
	return b

}

func (b *Block) Print() {
	fmt.Printf("timestamp        %d\n", b.timestamp)
	fmt.Printf("previous hash    %s\n", b.previoushash)
	fmt.Printf("nonce            %d\n", b.nonce)
	fmt.Printf("transactions     %s\n", b.transations)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previoushash string) *Block {
	b := NewBlock(nonce, previoushash)
	bc.chain = append(bc.chain, b)
	return b
}

func init() {
	log.SetPrefix("BChain")
}

func main() {
	blockchain := NewBlockchain()
	fmt.Println(blockchain)
}

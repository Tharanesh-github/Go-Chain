package core

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block represents a single "page" in our cryptographic ledger (notebook)
type Block struct {
	Timestamp     int64  // When the block was created
	Data          []byte // The actual transactions/data in the block
	PrevBlockHash []byte // The SHA-256 fingerprint of the previous block
	Hash          []byte // The SHA-256 fingerprint of THIS block
}

// CalculateHash generates the SHA-256 fingerprint for the block.
// This is the core cryptographic function that secures the chain.
func (b *Block) CalculateHash() {
	// 1. Combine all the block's data (time, data, and previous hash) into one giant string
	record := fmt.Sprintf("%d%s%s", b.Timestamp, b.Data, b.PrevBlockHash)
	
	// 2. Run that string through the SHA-256 algorithm
	hash := sha256.Sum256([]byte(record))
	
	// 3. Store the resulting 256-bit hash securely on the block
	b.Hash = hash[:]
}

// NewBlock creates a newly minted block and securely hashes it
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
	}
	
	// Generate the cryptographic seal for this new block immediately
	block.CalculateHash()
	
	return block
}

// NewGenesisBlock creates the very first block in the chain.
// Because it's the first, it has no "previous" block to link to.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
package core

// Blockchain represents our immutable ledger. 
// Under the hood, it's essentially a securely linked list of Blocks.
type Blockchain struct {
	Blocks []*Block
}

// AddBlock securely mints a new block and adds it to the chain.
func (bc *Blockchain) AddBlock(data string) {
	// 1. Get the very last block currently in the chain
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	
	// 2. Create a new block, linking it to the previous block's hash
	newBlock := NewBlock(data, prevBlock.Hash)
	
	// 3. Append the new block to the chain
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain initializes a brand new ledger.
// A blockchain must ALWAYS start with a Genesis Block!
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}

package blockchain

type BlockChain struct {
	Block []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Block[len(chain.Block)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Block = append(chain.Block, new)
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

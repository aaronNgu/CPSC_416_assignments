package blockchain

import "encoding/json"

type Blockchain struct {
	Chain               []Block       `json:"chain"`
	Difficulty          int           `json:"diff"`
	PendingTransactions []Transaction `json:"pendingTransactions"`
}

func NewBlockchain(difficulty int) *Blockchain {
	genesisBlock := CreateGenesisBlock()
	chain := make([]Block, 0)
	chain = append(chain, *genesisBlock)
	blockchain := &Blockchain{Chain: chain, Difficulty: difficulty}
	return blockchain
}

func CreateGenesisBlock() *Block {
	return NewBlock([]Transaction{{Post: "Genesis Block"}})
}

func (blockchain *Blockchain) GetLatestBlock() *Block {
	lenghtOfChain := len(blockchain.Chain)
	return &blockchain.Chain[lenghtOfChain-1]
}

func (blockchain *Blockchain) MinePendingTransactions(miningRewardAddress string) {
	newBlock := NewBlock(blockchain.PendingTransactions)
	newBlock.MineBlock(blockchain.Difficulty)
	blockchain.Chain = append(blockchain.Chain, *newBlock)
	blockchain.PendingTransactions = []Transaction{{Post: "", CharCredit: 100, Addr: Addr{miningRewardAddress}}}
}

func (blockchain *Blockchain) CreateTransaction(transaction Transaction) {
	blockchain.PendingTransactions = append(blockchain.PendingTransactions, transaction)
}

func (blockchain *Blockchain) GetBalanceOfAddress(Address string) int {
	balance := 0
	for _, block := range blockchain.Chain {
		for _, transaction := range block.Data {
			if transaction.Addr.Key == Address {
				balance += transaction.CharCredit
			}
		}
	}
	return balance
}

func (blockchain *Blockchain) AddBlock(newBlock *Block) {
	newBlock.PrevHash = blockchain.GetLatestBlock().Hash
	newBlock.MineBlock(blockchain.Difficulty)
	blockchain.Chain = append(blockchain.Chain, *newBlock)
}

func (blockchain *Blockchain) IsChainValid() bool {
	for i := 1; i < len(blockchain.Chain); i++ {
		currBlock := blockchain.Chain[i]
		prevBlock := blockchain.Chain[i-1]
		if currBlock.Hash != currBlock.CalculateHash() {
			return false
		}

		if currBlock.PrevHash != prevBlock.CalculateHash() {
			return false
		}
	}
	return true
}

func (blockchain Blockchain) String() string {
	data, _ := json.Marshal(blockchain)
	return string(data)
}

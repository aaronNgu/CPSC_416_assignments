package blockchain

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type Block struct {
	Data     []Transaction `json:"transaction"`
	PrevHash [16]byte      `json:"prevHash"`
	Hash     [16]byte      `json:"hash"`
	Nonce    int           `json:"nonce"`
}

func NewBlock(data []Transaction) *Block {
	block := Block{Data: data, Nonce: 0}
	block.Hash = block.CalculateHash()
	return &block
}

func (block *Block) CalculateHash() [16]byte {
	data := make([]byte, 0)
	for _, d := range block.Data {
		data = append(data, []byte(d.Post)...)
	}
	content := append(data, block.PrevHash[:]...)
	content = append(content, byte(block.Nonce))
	hash := md5.Sum(content)
	//fmt.Println("\nhash is")
	//fmt.Printf("%x", hash)
	//fmt.Println("\n in string")
	//fmt.Println(hash)
	return hash
}

func (block *Block) MineBlock(difficulty int) {
	//fmt.Println("difficulty is  ", difficulty)
	hasNZero := func() bool {
		for i := 0; i < difficulty; i++ {
			if block.Hash[i] != 0 {
				return false
			}
		}
		return true
	}

	//fmt.Println("or here")
	for !hasNZero() {
		//fmt.Printf("\ngot here %d", block.Nonce)
		block.Nonce++
		block.Hash = block.CalculateHash()
	}
}

func (block Block) String() string {
	fmt.Println("transaction is")
	fmt.Println(block.Data[0])
	data, _ := json.Marshal(block)
	return string(data)
}

package blockchain

import (
	"fmt"
)

type Transaction struct {
	CharCredit int    `json:"charCredit"` // if CharCredit negative it has a posts else its a new mined blocked
	Addr       Addr   `json:"node"`
	Post       string `json:"transaction"`
}

func (t Transaction) String() string {
	fmt.Println("got here")
	//tr := Transaction{CharCredit: t.CharCredit, Addr: "somthing", Post: t.Post}
	//data, _ := json.Marshal(tr)
	//return string(data)
	return "hey"
}

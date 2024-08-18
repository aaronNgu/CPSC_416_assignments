package blockchain

import (
	"encoding/json"
	"fmt"
)

type Addr struct {
	Key string
}

func (a Addr) MarshalJSON() ([]byte, error) {
	return json.Marshal("Thisg is an Addr")
}

func (a Addr) String() string {
	fmt.Println("got here")
	//tr := Transaction{CharCredit: t.CharCredit, Addr: "somthing", Post: t.Post}
	//data, _ := json.Marshal(tr)
	//return string(data)
	return "hey its addr"
}

package basic

import (
	"encoding/json"
	"fmt"
)

type Person1 struct {
	Age       int    `json:"age,string"`
	Name      string `json:"name"`
	Niubility bool   `json:"niubility"`
}

func ToStruct() {
	// 假数据
	b := []byte(`{"age":"18","name":"5lmh.com","marry":false}`)
	var p Person1
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}

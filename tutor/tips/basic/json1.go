package basic

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"-"`
	Hobby string `json:"hobby"`
}

func JsonMarshal() {
	p := Person{"5lmh.com", "女"}
	// 编码json
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))

	// 格式化输出
	b, err = json.MarshalIndent(p, "", "     ")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}
func JsonMar() {
	student := make(map[string]interface{})
	student["name"] = "5lmh.com"
	student["age"] = 18
	student["sex"] = "man"
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}

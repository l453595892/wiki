package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	a := T1{
		Ta: A(1),
		Tb: 2,
	}

	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(a)
	fmt.Println(b)

	var c T2
	err := json.NewDecoder(b).Decode(&c)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(c.Ta)
}

type A int

type B int

type T1 struct {
	Ta A   `json:"ta"`
	Tb int `json:"tb"`
}

type T2 struct {
	Ta B   `json:"ta"`
	Tb int `json:"tb"`
}




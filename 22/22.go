package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	a.SetString("121389274897128471289478127489127489", 10)
	b := new(big.Int)
	b.SetString("58928381273827138972819371283718927", 10)
	res := new(big.Int)

	fmt.Println(res.Add(a, b))
	fmt.Println(res.Sub(a, b))
	fmt.Println(res.Mul(a, b))
	fmt.Println(res.Div(a, b))
}

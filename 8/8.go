package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64 = 967296
	num = Set(5, num)
	fmt.Println(num)
}

func Set(i int, num int64) int64 {
	str := strconv.FormatInt(num, 2)
	var helpNum int64 = (1 << i)
	if num > helpNum && str[i] == 1 {
		return num &^ helpNum
	} else {
		return num ^ helpNum
	}
}

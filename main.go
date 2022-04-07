package main

import (
	"Assignment1/function"
	"os"
	"strconv"
)

func main() {
	getArg := os.Args[1]
	listNum, _ := strconv.ParseInt(getArg, 0, 64)
	function.PrintData(int(listNum))
}

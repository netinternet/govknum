package main

import (
	"fmt"
	"github.com/netinternet/govknum"
)

func main() {
	if valid := govknum.Vknum("3128402031"); valid {
		fmt.Println("Tax Number Is Valid")
	} else {
		fmt.Println("Tax Number Is Not Valid")
	}
}

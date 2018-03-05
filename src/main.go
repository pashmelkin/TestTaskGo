package main

import (
	"fmt"
	"lib"
)

func main() {

  checkout.NewCustomer("Uniliver")
	checkout.Add(checkout.Classic)
	checkout.Add(checkout.Classic)
	fmt.Println(checkout.Total())
}

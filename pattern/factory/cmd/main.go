package main

import (
	"fmt"
	"l2/pattern/factory/pkg"
)

func main() {
	ak47, _ := pkg.GetGun("ak47")
	musket, _ := pkg.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g pkg.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}

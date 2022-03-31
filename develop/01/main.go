package main

import (
	"01/ntp"
	"fmt"
	"os"
)

func main() {
	t, err := ntp.GetTime("0.ru.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(t)
}

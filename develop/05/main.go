package main

import (
	"fmt"
	"log"
)

func main() {
	cfg := &GrepConfig{
		after:       0,
		before:      0,
		contextRows: 0,
		count:       false,
		ignoreCase:  false,
		invert:      false,
		fixed:       false,
		strNum:      true,
		regExp:      "heh",
		filename:    "develop/dev05/test.txt",
	}
	res, err := grep(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

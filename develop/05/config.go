package main

import (
	"flag"
	"log"
)

type GrepConfig struct {
	after       int
	before      int
	contextRows int
	count       bool
	ignoreCase  bool
	invert      bool
	fixed       bool
	strNum      bool
	regExp      string
	filename    string
}

func NewConfig() *GrepConfig {
	cfg := GrepConfig{}
	flag.IntVar(&cfg.after, "A", 0, "Print +N rows after match")
	flag.IntVar(&cfg.before, "B", 0, "Print +N rows before match")
	flag.IntVar(&cfg.contextRows, "C", 0, "Print +N rows after and before match")
	flag.BoolVar(&cfg.count, "c", false, "Print count of match rows")
	flag.BoolVar(&cfg.ignoreCase, "i", false, "Ignore case")
	flag.BoolVar(&cfg.invert, "v", false, "Instead of a match, exclude")
	flag.BoolVar(&cfg.fixed, "F", false, "Exact match with a string, not a pattern")
	flag.BoolVar(&cfg.strNum, "n", false, "Print line number of match rows")

	flag.Parse()

	args := flag.Args()
	if cfg.contextRows > 0 {
		cfg.after, cfg.before = cfg.contextRows, cfg.contextRows
	}
	if cfg.count {
		cfg.after, cfg.before = 0, 0
	}
	if len(args) == 2 {
		cfg.regExp = args[0]
		cfg.filename = args[1]
	} else {
		log.Fatalf("The argument (path to the file name) must be one")
	}

	return &cfg
}

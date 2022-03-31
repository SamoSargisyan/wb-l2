package main

import (
	"flag"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func commands() *cli.App {
	var app = cli.NewApp()

	app.Name = "Simple CLI"
	app.Usage = "An example CLI"
	app.Author = "Samson Sargisyan"
	app.Version = "1.0.0"

	flag.Parse()

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input := string(bytes)

	app.UseShortOptionHandling = true
	app.Commands = []cli.Command{
		{
			Name:    "parse file",
			Aliases: []string{"parse"},
			Usage:   "add a new template",
			Flags: []cli.Flag{
				cli.IntFlag{Name: "column, c"},
				cli.BoolFlag{Name: "unique, u"},
				cli.BoolFlag{Name: "reverse, r"},
				cli.BoolFlag{Name: "numeric, n"},
			},
			Action: func(c *cli.Context) error {
				listOfStrings := strings.Split(input, "\n")
				var sortedLines []string

				if c.Bool("unique") {
					set := make(map[string]bool)
					for _, value := range listOfStrings {
						set[value] = true
					}

					for value := range set {
						sortedLines = append(sortedLines, value)
					}

				}

				sortedLines = listOfStrings
				sort.Strings(sortedLines)

				if c.Bool("reverse") {
					sort.Sort(sort.Reverse(sort.StringSlice(sortedLines)))
				}

				if c.Bool("numeric") {
					var keys []float64
					numKeyForString := make(map[float64][]string)

					for _, line := range sortedLines {
						currentLine := strings.Split(line, " ")
						key, err := strconv.ParseFloat(currentLine[0], 32)
						if err != nil {
							return err
						}
						keys = append(keys, key)

						if _, ok := numKeyForString[key]; ok {
							numKeyForString[key] = append(numKeyForString[key], line)
						} else {
							numKeyForString[key] = append(numKeyForString[key], line)
						}
					}
					keys = removeDuplicateInt(keys)

				}

				if c.Int("column") == 0 {
					sort.Sort(sort.Reverse(sort.StringSlice(sortedLines)))
				}

				if c.Int("column") > 0 {
					keys := make([]string, 0, len(sortedLines))
					m := make(map[string][]string)
					for _, line := range sortedLines {
						currentLine := strings.Split(line, " ")
						key := currentLine[c.Int("column")-1]
						keys = append(keys, key)

						if _, ok := m[key]; ok {
							m[key] = append(m[key], line)
						} else {
							m[key] = append(m[key], line)
						}
					}

					keys = removeDuplicateStr(keys)
					sort.Strings(keys)
				}

				return nil
			},
		},
	}

	return app
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func removeDuplicateInt(intSlice []float64) []float64 {
	allKeys := make(map[float64]bool)
	var list []float64
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	app := commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

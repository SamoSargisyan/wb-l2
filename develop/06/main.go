package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CutConfig struct {
	separated bool
	fields    string
	delimiter string
}

func readDataFromStdin() ([]string, error) {
	var listOfStrings []string

	fmt.Print("Enter text:\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			listOfStrings = append(listOfStrings, text)
		} else {
			break
		}
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return listOfStrings, nil
}

func parseFieldsToInt(fields string) ([]int, error) {
	switch {
	case strings.Contains(fields, "-"):
		numbers := strings.Split(fields, "-")
		if len(numbers) > 2 {
			return nil, errors.New("invalid fields")
		}
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, err
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, err
		}
		if num2 < num1 {
			return nil, errors.New("invalid fields")
		}
		res := make([]int, num2-num1+1)
		for i, j := 0, num1; j <= num2; i, j = i+1, j+1 {
			res[i] = j
		}
		return res, nil
	case strings.Contains(fields, ","):
		numbers := strings.Split(fields, ",")
		res := make([]int, len(numbers))
		for i := range numbers {
			num, err := strconv.Atoi(numbers[i])
			if err != nil {
				return nil, err
			}
			res[i] = num
		}
		sort.Ints(res)
		return res, nil
	default:
		res := make([]int, 1)
		number, err := strconv.Atoi(fields)
		if err != nil {
			return nil, err
		}
		res[0] = number
		return res, nil
	}
}

func cut(input []string, cfg *CutConfig) ([]string, error) {
	if len(input) == 0 {
		return nil, errors.New("")
	}
	result := make([]string, 0)
	fields, err := parseFieldsToInt(cfg.fields)
	if err != nil {
		return nil, err
	}

	for _, s := range input {
		if !strings.Contains(s, cfg.delimiter) && cfg.separated {
			continue
		}
		if !strings.Contains(s, cfg.delimiter) && !cfg.separated {
			result = append(result, s)
			continue
		}
		columns := strings.Split(s, cfg.delimiter)
		if len(fields) > 1 {
			var a strings.Builder
			for j, num := range fields {
				if len(columns) > num-1 {
					a.WriteString(columns[num-1])
					if len(fields)-1 != j {
						a.WriteString(cfg.delimiter)
					}
				}

			}

			result = append(result, a.String())
		} else {
			result = append(result, columns[fields[0]-1])
		}

	}
	return result, nil
}

func main() {

	cfg := CutConfig{
		separated: false,
		fields:    "10-13",
		delimiter: ";",
	}
	str, err := readDataFromStdin()
	if err != nil {
		log.Fatalln(err)
	}
	res, err := cut(str, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	for _, val := range res {
		fmt.Print(val)
	}
}

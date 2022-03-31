package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Text struct {
	text string
	num  int
}

func fileRead(filename string) ([]Text, error) {
	var rows []Text
	file, err := os.Open(filename)
	if err != nil {
		return rows, err
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	count := 1
	for sc.Scan() {
		rows = append(rows, Text{
			text: sc.Text(),
			num:  count,
		})
		count++
	}
	return rows, nil
}

func grep(cfg *GrepConfig) (string, error) {
	var prefix, postfix string
	if cfg.ignoreCase {
		prefix = "(?i)"
	}

	if cfg.fixed {
		prefix += "^"
		postfix = "$"
	}
	re, err := regexp.Compile(prefix + cfg.regExp + postfix)
	if err != nil {
		return "", err
	}
	file, err := fileRead(cfg.filename)
	if err != nil {
		return "", err
	}
	indexStr := make([]int, 0)
	if cfg.invert {
		for i, str := range file {
			if !re.MatchString(str.text) {
				indexStr = append(indexStr, i)
			}
		}
	} else {
		for i, str := range file {
			if re.MatchString(str.text) {
				indexStr = append(indexStr, i)
			}
		}
	}

	if len(indexStr) == 0 {
		return "", nil
	}

	result := make([]Text, 0, len(indexStr))
	switch {
	case cfg.after > 0 || cfg.before > 0:
		after, before := cfg.after, cfg.before
		start, end := indexStr[0]-before, indexStr[0]+after
		if indexStr[0]-before < 0 {
			start = 0
		}
		if end >= len(file) {
			end = len(file) - 1
		}
		result = append(result, file[start:end+1]...)

		if len(indexStr) > 1 {
			for _, index := range indexStr[1:] {

				if before != 0 {

					if index > end && end >= index-before {
						result = append(result, file[end+1:index+1]...)
						end = index
					}
					if index > end && end < index-before {
						result = append(result, file[index-before:index+1]...)
						end = index
					}
				}

				if index > end {
					result = append(result, file[index])
				}

				if after != 0 {
					lastA := end
					if len(file) <= index+after {
						end = len(file) - 1
					} else {
						end = index + after
					}
					if index <= lastA && lastA <= index+after {
						result = append(result, file[lastA+1:end+1]...)
						continue
					}
					result = append(result, file[index+1:end+1]...)
				}

			}

		}
	case cfg.count:
		return strconv.Itoa(len(indexStr)), nil
	default:
		for _, index := range indexStr {
			result = append(result, file[index])
		}
	}

	return join(result, cfg.strNum), nil
}

func join(elems []Text, strNum bool) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		if strNum {
			return strconv.Itoa(elems[0].num) + ": " + elems[0].text
		}
		return elems[0].text
	}
	sep := "\n"
	var b strings.Builder
	if strNum {
		b.WriteString(strconv.Itoa(elems[0].num) + ": " + elems[0].text)
		for _, s := range elems[1:] {
			b.WriteString(sep)
			b.WriteString(strconv.Itoa(s.num) + ": " + s.text)
		}
		return b.String()
	}

	b.WriteString(elems[0].text)
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s.text)
	}
	return b.String()
}

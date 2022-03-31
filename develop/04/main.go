package main

import (
	"fmt"
	"sort"
	"strings"
)

func deleteRepeated(in []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool)

	for _, v := range in {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

func AnagramDict(in []string) map[string][]string {
	for i := range in {
		in[i] = strings.ToLower(in[i])
	}
	woRepeated := deleteRepeated(in)
	tempM := make(map[string][]string, 0) //промежуточная мапа, ключ - отсортированное слов

	for _, v := range woRepeated {
		sorted := []rune(v)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
		sortedS := string(sorted) //отсортированное слово

		tempM[sortedS] = append(tempM[sortedS], v)
	}

	//результирующая мапа
	resultM := make(map[string][]string, 0)

	for _, v := range tempM {
		if len(v) > 1 { //если всего один элемент - в словрь не попадает
			resultM[v[0]] = v //нулевой элемент, это первый добавленный (первый просмотренный)
			sort.Strings(v)
		}
	}

	return resultM

}

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(input)
	fmt.Println(AnagramDict(input))
}

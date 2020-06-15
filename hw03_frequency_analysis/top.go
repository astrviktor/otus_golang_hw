package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

// Top10 - топ 10 слов в строке.
func Top10(text string) []string {
	// символы под замену
	var symbols = []string{".", "!", "?", ",", ";", " -", "- ",
		" - ", "—", " —", "— ", " — ", "\n", "\t", `"`}

	// переводим в нижний регистр
	text = strings.ToLower(text)

	// заменяем символы на пробелы
	for _, symbol := range symbols {
		text = strings.ReplaceAll(text, symbol, " ")
	}

	// разделяем на слова
	strs := strings.Split(text, " ")

	// перекладываем в map
	frequencyStrings := make(map[string]int)

	for _, s := range strs {
		if s != "" {
			frequencyStrings[s]++
		}
	}
	// если map пустой то нет слов
	if len(frequencyStrings) == 0 {
		return nil
	}

	// перекладываем из map в slice структур которые можно будет сортировать
	type StrCount struct {
		Str   string
		Count int
	}

	StrCountSlice := make([]StrCount, 0)

	for key, value := range frequencyStrings {
		StrCountSlice = append(StrCountSlice, StrCount{key, value})
	}

	// сортируем slice из слов и их количества
	sort.Slice(StrCountSlice, func(i, j int) bool {
		return StrCountSlice[i].Count > StrCountSlice[j].Count
	})

	// если количество элементов > 10 то отсекаем
	if len(StrCountSlice) > 10 {
		StrCountSlice = StrCountSlice[:10]
	}

	// нужно вернуть slice string
	res := make([]string, 0)

	for _, val := range StrCountSlice {
		res = append(res, val.Str)
	}

	return res
}

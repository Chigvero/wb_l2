package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var columnToSort = flag.Int("k", 0, "Указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
var intToSort = flag.Bool("n", false, "Cортировка строк linux по числовому значению")
var reverseSort = flag.Bool("r", false, "Сортировать в обратном порядке")
var noDuplicat = flag.Bool("u", false, "Не выводить повторяющиеся строки")

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := make([]string, 0)
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func MySort(data []byte, r, n, u bool, k int) (string, error) {
	rows := strings.Split(string(data), "\n")
	var result string

	if u { // if user asked for unique
		rows = removeDuplicateStr(rows)
	}

	if n {
		numbers := make([]int, 0)
		for _, row := range rows {
			if numRow, err := strconv.Atoi(row); err == nil {
				numbers = append(numbers, numRow)
			} else {
				return "", fmt.Errorf("not numerical data")
			}
		}
		sort.Ints(numbers)
		if r {
			sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		}
		for _, row := range numbers {
			result += fmt.Sprintln(row)
		}
		return result, nil
	}

	rowsOfSlices := make([][]string, 0) 
	for _, row := range rows {
		rowSlice := strings.Split(row, " ")
		rowsOfSlices = append(rowsOfSlices, rowSlice)
	}

	if k < 0 || k >= len(rowsOfSlices[0]) {
		return "", fmt.Errorf("incorrect column number: %d", k)
	}

	sort.Slice(rowsOfSlices, func(i, j int) bool { 
		for x := k; x < len(rowsOfSlices[i]); x++ {
			if rowsOfSlices[i][k] == rowsOfSlices[j][k] {
				continue
			}
			if r {
				return rowsOfSlices[i][k] > rowsOfSlices[j][k]
			} else {
				return rowsOfSlices[i][k] < rowsOfSlices[j][k]
			}
		}
		return true
	})

	for _, rowSlice := range rowsOfSlices { // Print out
		for i := 0; i < len(rowSlice); i++ {
			result += rowSlice[i]
			if i != len(rowSlice)-1 { // if last word of row
				result += " "
			} else {
				result += "\n"
			}
		}
	}

	return result, nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	src := args[0]
	r := *reverseSort
	n := *intToSort
	u := *noDuplicat
	k := *columnToSort
	data, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}
	fmt.Print(MySort(data, r, n, u, k))
}
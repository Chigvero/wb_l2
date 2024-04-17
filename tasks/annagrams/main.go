package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Start")
	myarray := []string{"пятак", "а", "б", "а", "листок", "пятка", "тяпка", "мор", "тяпка", "слиток", "ром", "столик"}

	mymap := make(map[string][]string)
	finder(mymap, &myarray)
	fmt.Println(myarray)
	fmt.Print(mymap)
}

func find(key *string, elemVal *string, adder *bool) {
	*elemVal = strings.ToLower(*elemVal)
	keymap := make(map[rune]int)
	valmap := make(map[rune]int)
	for _, v := range *key {
		keymap[v]++
	}
	for _, v := range *elemVal {
		valmap[v]++
	}
	for k, v := range keymap {
		if valmap[k] != v {
			*adder = false
		}
	}
}

func finder(mymap map[string][]string, myarray *[]string) {
	for _, elemVal := range *myarray {
		if len(elemVal) < 4 {
			continue
		}
		checkKey := false
		for key, _ := range mymap {
			if len(key) == len(elemVal) {
				adder := true
				checkFlag := true
				find(&key, &elemVal, &adder)
				checkValue(mymap[key], &elemVal, &checkFlag)
				if adder {
					if checkFlag {
						mymap[key] = append(mymap[key], elemVal)
					}
					checkKey = true
				}
			}
		}
		if !checkKey {
			mymap[elemVal] = append(mymap[elemVal], elemVal)
		}
	}
}

func removeElement(i int, myarray **[]string) {

	(**myarray)[i] = (**myarray)[0]
	**myarray = (**myarray)[1:]
}

func checkValue(value []string, elemVal *string, checkFlag *bool) {
	for _, v := range value {
		if v == *elemVal {
			*checkFlag = false
		}
	}
}

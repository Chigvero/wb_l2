package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "a4b3ce"
	s1 := "abcd"
	s2 := ""

	out, err := Unpack(s1)
	fmt.Println(out, err)
	out, err = Unpack(s)
	fmt.Println(out, err)
	fmt.Scan(&s2)
	out, err = Unpack(s2)
	fmt.Println(out, err)

}

func Unpack(s1 string) (string, error) {
	var err error
	if len(s1) > 0 {
		_, x := strconv.Atoi(string(s1[0]))
		if x == nil {
			err = fmt.Errorf("некорректная строка")
			return s1, err
		}
		s1slice := []byte{}
		var prev byte = 0
		s_ := 0
		for _, curr := range s1 {
			c, _ := strconv.Atoi(string(curr))
			_, x := strconv.Atoi(string(prev))
			if s_ == 1 && curr != '\\' {
				s1slice = append(s1slice, byte(curr))
				s_ = 0
			} else if curr == '\\' {
				s_++
			} else if c != 0 && prev != '\\' {
				if x == nil {
					err = fmt.Errorf("некорректная строка : %v", string(curr))
					return s1, err
				}
				for j := 0; j < c-1; j++ {
					s1slice = append(s1slice, prev)
				}
			} else if c != 0 && prev == '\\' {
				if x == nil {
					err = fmt.Errorf("некорректная строка : %v", string(curr))
					return s1, err
				}
				for j := 0; j < c; j++ {
					s1slice = append(s1slice, prev)
				}
				s_ = 0
			} else if c == 0 {
				s1slice = append(s1slice, byte(curr))
			}
			prev = byte(curr)
		}
		s1 = string(s1slice)
	}
	return s1, err
}

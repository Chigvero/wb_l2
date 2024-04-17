package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/pborman/getopt/v2"
)

type Flages struct {
	after   *bool
	before  *bool
	context *bool
	count   *bool
	iCase   *bool
	invert  *bool
	fixed   *bool
	lineN   *bool
}

func (f *Flages) Parse() {
	f.after = getopt.BoolLong("after", 'A')
	f.before = getopt.BoolLong("before", 'B')
	f.context = getopt.BoolLong("context", 'C')
	f.count = getopt.BoolLong("count", 'c')
	f.iCase = getopt.BoolLong("ignore-case", 'i')
	f.invert = getopt.BoolLong("invert", 'v')
	f.fixed = getopt.BoolLong("fixed", 'F')
	f.lineN = getopt.Bool('n', "lineNum")

	getopt.ParseV2()
	//fmt.Printf("A:%v\n B:%v\n C:%v\n c:%v\n i:%v\n v:%v\n F:%v\n n:%v\n ", *f.after, *f.before, *f.context, *f.count, *f.iCase, *f.invert, *f.fixed, *f.lineN)

}

func main() {
	fl := Flages{}
	fl.Parse()
	countMatch := 0
	numberString := 1
	//opening file
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//compilation (flag i e)
	comp := regCompilation(&fl)

	reader := bufio.NewReader(file)
	// matching (another flags after before context)
	if *fl.after {
		afterMatch(comp, reader)
	} else if *fl.before {
		beforeMatch(comp, reader)
	} else if *fl.context {
		contextMatch(comp, reader)
	} else {
		Matching(comp, reader, &fl, &countMatch, &numberString)
	}
}

func Matching(comp *regexp.Regexp, reader *bufio.Reader, fl *Flages, countMatch *int, numberString *int) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		if *fl.invert && !comp.MatchString(line) {
			*countMatch++
			if !*fl.count && !*fl.lineN {
				fmt.Print(line)
			} else if !*fl.count {
				fmt.Printf("%d:%s", *numberString, line)
			}
		} else if !*fl.invert && comp.MatchString(line) {
			*countMatch++
			if !*fl.count && !*fl.lineN {
				fmt.Print(line)
			} else if !*fl.count {
				fmt.Printf("%d:%s", *numberString, line)
			}
		}
		*numberString++
	}
	if *fl.count {
		fmt.Println(*countMatch)
	}
}

func regCompilation(fl *Flages) *regexp.Regexp {
	var (
		reg string
	)
	if *fl.iCase {
		reg = fmt.Sprint("(?i)", os.Args[len(os.Args)-2])
	} else {
		reg = os.Args[len(os.Args)-2]
	}
	comp, err := regexp.Compile(reg)
	if err != nil {
		panic(err)
	}
	return comp
}

func afterMatch(comp *regexp.Regexp, reader *bufio.Reader) {
	Printcommand := false
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		if comp.MatchString(line) {
			Printcommand = true
			break
		}
	}
	if Printcommand {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println(err)
					return
				}
			}
			fmt.Print(line)
		}
	}
}

func beforeMatch(comp *regexp.Regexp, reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		if !comp.MatchString(line) {
			fmt.Print(line)
		} else {
			break
		}
	}
}

func contextMatch(comp *regexp.Regexp, reader *bufio.Reader) {
	Printcommand := false
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		if !comp.MatchString(line) {
			fmt.Print(line)
		} else {
			Printcommand = true
			break
		}
	}
	if Printcommand {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println(err)
					return
				}
			}
			fmt.Print(line)
		}
	}
}

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

type flages struct {
	d    bool
	f    bool
	s    bool
	dArg string
	fArg []string
}

func (fl *flages) NewParse() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "-d" {
			fl.d = true
			fl.dArg = os.Args[i+1]
		}
		if os.Args[i] == "-f" {
			fl.f = true
			fl.fArg = strings.Split(os.Args[i+1], ",")
		}
		if os.Args[i] == "-s" {
			fl.s = true
		}
	}
	fmt.Print(*fl)
}

func main() {
	fmt.Println("Start")
	fl := flages{}
	fl.NewParse()
	reader := bufio.NewReader(os.Stdin)

	file, err := os.Create("myfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Создаем контекст для корректного завершения работы горутин
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	// Используем канал для сигнала завершения работы
	done := make(chan struct{})

	var mainwait sync.WaitGroup

	mainwait.Add(1)
	go func() {
		defer mainwait.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				line, err := reader.ReadBytes('\n')
				if err != nil {
					if err != io.EOF {
						fmt.Println("Error reading input:", err)
						return
					}
					return
				}
				if _, err := file.Write(line); err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}
		}
	}()

	mainwait.Add(1)
	go func() {
		defer mainwait.Done()
		<-sigs
		fmt.Println("Signal received, reading from file...")
		file1, err := os.Open("myfile.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file1.Close()
		scanner := bufio.NewScanner(file1)
		for scanner.Scan() {
			if fl.d == true {
				text := strings.Split(scanner.Text(), fl.dArg)
				if fl.f == true {
					if fl.s == true && len(text) == 1 {
						continue
					} else {
						for _, v := range fl.fArg {
							i, _ := strconv.Atoi(v)
							if i < len(text) {
								fmt.Printf("%s ", text[i-1])
							}
						}
						fmt.Println()
					}
				} else {
					fmt.Print(text)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning file:", err)
		}
		// После завершения работы выводим сигнал в канал done
		close(done)
	}()

	// Ждем сигнала завершения от основной горутины или от сигнала
	select {
	case <-done:
		fmt.Println("Exiting gracefully")
	case <-ctx.Done():
		fmt.Println("Context canceled, exiting")
	}

	// Отменяем контекст и ждем завершения всех горутин
	cancel()
	mainwait.Wait()
	fmt.Println("All goroutines exited")
}

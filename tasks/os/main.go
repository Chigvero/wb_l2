package main

import (
	"bufio"
	"fmt"
	"os"
	path "path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	//fmt.Print("START")
	//fmt.Print(os.Args)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("shell>")
		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		command = strings.Trim(command, "\n")
		Args := strings.Split(command, " ")
		//fmt.Print(Args)
		if len(Args) == 0 {
			continue
		}
		switch Args[0] {
		case "cd":
			if len(Args) < 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(Args[1])
			if err != nil {
				fmt.Println("Usage: cd <directory>")
				continue
			}
		case "pwd":
			fmt.Print("pwd:")
			pwd, err := path.Abs("..")
			if err != nil {
				panic(err)
			}
			fmt.Println(pwd)
		case "echo":
			back_c := false
			for i := 1; i < len(Args); i++ {
				for _, val := range Args[i] {
					if back_c {
						fmt.Print(string(val))
						back_c = false
					} else if val == '\\' {
						back_c = true
						continue
					} else {
						fmt.Print(string(val))
					}
				}
			}
		case "ps":
			// Получаем информацию о текущем процессе
			currentProcess, err := os.FindProcess(os.Getpid())
			if err != nil {
				fmt.Println("Error getting process information:", err)
				continue
			}

			fmt.Println(currentProcess.Pid)

		case "kill":
			if len(Args) < 2 {
				fmt.Println("Usage: kill <pid>")
				continue
			}
			pidStr := Args[1]
			pid, err := strconv.Atoi(pidStr)
			if err != nil {
				fmt.Println("Invalid PID:", pidStr)
				continue
			}
			process, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println("Error finding process:", err)
				continue
			}
			err = process.Signal(syscall.SIGKILL)
			if err != nil {
				fmt.Println("Error killing process:", err)
			}
		case "exit":
			fmt.Println("Exiting shell...")
			return
		}

	}
}

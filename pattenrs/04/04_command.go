package main

/*
	Суть паттерна - создать прослойку между бизнес логикой и интерфейсом, тем самым
	обеспечив позволяя инкапсулировать запрос на выполнение определенной операции в
	виде отдельного объекта. Этот объект-команда хранит в себе все необходимые данные
	для выполнения этой операции.

	+:
		1)	Убирает прямую зависимость между объектами, которые вызывают запросы, и
			объектами, выполняющими эти запросы.
		2)	Позволяют реализовать простую отмену и повтор операции.
		3)	Позволяют реализовать отложенный запуск операции.
		4)	Реализуют принцип открытости/закрытости.

	-:
		1)	Усложняют код программы из-за добавления большого количества допольни-
			тельных классов.
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import (
	"fmt"
)

// ICommands - интерфейс команды
type ICommands interface {
	Execute()
}

// Command - структура, представляющая команду
type Command struct {
	Name string
}

func (c *Command) Execute() {
	fmt.Println(c.Name)
}

type CommandQueue struct {
	commands []ICommands
}

func (cq *CommandQueue) AddCommand(c *Command) {
	cq.commands = append(cq.commands, c)
}

func (cq *CommandQueue) ExecuteAll() {
	for _, value := range cq.commands {
		value.Execute()
	}
}

func main() {
	command1 := &Command{
		Name: "ls -la",
	}
	command2 := &Command{
		Name: "pwd",
	}
	command3 := &Command{
		Name: "echo 'Hello World!'",
	}

	commands := &CommandQueue{}

	commands.AddCommand(command1)
	commands.AddCommand(command2)
	commands.AddCommand(command3)

	commands.ExecuteAll()
}

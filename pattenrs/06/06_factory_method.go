package main

import (
	"fmt"
)

/*
	Фабричный метод - предоставляет интерфейс, создающий объект, но при этом
	позволяющий подклассам самим определять конкретный класс объекта. То есть
	мы определяем общий интерфейс для создания объектов, но детали конкретной
	реализации скрыты в подклассах

	+:
		1)	Упрощает добавление новых продуктов в программу.
		2)	Избавляет класс от привязки к конкретным классам продуктов.
		3)	Выделяет код производства продуктов в одно место, упрощая поддержку
			кода.
		4)	Реализует принцип открытости / закрытости.
		5)	Позволяет легко расширять программу.

	-:
		1)	Может привести к созданию большого количества паралелльных иерархий
			классов, так как для каждого класса продукта надо создать свой
			подкласс создателя.

*/

// Transport - интерфейс транспортного объекта
type ITransport interface {
	Move()
	SetSpeed(speed int)
	SetWeight(weight int)
	GetSpeed() int
	GetWeight() int
}

type Transport struct {
	name   string
	speed  int
	weight int
}

func (t *Transport) Move() {
	fmt.Println(t.name, " moves at a speed ", t.speed, " and a weight of ", t.weight)
}

func (t *Transport) SetSpeed(speed int) {
	t.speed = speed
}

func (t *Transport) SetWeight(weight int) {
	t.weight = weight
}

func (t *Transport) GetSpeed() int {
	return t.speed
}

func (t *Transport) GetWeight() int {
	return t.weight
}

type Car struct {
	Transport
}

type Train struct {
	Transport
}

type Plane struct {
	Transport
}

type TransportFactory struct {
}

func (t *TransportFactory) CreateTransport(isName string, isSpeed int, isWeight int) ITransport {
	switch isName {
	case "car":
		return &Car{Transport{name: isName, speed: isSpeed, weight: isWeight}}
	case "train":
		return &Train{Transport{name: isName, speed: isSpeed, weight: isWeight}}
	case "plane":
		return &Plane{Transport{name: isName, speed: isSpeed, weight: isWeight}}
	default:
		return nil
	}
}

func main() {
	factory := &TransportFactory{}

	car := factory.CreateTransport("car", 100, 200)
	train := factory.CreateTransport("train", 180, 3000)
	plane := factory.CreateTransport("plane", 700, 900)

	car.Move()
	train.Move()
	plane.Move()
}

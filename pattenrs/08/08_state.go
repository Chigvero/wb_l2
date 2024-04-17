package main

import "fmt"

/*
	Состояние - поведенческий паттерн, который позволяет менять объектам свое
	поведение в зависимости от своего внутреннего состояния. То есть объект
	может находится в одном из нескольких состояний, и его поведение от этого
	меняется.

	+:
		1)	Избавляет от множества больших условных операторов машины состояний.		2)	Изолирует код и алгоритмы от остальных классов.
		3)	Концентрирует в одном месте код связанный с определенным состоянием.
		4)	Упрощает код контекста.

	-:
		1)	Неоправданно усложняет код, если состояний мало, или они редко меняются.

*/

type RobotState interface {
	Move() string
	Stop() string
	PickUp() string
	SetLight() string
}

type Robot struct {
	state RobotState
}

func (r *Robot) SetState(s RobotState) {
	r.state = s
}

// Методы для работы робота, использующие текущие состояния

func (r *Robot) Move() string {
	return r.state.Move()
}

func (r *Robot) Stop() string {
	return r.state.Stop()
}

func (r *Robot) PickUp() string {
	return r.state.PickUp()
}

func (r *Robot) SetLight() string {
	return r.state.SetLight()
}

// Реализация интерфейса для состояния Движение

type MoveState struct{}

func (m *MoveState) Move() string {
	return "Робот движется"
}

func (m *MoveState) Stop() string {
	return "Нельзя остановить робота во время движения"
}

func (m *MoveState) PickUp() string {
	return "Нельзя подобрать предметы во время движения"
}

func (m *MoveState) SetLight() string {
	return "Робот освещает дорогу"
}

// Реализация интерфейса для состояния Остановка

type StopState struct{}

func (m *StopState) Move() string {
	return "Робот начинает движение"
}

func (m *StopState) Stop() string {
	return "Робот уже остановлен"
}

func (m *StopState) PickUp() string {
	return "Робот может подбирать предметы"
}

func (m *StopState) SetLight() string {
	return "Робот не освещает дорогу во время остановки"
}

// Реализация интерфейса для состояния Остановка

type PickUpState struct{}

func (p *PickUpState) Move() string {
	return "Робот не может двигаться во время поднятия предмета"
}

func (p *PickUpState) Stop() string {
	return "Робот не может останавливаться во время поднятия предмета"
}

func (p *PickUpState) PickUp() string {
	return "Робот уже подбирает предмет"
}

func (p *PickUpState) SetLight() string {
	return "Невозможно освещать дорогу во время поднятия предмета"
}

// Реализация интерфейса для состояния Остановка

type SetLightState struct{}

func (s *SetLightState) Move() string {
	return "Робот уже движется"
}

func (s *SetLightState) Stop() string {
	return "Робот остановлен"
}

func (s *SetLightState) PickUp() string {
	return "Нельзя подобрать предметы во время движения"
}

func (s *SetLightState) SetLight() string {
	return "Робот уже освещает дорогу"
}

func main() {
	robot := &Robot{}
	moveState := &MoveState{}
	stopState := &StopState{}
	pickUpState := &PickUpState{}
	setLightState := &SetLightState{}

	robot.SetState(moveState)
	fmt.Println(robot.Move())
	fmt.Println(robot.Stop())
	fmt.Println(robot.PickUp())
	fmt.Println(robot.SetLight())
	fmt.Println()

	robot.SetState(stopState)
	fmt.Println(robot.Move())
	fmt.Println(robot.Stop())
	fmt.Println(robot.PickUp())
	fmt.Println(robot.SetLight())
	fmt.Println()

	robot.SetState(pickUpState)
	fmt.Println(robot.Move())
	fmt.Println(robot.Stop())
	fmt.Println(robot.PickUp())
	fmt.Println(robot.SetLight())
	fmt.Println()

	robot.SetState(setLightState)
	fmt.Println(robot.Move())
	fmt.Println(robot.Stop())
	fmt.Println(robot.PickUp())
	fmt.Println(robot.SetLight())
}

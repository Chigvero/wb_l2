package main

import (
	"errors"
	"fmt"
)

/*
	Суть паттерна - создать цепочку обработчиков, которые решают, могут ли обработать
	запрос сами, и в противном случае передают запрос дальше по цепочке

	+:
		1)	Уменьшает зависимость между клиентом и обработчиками.
		2)	Реализует принцип единственной обязанности.
		3)	Реализует принцип открытости / закрытости.

	-:
		1)	Запрос может в итоге быть необработанным ни одним из обработчиков.
			Реализовать паттерн «комманда».
*/

type request struct {
	userID   int
	isSecure bool
	resource string
}

type handler interface {
	setNext(handler)
	handle(request) error
}

type authHandler struct {
	next handler
}

func (a *authHandler) setNext(next handler) {
	a.next = next
}

func (a *authHandler) handle(req request) error {
	if req.userID != 0 {
		if a.next != nil {
			return a.next.handle(req)
		}
		return nil
	}
	return errors.New("Unauthorized")
}

type securityHandler struct {
	next handler
}

func (s *securityHandler) setNext(next handler) {
	s.next = next
}

func (s *securityHandler) handle(req request) error {
	if req.isSecure != false {
		if s.next != nil {
			return s.next.handle(req)
		}
		return nil
	}
	return errors.New("Request is not secure")
}

type resourceHandler struct {
	next handler
}

func (r *resourceHandler) setNext(next handler) {
	r.next = next
}

func (r *resourceHandler) handle(req request) error {
	if req.resource != "" {
		if r.next != nil {
			return r.next.handle(req)
		}
		return nil
	}
	return errors.New("Resource not found")
}

func main() {
	request1 := &request{userID: 0, isSecure: true, resource: "/"}
	request2 := &request{userID: 1, isSecure: false, resource: "/admin"}
	request3 := &request{userID: 2, isSecure: true, resource: ""}
	request4 := &request{userID: 3, isSecure: true, resource: "/users"}

	auth := &authHandler{}
	secure := &securityHandler{}
	resource := &resourceHandler{}

	auth.setNext(secure)
	secure.setNext(resource)

	obj := make(map[int]request)
	obj[0] = *request1
	obj[1] = *request2
	obj[2] = *request3
	obj[3] = *request4

	for _, value := range obj {
		if err := auth.handle(value); err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Correct data: ", value)
		}
	}
}

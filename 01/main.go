/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).*/
package main

import "fmt"

type (
	Human struct {
		lastName  string
		firstName string
		surName   string
	}

	Action struct {
		Human
	}
)

func (h Human) FullName() {
	fmt.Printf("Имя: %s\nФамилия: %s\nОтчество: %s\n", h.lastName, h.firstName, h.surName)
}

func (h *Human) SetName(l string, f string, s string) {
	h.lastName = l
	h.firstName = f
	h.surName = s
}

func main() {
	// action := &Action{}
	action := new(Action)
	action.SetName("Ivanov", "Ivan", "Ivanovich")
	action.FullName()
}

/*Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}. */
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var test interface{}
	test = "test"

	// с помощью глагола аннотации %T
	fmt.Printf("%T: %v\n", test, test)

	// с помощью switch type
	switch t := test.(type) {
	case int:
		fmt.Printf("int: %v\n", t)
	case string:
		fmt.Printf("string: %v\n", t)
	case bool:
		fmt.Printf("bool: %v\n", t)
	case chan int:
		fmt.Printf("chan int: %v\n", t)
	}

	// с помощью пакета reflect
	fmt.Printf("%s: %v", reflect.TypeOf(test), test)
}

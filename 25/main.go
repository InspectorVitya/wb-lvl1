/*
	Реализовать собственную функцию sleep.
*/
package main

import (
	"fmt"
	"time"
)

func Sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}

func main() {
	fmt.Println("Before sleep...")
	Sleep(2)
	fmt.Println("After sleep...")
}

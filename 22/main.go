// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20
package main

import (
	"fmt"
	"math/big"
)

func bigAdd(a, b *big.Float) *big.Float {
	sum := new(big.Float)

	return sum.Add(a, b)
}

func bigMultiply(a, b *big.Float) *big.Float {
	mul := new(big.Float)
	return mul.Mul(a, b)
}

func bigDiv(a, b *big.Float) *big.Float {
	div := new(big.Float)
	return div.Quo(a, b)
}

func bigSub(a, b *big.Float) *big.Float {
	sub := new(big.Float)
	return sub.Sub(a, b)
}

func main() {
	a := big.NewFloat(299792)
	b := big.NewFloat(86400)

	fmt.Println(bigAdd(a, b))
	fmt.Println(bigDiv(a, b))
	fmt.Println(bigSub(a, b))
	fmt.Println(bigMultiply(a, b))
}

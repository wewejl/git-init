package main

import (
	"fmt"
	"gitTest/calc"
)

func main() {
		fmt.Println("calc.Add called!")
		res := calc.Add(10, 20)
		fmt.Println("Add(10 ,20) :", res)
		fmt.Println("calc.Sub called!")
		res = calc.Sub(30, 20)
		fmt.Println("Sub(10 ,20) :", res)

		fmt.Println("calc.chufa called!")
		res = calc.Calc(100, 10)

		fmt.Println("calc.multi called!")
		res = calc.Sub(20, 10)

		fmt.Println("Sub(10 ,20) :", res)
}
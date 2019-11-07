package calc


import "fmt"

func Calc(a ,b int) int {

	fmt.Println("chufa call...")

	if b==0 {
		panic("分母不能为空")
	}
	return a/b
}

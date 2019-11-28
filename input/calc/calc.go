package main

import (
	"../../interface/stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 编写一个简单的逆波兰式计算器，
// 它接受用户输入的整型数（最大值 999999）
// 和运算符 +、-、*、/。
// 输入的格式为：number1 ENTER number2 ENTER
// operator ENTER --> 显示结果
// 当用户输入字符 'q' 时，程序结束。
// 请使用您在练习11.13中开发的 stack 包。

var s = new(stack.Stack)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err)
			os.Exit(1)
		}
		Calc(input)
	}
}

func Calc(input string) {
	key := input[0 : len(input)-2]
	switch {
	case key == "q":
		os.Exit(0)
	case key >= "0" && key <= "999999":
		i, _ := strconv.Atoi(key)
		s.Push(i)
	case key == "+":
		i, _ := s.Pop()
		j, _ := s.Pop()
		fmt.Printf("result: %d+%d=%d", i.(int), j.(int), i.(int)+j.(int))
	case key == "-":
		i, _ := s.Pop()
		j, _ := s.Pop()
		fmt.Printf("result: %d-%d=%d", i.(int), j.(int), i.(int)-j.(int))
	case key == "*":
		i, _ := s.Pop()
		j, _ := s.Pop()
		fmt.Printf("result: %d*%d=%d", i.(int), j.(int), i.(int)*j.(int))
	case key == "/":
		i, _ := s.Pop()
		j, _ := s.Pop()
		fmt.Printf("result: %d/%d=%f", i.(int), j.(int), i.(float64)/j.(float64))
	default:
		fmt.Print("none")
	}

}

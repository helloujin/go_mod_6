package main

import (
	"github.com/helloujin/go_mod_6/calc"

	"fmt"
)

func main() {
	var oprnd1, oprnd2, result float64
	var operator string

	fmt.Print("введите первое число: ")
	_, err := fmt.Scanln(&oprnd1)
	if err != nil {
		fmt.Println(fmt.Sprintf("Число не прочитано %s", err))

	}
	// fmt.Println(oprnd1) - убрал для визуального удобства в терминале вывода инфы.

	fmt.Print("укажите операцию: (+,-,*,/) ")
	_, err = fmt.Scanln(&operator)
	if err != nil { // не знаю как реализовать проверку на корректный ввод операции. Через break
		// или по-другому условие прописать надо.
		fmt.Println(fmt.Sprintf("некорректный формат %s", err))
		if operator != "+, -, *, /" {
			fmt.Println("oshibka")
		}
	}

	// fmt.Println(operator)- убрал для визуального удобства в терминале вывода инфы.

	fmt.Print("введите второе число: ")
	_, err = fmt.Scanln(&oprnd2)
	if err != nil {
		fmt.Println(fmt.Sprintf("Число не прочитано %s", err))
	}
	// fmt.Println(oprnd2)- убрал для визуального удобства в терминале вывода инфы.

	calculator := calc.NewCalculator()
	result = calculator.Calculate(oprnd1, oprnd2, operator)

	fmt.Println(result)
}

package main

import "fmt"

func main() {
	var oprnd1, oprnd2, result float32
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

	switch operator {
	case "+":
		result = oprnd1 + oprnd2
		fmt.Println("Результат суммы: ", result)

	case "-":
		result = oprnd1 - oprnd2
		fmt.Println("Результат разности: ", result)

	case "*":
		result = oprnd1 * oprnd2
		fmt.Println("Результат умножения: ", result)

	case "/":
		if oprnd2 == 0 { // проверим делимое на условие не равности нулю
			fmt.Println("Не делим на ноль")
		} else {
			result = oprnd1 / oprnd2
			fmt.Println("Результат частного: ", result)
		}

	default:
		fmt.Println("Не указана запрашиваемая операция")

	}

}

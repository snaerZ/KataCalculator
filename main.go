package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

Calc:
	for {
		var num1, action, num2, action2, num3 string

		fmt.Println("Введите математическое вырежение:")
		fmt.Scanln(&num1, &action, &num2, &action2, &num3)

		toNum1, _ := strconv.Atoi(num1) // Преобразование тексата из переменной num1 в число (переменная toNum1)
		toNum2, _ := strconv.Atoi(num2) // Преобразование тексата из переменной num1 в число (переменная toNum1)

		// Проверка чисел для вывода ошибок

		if action2 == "" { // Проверяем на наличие второго оператора или операнда
		} else if num3 == "" {
		} else {
			log.Fatalln("Ошибка! Выражение должно состоять из двух операндов и одного оператора!")
		}

		if toNum1 == 0 && toNum2 != 0 { // Проверяем, что оба числа арабские или римские
			log.Fatalln("Ошибка! Оба числа должны быть арабскими или римскими!")
		} else if toNum1 != 0 && toNum2 == 0 {
			log.Fatalln("Ошибка! Оба числа должны быть арабскими или римскими!")
		} else {
		}

		if action == "+" || action == "-" { // Проверка математического знака
		} else if action == "*" || action == "/" {
		} else {
			log.Fatalln("Ошибка! Неизвестное маметатическое действие!")
		}

		if toNum1 > 10 { //Проверка, чтобы числа были не больше 10
			log.Fatalln("Ошибка! Числа должны быть не больше 10!")
			//panic("Ошибка! Числа должны быть не больше 10!")
		}

		if toNum2 > 10 { //Проверка, чтобы числа были не больше 10
			log.Fatalln("Ошибка! Числа должны быть не больше 10!")
			//panic("Ошибка! Числа должны быть не больше 10!")
		}

		// Проверяем преобразованные числа, переходим к нужному калькулятору

		if toNum1 != 0 && toNum2 != 0 { //Если числа string преобразовались в int :
			goto ArabCalc // То переходим к арабскому калькулятору
		} else { // Если нет :
			goto RimCalc // То переходим к римскому калькулятору
		}

	ArabCalc:
		// **********************************************************************************
		//                      Калькулятор для целых арабских чисел
		// **********************************************************************************

		switch action { // Создаём дествия для разных значений в переменной action

		case "*":
			fmt.Println(toNum1, action, toNum2, "=", toNum1*toNum2)

		case "/":
			fmt.Println(toNum1, action, toNum2, "=", toNum1/toNum2)

		case "+":
			fmt.Println(toNum1, action, toNum2, "=", toNum1+toNum2)

		case "-":
			fmt.Println(toNum1, action, toNum2, "=", toNum1-toNum2)
		}

	RimCalc:
		// **********************************************************************************
		//                      Калькулятор для целых римских чисел
		// **********************************************************************************

		if toNum1 > 0 && toNum2 > 0 {
			goto Calc // break
		}

		var Rim1, Rim2 int
		var resultInt int

		switch num1 { // Расшифровываем римские цифры, записывая значение в переменную Rim1
		case "I":
			Rim1 = 1
		case "II":
			Rim1 = 2
		case "III":
			Rim1 = 3
		case "IV":
			Rim1 = 4
		case "V":
			Rim1 = 5
		case "VI":
			Rim1 = 6
		case "VII":
			Rim1 = 7
		case "VIII":
			Rim1 = 8
		case "IX":
			Rim1 = 9
		case "X":
			Rim1 = 10
		}

		if Rim1 == 0 {
			log.Fatalln("Ошибка! Римское число больше 10!")
		}

		switch num2 { // Расшифровываем римские цифры, записывая значение в переменную Rim2
		case "I":
			Rim2 = 1
		case "II":
			Rim2 = 2
		case "III":
			Rim2 = 3
		case "IV":
			Rim2 = 4
		case "V":
			Rim2 = 5
		case "VI":
			Rim2 = 6
		case "VII":
			Rim2 = 7
		case "VIII":
			Rim2 = 8
		case "IX":
			Rim2 = 9
		case "X":
			Rim2 = 10
		}

		if Rim2 == 0 {
			log.Fatalln("Ошибка! Римское число больше 10!")
		}

		switch action { // Создаём дествия для разных значений в переменной action

		case "*":
			resultInt = Rim1 * Rim2

		case "/":
			if Rim2 > Rim1 {
				log.Fatalln("Ошибка! Римское число не может получится отрицательным!")
			} else {
				resultInt = Rim1 / Rim2
			}

		case "+":
			resultInt = Rim1 + Rim2

		case "-":
			if Rim2 >= Rim1 {
				log.Fatalln("Ошибка! Римское число не может получиться отрицательным или 0!")
			} else {
				resultInt = Rim1 - Rim2
			}
		}

		//Переводим ответ из int обратно в римские цифры

		resultRim1 := resultInt / 10 // Десятки результата выражения
		resultRim2 := resultInt % 10 // Остаток от деления результата выражения, т.е. единицы

		var resultRimDes string // Переменная для десятков
		var resultRimEd string  // Переменная для единиц

		switch resultRim1 { // Переводим десятки в римские числа
		case 0:
			resultRimDes = ""
		case 1:
			resultRimDes = "X"
		case 2:
			resultRimDes = "XX"
		case 3:
			resultRimDes = "XXX"
		case 4:
			resultRimDes = "XL"
		case 5:
			resultRimDes = "L"
		case 6:
			resultRimDes = "LX"
		case 7:
			resultRimDes = "LXX"
		case 8:
			resultRimDes = "LXXX"
		case 9:
			resultRimDes = "XC"
		case 10:
			resultRimDes = "C"
		}

		switch resultRim2 { // Переводим еденицы в римские числа
		case 0:
			resultRimEd = ""
		case 1:
			resultRimEd = "I"
		case 2:
			resultRimEd = "II"
		case 3:
			resultRimEd = "III"
		case 4:
			resultRimEd = "IV"
		case 5:
			resultRimEd = "V"
		case 6:
			resultRimEd = "VI"
		case 7:
			resultRimEd = "VII"
		case 8:
			resultRimEd = "VIII"
		case 9:
			resultRimEd = "IX"
		}

		resultMain := strings.TrimSpace(resultRimDes + resultRimEd) // Записываем результат десятков и едениц, удаляя пробелы
		fmt.Println(num1, action, num2, "=", resultMain)            // Выводим результат без пробелов
	}
}

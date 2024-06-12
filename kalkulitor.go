package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a, b, c, err string
	fmt.Scanln(&a, &b, &c, &err)
	if err != "" { //проверка на соответствие матиматической операции
		panic("Не более двух чисел и одного арифметического действия!!")
	}

	m := map[string]int{ //мапа римских чисел
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	rimA := 0
	rimC := 0
	rimAA := 10
	rimCC := 10
	for kluch, allement := range m { // проверка на римское число и присвоение значения
		if kluch == a {
			rimA = allement
			rimAA = allement
		}
		if kluch == c {
			rimC = allement
			rimCC = allement
		}
	}

	if b != "+" && b != "-" && b != "/" && b != "*" { //проверка на верность введённого арифметического действия
		panic("Введите корректное арифметическое действие!!!")
	}

	if (rimA == 0 && rimC != 0) || (rimC == 0 && rimA != 0) { //проверка на ввод  двух арабсикх или двух римских чисел
		panic("Нельзя производить операции между арабскими и римскими числами!!")
	}

	if (rimAA <= 0 || rimAA > 10) || (rimCC <= 0 || rimCC > 10) { //проверка римских чисел  от 1 до 10 включительно
		panic("Римские числа от I до X включительно!!")
	}

	if (rimA >= 1 && rimA <= 10) && (rimC >= 1 && rimC <= 10) { //если числа римские срабатывает условие и происходт решение примера через вспомогательные функции
		switch b {
		case "+":
			fmt.Println(rim(plus(rimA, rimC)))
		case "-":
			fmt.Println(rim(minus(rimA, rimC)))
		case "/":
			fmt.Println(rim(delenie(rimA, rimC)))
		case "*":
			fmt.Println(rim(umnoj(rimA, rimC)))
		}
	} else { //если число арабское производит расчёт арабского и выдаёт ответ
		switch b {
		case "+":
			fmt.Println(plus(arab(a), arab(c)))
		case "-":
			fmt.Println(minus(arab(a), arab(c)))
		case "/":
			fmt.Println(delenie(arab(a), arab(c)))
		case "*":
			fmt.Println(umnoj(arab(a), arab(c)))
		}
	}

}

func plus(a int, c int) int { // функция сложения
	rezultat := a + c
	return rezultat
}

func minus(a int, c int) int { // функция вычетания
	rezultat := a - c
	return rezultat
}

func delenie(a int, c int) int { // функция деления
	rezultat := a / c
	return rezultat
}

func umnoj(a int, c int) int { // функция умножения
	rezultat := a * c
	return rezultat
}

func arab(n string) int { // функция преобразования из строки в число арабское
	x, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if x <= 0 || x > 10 { //проверка арабсикх чисел  от 1 до 10 включительно
		panic("Арабские числа от 1 до 10 включительно!!")
	}
	return x
}

func rim(n int) string { // функция преобразования из арабского числа в  римское
	if n <= 0 { //проверка римского ответа на положительное число
		panic("Результат арифметических действий римских чисел не может быть ниже или равен нулю")
	}
	rml := []struct {
		val int
		rom string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	result := ""
	for _, v := range rml {
		for n >= v.val {
			result += v.rom
			n -= v.val
		}
	}
	return result

}

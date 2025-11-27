package main

import (
	calculator_ "calculator/Calculator_"
	convector_ "calculator/Convector_"
	"fmt"
)

func main() {
	for {
		var text string
		fmt.Print("Введите пример (или 'exit' для выхода): ")
		fmt.Scanln(&text)

		if text == "exit" {
			fmt.Println("Выход из программы...") // это для возможности выйти из цикла проги
			break
		}

		a, operator, b, err := convector_.StringToRunes(text) //Вызвал функцию из пакета convector для вывода результатов

		if err != nil {
			fmt.Println("Ошибка:", err)
			//делаю лютую проверку ошибок     типо если в пакете convector произошла ошибка по оператору или числам то сообщаю об этом
			continue
		}
		calc := calculator_.New(operator) //вызываю результат функцию которая принемает в качестве аргумента (оператора + - * /)
		// внутри функции оператор гоняется по switch. Если там все совпадает то аргументы a и b попадают в метод который выполняется в нужной структуре и проверяетя условно через интерфейс
		result := calc.Calculate(a, b) // присваиваю результат метода Calculate из calc
		fmt.Printf("Результат: %d %s %d = %d\n\n", a, operator, b, result)
	}

}

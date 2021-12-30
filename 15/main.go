package main

/*
	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.

	var justString string
	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}
*/

/*
	Использование глобальной переменной - плохая практика
	В го строки = слайс байтов, доступный только для чтения
	Так взятие sub-slice от v в данном случае приведет к тому, что в памяти будет храниться
	целый слайс созданный createHugeString, а не только интересующую нас часть в 100 элементов
	такое использование не даст очистить сборщику мусора неиспользуемую часть данных

	решение с использованием глобальной переменной:
	var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		out := append([]byte{}, v[:100]...)
		justString = string(out)
	}
	func main() {
		someFunc()
	}

	решение без использованием глобальной переменной:
		func someFunc() []byte {
			v := createHugeString(1 << 10)
			out := append([]byte{}, v[:100]...)
			return out
		}
		func main() {
			str := someFunc()
		}
*/

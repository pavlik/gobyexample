// [_Аргументы командной строки_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// это один из самых популярных способов параметризации программы.
// К примеру, `go run hello.go` использует `run` и
// `hello.go` как аргументы для запуска программы `go`.

package main

import "os"
import "fmt"

func main() {

	// `os.Args` позволяет получить доступ к аргументам
	// командной строки. Обратите внимание, что первое
	// значение это путь к самой программе, а
	// `os.Args[1:]` содержит только аргументы.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// К отдельным аргументам можно обратиться по индексу.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

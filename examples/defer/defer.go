// _Defer_ гарантирует отложенный вызов функциии, обычно
// отчистки,  по мере выполнения программы. `defer` часто
// используется подобно `ensure` и `finally` в других языках.

package main

import "fmt"
import "os"

// Предположим, мы хотели бы создать файл, записать в него
// данные и затем закрыть его, когда закончим. Вот, что мы
// могли бы сделать с `defer`.
func main() {

    // После получения объекта файла с помощью функции
    // `createFile`, мы откладываем закрытие файла с
    // помощью функции `closeFile`. Выполнение произойдет
    // в конце блока текущей функции (`main`) после того,
    // как `writeFile` выполнит свою работу.
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}

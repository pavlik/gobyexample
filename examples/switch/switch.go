// Оператор _switch_ предназначен для организации
// выбора из множества различных вариантов.

package main

import "fmt"
import "time"

func main() {

    // Пример простого оператора `switch`.
    i := 2
    fmt.Print("write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // Вы можете использовать запятые для разделения
    // нескольких выражений в одном операторе `case`.
    // Также в этом примере мы используем
    // необязательный оператор `default`.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

    // `switch` без выражения - это альтернативный путь
    // выразить логику if/else. В этом примере
    // также показано, как `case` выражения
    // могут быть не постоянными.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    default:
        fmt.Println("it's after noon")
    }
}

// todo: type switches

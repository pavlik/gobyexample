// Основные отправка и приём из каналов являются
// блокирующими. Но мы можем использовать `select`
// с оператором `default` для реализации _неблокирующих_
// отправки, приёма и даже неблокирующего
// многовариантного `select`.

package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    // Здесь неблокирующий приём. Если значение доступно
    // в `messages`, то `select` выберет вариант
    // `<-messages` с этим значением. Если нет
    // то будет сразу выбран вариант `default`.

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // Неблокирующая отправка работает аналогично.
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    // Мы можем использовать несколько `case` перед
    // оператором `default` для реализации многовариантного
    // неблокирующего выбора. Здесь мы пытаемся неблокирующим
    // способом получить как `messages` так и `signals`.

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}

// В Go _select_ позволяет ожидать выполнения
// многоканальных операций. Cочетание горутин
// и каналов с оператором select является
// сильной стороной Go.

package main

import "time"
import "fmt"

func main() {

    // В нашем примере будем делать выбор из двух каналов.
    c1 := make(chan string)
    c2 := make(chan string)

    // Каждый канал будет получать значение после некоторого
    // промежутка времени, например, для имитации
    // блокирования RPC операций в одновременно
    // выполняющихся горутинах.
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    // Будем использовать `select` для ожидания обоих
    // значений одновременно и печати каждого по мере
    // получения.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}

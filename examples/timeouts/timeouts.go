// _Тайм-ауты_ важны для программ, которые подключаются ко
// внешним ресурсам или нуждаются в ограничении
// времени выполнения. Реализация тайм-аутов в Go проста и
// красива благодаря каналам и оператору `select`.

package main

import "time"
import "fmt"

func main() {

    // В нашем примере предположим, что мы выполняем внешний
    // вызов, который возвращает свой результат в канал `c1`
    // спустя 2 сек.
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    // Здесь `select` выполняет вариант с тайм-аутом.
    // `res := <-c1` ожидает результата и `<-Time.After`
    // ожидает значение для отправки после тайм-аута
    // в 1 сек. Поскольку `select` продолжается
    // при первом успешном приёме, мы выберем вариант
    // с тайм-аутом, если операция продолжится
    // дольше дозволенной 1 сек.

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    // Если мы указываем длинный тайм-аут в 3 сек., то
    // приём из `c2` будет успешен и мы напечатаем `result 2`.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}

// todo: cancellation?

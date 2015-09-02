// При использовании каналов в качестве параметров функции
// вы можете указать, будет ли канал предназначен только
// для отправки или только для получения значений. Это
// указание повышает безопасность типов программы.

package main

import "fmt"

// Функция `ping` принимает только канал для отправки
// значений. Попытка получения в этом канале вызовет
// ошибку компиляции.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Функция `pong` принимает один канал для получения
// (`pings`) и второй для отправки (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// Базовые операции отправки и приёма из каналов являются
// блокирующими. Но можно использовать `select`
// с оператором `default` для реализации _неблокирующих_
// отправки, приёма и даже неблокирующего
// многовариантного `select`.

package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	// Здесь показан неблокирующий приём.
	// Если значение доступно
	// в `messages`, то `select` выберет вариант
	// `<-messages` с этим значением. Если нет,
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

	// Можно использовать несколько `case` перед
	// оператором `default` для реализации многовариантного
	// неблокирующего выбора. Здесь показана попытка
	// неблокирующим способом получить
	// как `messages` так и `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

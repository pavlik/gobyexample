// _[Ограничение скорости
// (Rate limiting)](http://en.wikipedia.org/wiki/Rate_limiting)_
// — важный способ для управления использованием ресурсов
// и обеспечения качества обслуживания. Go предлагает
// элегантный способ ограничения скорости с помощью горутин,
// каналов и [счётчиков тиков](tickers).

package main

import "time"
import "fmt"

func main() {

	// Сначала обратим внимание на базовое ограничение
	// скорости. Предположим, что нужно ограничить скорость
	// обработки входящих запросов.  Будем брать эти запросы
	// из канала  `requests`.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Канал `limiter` будет получать значение каждые
	// 200 миллисекунд. Это будет регулятор скорости
	// в нашей схеме.
	limiter := time.Tick(time.Millisecond * 200)

	// С помощью блокировки приёма значения из канала `limiter`
	// перед обработкой каждого запроса, ограничимся 1 запросом
	// в 200 миллисекунд.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Может потребоваться организовать небольшие пакеты
	// запросов в схеме ограничения скорости с сохранением
	// общего ограничения скорости. Это можно сделать с
	// помощью буфера канала-ограничителя `burstyLimiter`.
	// Такой канал позволяет использовать размер пакета
	// велчиной до 3 событий.
	burstyLimiter := make(chan time.Time, 3)

	// Заполним канал для того, чтобы показать позволенный
	// размер пакета
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Каждые 200 миллисекунд пробуем добавить новое
	// значение в канал `burstyLimiter` до предела в 3
	// значения.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// Теперь имитируем 5 входящих запросов. Первые
	// 3 из них получат преимущество из-за возможности
	// пакетной обработки `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

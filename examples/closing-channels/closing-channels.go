// _Закрытие_ канала показывает, что больше нельзя
// будет отправить значения через него.
// Это может быть полезно для сообщения о завершении
// приёмникам канала.

package main

import "fmt"

// В этом примере мы используем канал `jobs` для
// сообщения о завершении работы от горутины `main()`
// рабочей горутине. Когда у нас больше нет задач для
// рабочей горутины мы закрываем канал `jobs`.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.

	// Это рабочая горутина. Она постоянно получает данные из
	// `jobs` посредством `j, more := <-jobs`. В этой
	// специальной форме получения данных с 2-мя значениями,
	// значение `more` будет `false` если `jobs` был закрыт
	// и все значения в канале уже были получены.
	// Мы используем этот подход для уведомления на
	// `done` в момент, когда обработаны все задачи.

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Отправляем 3 задачи рабочей горутине через
	// канал `jobs` и закрываем его.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.
	// Ждем рабочую горутину используя подход
	// [синхронизации](channel-synchronization),
	// который был рассмотрен ранее.
	<-done
}

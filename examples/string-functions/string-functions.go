// Пакет стандартной библиотеки `strings` предоставляет много
// полезных функций для работы со строками. Вот, несколько примеров,
// отображающие возможности пакета.

package main

import s "strings"
import "fmt"

// Присвоим переменной `p` функцию `fmt.Println` для сокращения
// имени, которую мы будем использовать ниже.
var p = fmt.Println

func main() {

	// Данные функции доступны в пакете `strings`.
	// Обратите внимание, что все эти функции из пакета, а
	// не методы строковых объектов. Это означает, что нам
	// необходимо передать первым аргументом функции, строку,
	// над которой мы производим операцию.
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// Вы можете найти больше функций из документации
	// к пакету [`strings`](http://golang.org/pkg/strings/)

	// Примеры ниже не относятся к пакету `strings`, но о них
	// стоит упомянуть - это механизмы для получения длины строки
	// и получение символа по индексу.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
